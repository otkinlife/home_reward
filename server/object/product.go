package object

import (
	"encoding/json"
	"errors"
	"home-reward/server/helper"
)

const StatusNeed = 1
const StatusGot = 2

var ProductList = map[int64]*Product{}
var LastProductID int64

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Reward      int64  `json:"reward"`
	Status      int64  `json:"status"`
	CharacterID int64  `json:"character_id"`
	Publisher   int64  `json:"publisher"`
	Other       string `json:"other"`
}

func InitProductList() {
	var err error
	helper.ProductsFile, err = helper.InitFile(GlobalConfig.DataPath + "/" + GlobalConfig.ProductFileName)
	if err != nil {
		panic(err)
	}
	bytes, err := helper.ProductsFile.GetBytes()
	if err != nil {
		panic(err)
	}
	if len(bytes) == 0 {
		return
	}
	err = json.Unmarshal(bytes, &ProductList)
	if err != nil {
		panic(err)
	}
	max := int64(0)
	for k, _ := range ProductList {
		if k > max {
			max = k
		}
	}
	LastProductID = max
}

func CreateProduct(name string, reward int64, other string) error {
	ProductList[LastProductID+1] = &Product{
		ID:          LastProductID + 1,
		Name:        name,
		Reward:      reward,
		Status:      StatusNeed,
		CharacterID: 0,
		Publisher:   CurrentCharacter(),
		Other:       other,
	}
	LastProductID += 1
	return commitToProductFile()
}

func DeleteProduct(ID int64) error {
	delete(ProductList, ID)
	return commitToProductFile()
}

func BuyProduct(productID int64) error {
	if product, ok := ProductList[productID]; ok {
		if product.Status == StatusNeed {
			if _, ok := Characters[CurrentCharacter()]; ok {
				err := Characters[CurrentCharacter()].DownReward(product.Reward)
				if err != nil {
					return err
				}
			}
			ProductList[productID].Status = StatusGot
			ProductList[productID].CharacterID = CurrentCharacter()
			return commitToProductFile()
		} else {
			return errors.New("该物品不能购买！")
		}
	} else {
		return errors.New("该物品不存在！")
	}
}

func CancelBuyProduct(productID int64) error {
	if product, ok := ProductList[productID]; ok {
		if product.Status == StatusGot {
			if _, ok := Characters[product.CharacterID]; ok {
				err := Characters[product.CharacterID].UpReward(product.Reward)
				if err != nil {
					return err
				}
			}
			ProductList[productID].Status = StatusNeed
			ProductList[productID].CharacterID = 0
			return commitToProductFile()
		}
		if product.Status == StatusNeed {
			return errors.New("该物品未被购买！")
		}
		return nil
	} else {
		return errors.New("该物品不存在！")
	}
}

func commitToProductFile() error {
	jsonString, err := json.Marshal(ProductList)
	if err != nil {
		return err
	}
	_ = helper.ProductsFile.Truncate(0)
	_, _ = helper.ProductsFile.Seek(0, 0)
	_, err = helper.ProductsFile.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}

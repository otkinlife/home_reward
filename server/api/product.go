package api

import (
	"encoding/json"
	"fmt"
	"home-reward/server/object"
	"net/http"
	"strconv"
)

func GetProductList(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	data := map[int64][]*object.Product{}
	for _, product := range object.ProductList {
		data[product.Status] = append(data[product.Status], product)
	}
	resp.Data = data
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	name := r.URL.Query().Get("name")
	if name == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("物品名不存在！")
		return
	}
	reward := r.URL.Query().Get("reward")
	if reward == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("赏金不存在！")
		return
	}
	rewardNumber, err := strconv.Atoi(reward)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	other := r.URL.Query().Get("other")
	err = object.CreateProduct(name, int64(rewardNumber), other)
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "物品发布成功"
}

func DelProduct(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("物品不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = object.DeleteProduct(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "物品删除成功"
}

func BuyProduct(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("物品不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}

	err = object.BuyProduct(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "物品购买成功"
}

func CancelBuyProduct(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("物品不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = object.CancelBuyProduct(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "物品取消购买成功"
}

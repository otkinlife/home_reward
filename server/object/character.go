package object

import (
	"encoding/json"
	"errors"
	"home-reward/server/dao"
	"home-reward/server/helper"
)

var Characters = map[int64]*Character{}

type Character struct {
	ID     int64  `json:"id"`
	Nick   string `json:"nick"`
	Reward int64  `json:"reward"`
	Avatar string `json:"avatar"`
}

func InitCharacters() {
	rows, err := dao.DB.Query("select * from `character`")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		character := Character{}
		err = rows.Scan(&character.ID, &character.Nick, &character.Reward, &character.Avatar)
		if err != nil {
			panic(err)
		}
		Characters[character.ID] = &character
	}
}

func (c *Character) UpReward(reward int64) error {
	c.Reward += reward
	return commitToCharacterFile()
}

func (c *Character) DownReward(reward int64) error {
	if reward > c.Reward {
		return errors.New("余额不足")
	}
	c.Reward -= reward
	return commitToCharacterFile()
}

func commitToCharacterFile() error {
	jsonString, err := json.Marshal(Characters)
	if err != nil {
		return err
	}
	_ = helper.CharactersFile.Truncate(0)
	_, _ = helper.CharactersFile.Seek(0, 0)
	_, err = helper.CharactersFile.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}

func CurrentCharacter() int64 {
	return int64(1)
}

package wish

import (
	"home-reward/server/character"
)

func List() (map[int64]Wish, error) {
	wishes := map[int64]Wish{}
	err := getWishes(&wishes)
	if err != nil {
		return wishes, err
	}
	return wishes, nil
}

func One(ID int64) (Wish, error) {
	wish := Wish{}
	err := getWishByID(ID, &wish)
	if err != nil {
		return wish, err
	}
	return wish, nil
}

func Create(name string, reward int64, other string) error {
	w := Wish{
		Name:        name,
		Reward:      reward,
		Status:      StatusWant,
		Other:       other,
		CharacterID: 0,
		Publisher:   character.Current().ID,
	}
	return save(w)
}

func Delete(ID int64) error {
	w, err := One(ID)
	if err != nil {
		return err
	}
	return delete(w)
}

func Finish(ID int64) error {
	w, err := One(ID)
	if err != nil {
		return err
	}
	if w.Status == StatusWant {
		w.Status = StatusFinished
		err = character.ReduceReward(*character.Current(), w.Reward)
		if err != nil {
			return err
		}
		return save(w)
	}
	return nil
}

func CancelFinish(ID int64) error {
	w, err := One(ID)
	if err != nil {
		return err
	}
	if w.Status == StatusFinished {
		w.Status = StatusWant
		err = character.AddReward(*character.Current(), w.Reward)
		if err != nil {
			return err
		}
		return save(w)
	}
	return nil
}

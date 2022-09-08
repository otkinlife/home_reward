package wish

import (
	"home-reward/server/character"
)

type Logic struct {
	IP               string
	CurrentCharacter *character.Character
}

func NewLogic(ip string) *Logic {
	return &Logic{
		IP:               ip,
		CurrentCharacter: character.Current(ip),
	}
}

func (l *Logic) List() (map[int64]Wish, error) {
	wishes := map[int64]Wish{}
	err := getWishes(&wishes)
	if err != nil {
		return wishes, err
	}
	return wishes, nil
}

func (l *Logic) One(ID int64) (Wish, error) {
	wish := Wish{}
	err := getWishByID(ID, &wish)
	if err != nil {
		return wish, err
	}
	return wish, nil
}

func (l *Logic) Create(name string, reward int64, other string) error {
	w := Wish{
		Name:        name,
		Reward:      reward,
		Status:      StatusWant,
		Other:       other,
		CharacterID: 0,
		Publisher:   l.CurrentCharacter.ID,
	}
	return save(w)
}

func (l *Logic) Delete(ID int64) error {
	w, err := l.One(ID)
	if err != nil {
		return err
	}
	return delete(w)
}

func (l *Logic) Finish(ID int64) error {
	w, err := l.One(ID)
	if err != nil {
		return err
	}
	if w.Status == StatusWant {
		w.Status = StatusFinished
		characterLogic := character.NewLogic(l.IP)
		err = characterLogic.ReduceReward(l.CurrentCharacter, w.Reward)
		if err != nil {
			return err
		}
		return save(w)
	}
	return nil
}

func (l *Logic) CancelFinish(ID int64) error {
	w, err := l.One(ID)
	if err != nil {
		return err
	}
	if w.Status == StatusFinished {
		w.Status = StatusWant
		characterLogic := character.NewLogic(l.IP)
		err = characterLogic.AddReward(l.CurrentCharacter, w.Reward)
		if err != nil {
			return err
		}
		return save(w)
	}
	return nil
}

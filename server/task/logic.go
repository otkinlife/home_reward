package task

import (
	"errors"
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

func (l *Logic) List() (map[int64]Task, error) {
	tasks := map[int64]Task{}
	err := getTasks(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (l *Logic) One(ID int64) (Task, error) {
	task := Task{}
	err := getTaskByID(ID, &task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (l *Logic) Create(name string, reward int64) error {
	t := Task{
		Name:        name,
		Reward:      reward,
		Status:      StatusToDo,
		CharacterID: 0,
	}
	return save(t)
}

func (l *Logic) Delete(ID int64) error {
	t, err := l.One(ID)
	if err != nil {
		return err
	}
	return delete(t)
}

func (l *Logic) Get(ID int64) error {
	t, err := l.One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != 0 {
		return errors.New("该任务已被领取")
	}
	if t.Status == StatusToDo {
		t.Status = StatusDoing
		t.CharacterID = l.CurrentCharacter.ID
		return save(t)
	}
	return nil
}

func (l *Logic) CancelGet(ID int64) error {
	t, err := l.One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != l.CurrentCharacter.ID {
		return errors.New("不是自己的任务不能取消")
	}
	if t.Status == StatusDoing {
		t.Status = StatusToDo
		t.CharacterID = 0
		return save(t)
	}
	return nil
}

func (l *Logic) Finish(ID int64) error {
	t, err := l.One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != l.CurrentCharacter.ID {
		return errors.New("不是自己的任务不能取消")
	}
	if t.CharacterID != l.CurrentCharacter.ID {
		return errors.New("不是自己的任务不能完成")
	}
	if t.Status == StatusDoing {
		t.Status = StatusDone
		characterLogic := character.NewLogic(l.IP)
		err = characterLogic.AddReward(l.CurrentCharacter, t.Reward)
		if err != nil {
			return err
		}
		return save(t)
	}
	return nil
}

func (l *Logic) CancelFinish(ID int64) error {
	t, err := l.One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != l.CurrentCharacter.ID {
		return errors.New("不是自己的任务不能取消")
	}
	if t.Status == StatusDone {
		t.Status = StatusDoing
		characterLogic := character.NewLogic(l.IP)
		err = characterLogic.ReduceReward(l.CurrentCharacter, t.Reward)
		if err != nil {
			return err
		}
		return save(t)
	}
	return nil
}

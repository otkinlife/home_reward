package task

import (
	"errors"
	"home-reward/server/character"
)

func List() (map[int64]Task, error) {
	tasks := map[int64]Task{}
	err := getTasks(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func One(ID int64) (Task, error) {
	task := Task{}
	err := getTaskByID(ID, &task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func Create(name string, reward int64) error {
	t := Task{
		Name:        name,
		Reward:      reward,
		Status:      StatusToDo,
		CharacterID: 0,
	}
	return save(t)
}

func Delete(ID int64) error {
	t, err := One(ID)
	if err != nil {
		return err
	}
	return delete(t)
}

func Get(ID int64) error {
	t, err := One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != 0 {
		return errors.New("该任务已被领取")
	}
	if t.Status == StatusToDo {
		t.Status = StatusDoing
		t.CharacterID = character.Current().ID
		return save(t)
	}
	return nil
}

func CancelGet(ID int64) error {
	t, err := One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != character.Current().ID {
		return errors.New("不是自己的任务不能取消")
	}
	if t.Status == StatusDoing {
		t.Status = StatusToDo
		t.CharacterID = 0
		return save(t)
	}
	return nil
}

func Finish(ID int64) error {
	t, err := One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != character.Current().ID {
		return errors.New("不是自己的任务不能取消")
	}
	if t.CharacterID != character.Current().ID {
		return errors.New("不是自己的任务不能完成")
	}
	if t.Status == StatusDoing {
		t.Status = StatusDone
		err = character.AddReward(*character.Current(), t.Reward)
		if err != nil {
			return err
		}
		return save(t)
	}
	return nil
}

func CancelFinish(ID int64) error {
	t, err := One(ID)
	if err != nil {
		return err
	}
	if t.CharacterID != character.Current().ID {
		return errors.New("不是自己的任务不能取消")
	}
	if t.Status == StatusDone {
		t.Status = StatusDoing
		err = character.ReduceReward(*character.Current(), t.Reward)
		if err != nil {
			return err
		}
		return save(t)
	}
	return nil
}

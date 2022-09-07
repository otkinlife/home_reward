package object

import (
	"encoding/json"
	"errors"
	"home-reward/server/helper"
)

const StatusToDo = 1
const StatusDoing = 2
const StatusDone = 3

var TaskList = map[int64]*Task{}
var LastID int64

type Task struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Reward      int64  `json:"reward"`
	Status      int64  `json:"status"`
	CharacterID int64  `json:"character_id"`
}

func InitTaskList() {
	var err error
	helper.TasksFile, err = helper.InitFile(GlobalConfig.DataPath + "/" + GlobalConfig.TaskFileName)
	if err != nil {
		panic(err)
	}
	bytes, err := helper.TasksFile.GetBytes()
	if err != nil {
		panic(err)
	}
	if len(bytes) == 0 {
		return
	}
	err = json.Unmarshal(bytes, &TaskList)
	if err != nil {
		panic(err)
	}
	max := int64(0)
	for k, _ := range TaskList {
		if k > max {
			max = k
		}
	}
	LastID = max
}

func CreateTask(name string, reward int64) error {
	TaskList[LastID+1] = &Task{ID: LastID + 1, Name: name, Reward: reward, Status: StatusToDo, CharacterID: 0}
	LastID += 1
	return commitToTaskFile()
}

func DeleteTask(ID int64) error {
	delete(TaskList, ID)
	return commitToTaskFile()
}

func DoTask(taskID int64) error {
	if task, ok := TaskList[taskID]; ok {
		if task.Status == StatusToDo && task.CharacterID == 0 {
			TaskList[taskID].Status = StatusDoing
			TaskList[taskID].CharacterID = CurrentCharacter()
			return commitToTaskFile()
		} else {
			return errors.New("该悬赏已被领取！")
		}
	} else {
		return errors.New("该悬赏不存在！")
	}
}

func UnDoTask(taskID int64) error {
	if task, ok := TaskList[taskID]; ok {
		if task.Status == StatusDoing {
			TaskList[taskID].Status = StatusToDo
			TaskList[taskID].CharacterID = 0
			return commitToTaskFile()
		}
		if task.Status == StatusToDo {
			return errors.New("该悬赏未被领取！")
		}
		if task.Status == StatusDone {
			return errors.New("该悬赏已被完成！")
		}
		return nil
	} else {
		return errors.New("该悬赏不存在！")
	}
}

func DoneTask(taskID int64) error {
	if task, ok := TaskList[taskID]; ok {
		if task.Status == StatusDoing {
			TaskList[taskID].Status = StatusDone
			if _, ok := Characters[task.CharacterID]; ok {
				err := Characters[task.CharacterID].UpReward(task.Reward)
				if err != nil {
					return err
				}
			}
			return commitToTaskFile()
		}
		if task.Status == StatusToDo {
			return errors.New("该悬赏未被领取！")
		}
		if task.Status == 3 {
			return errors.New("该悬赏已被完成！")
		}
		return nil
	} else {
		return errors.New("该悬赏不存在！")
	}
}

func commitToTaskFile() error {
	jsonString, err := json.Marshal(TaskList)
	if err != nil {
		return err
	}
	_ = helper.TasksFile.Truncate(0)
	_, _ = helper.TasksFile.Seek(0, 0)
	_, err = helper.TasksFile.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}

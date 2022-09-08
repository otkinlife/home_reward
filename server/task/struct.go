package task

const StatusToDo = 1
const StatusDoing = 2
const StatusDone = 3

type Task struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Reward      int64  `json:"reward"`
	Status      int64  `json:"status"`
	CharacterID int64  `json:"character_id"`
}

package wish

const StatusWant = 1
const StatusFinished = 2

type Wish struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Reward      int64  `json:"reward"`
	Status      int64  `json:"status"`
	CharacterID int64  `json:"character_id"`
	Publisher   int64  `json:"publisher"`
	Other       string `json:"other"`
}

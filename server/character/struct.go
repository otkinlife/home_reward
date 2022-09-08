package character

import "errors"

type Character struct {
	ID     int64  `json:"id"`
	Nick   string `json:"nick"`
	Reward int64  `json:"reward"`
	Avatar string `json:"avatar"`
}

func (c *Character) AddReward(reward int64) {
	c.Reward += reward
}

func (c *Character) Reduce(reward int64) error {
	if reward > c.Reward {
		return errors.New("余额不足")
	}
	c.Reward -= reward
	return nil
}

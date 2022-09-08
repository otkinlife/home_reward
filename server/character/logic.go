package character

func Current() *Character {
	c, err := One(1)
	if err != nil {
		return nil
	}
	return &c
}

func List() (map[int64]Character, error) {
	characters := map[int64]Character{}
	err := getCharacters(&characters)
	if err != nil {
		return characters, err
	}
	return characters, nil
}

func One(ID int64) (Character, error) {
	character := Character{}
	err := getCharacterByID(ID, &character)
	if err != nil {
		return character, err
	}
	return character, nil
}

func Create(nick string, avatar string) error {
	c := Character{
		Nick:   nick,
		Reward: 0,
		Avatar: avatar,
	}
	return save(c)
}

func AddReward(character Character, reward int64) error {
	character.AddReward(reward)
	return save(character)
}

func ReduceReward(character Character, reward int64) error {
	err := character.Reduce(reward)
	if err != nil {
		return err
	}
	return save(character)
}

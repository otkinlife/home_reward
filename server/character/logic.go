package character

type Logic struct {
	IP               string
	CurrentCharacter *Character
}

func NewLogic(ip string) *Logic {
	return &Logic{
		IP:               ip,
		CurrentCharacter: Current(ip),
	}
}

func Current(ip string) *Character {
	c := Character{}
	err := getCharacterByIp(ip, &c)
	if err != nil {
		return nil
	}
	return &c
}

func Bind(ip string, id int64) error {
	return bind(ip, id)
}

func UnBind(ip string, id int64) error {
	return unbind(ip, id)
}

func (l *Logic) List() (map[int64]Character, error) {
	characters := map[int64]Character{}
	err := getCharacters(&characters)
	if err != nil {
		return characters, err
	}
	return characters, nil
}

func (l *Logic) One(ID int64) (Character, error) {
	character := Character{}
	err := getCharacterByID(ID, &character)
	if err != nil {
		return character, err
	}
	return character, nil
}

func (l *Logic) Create(nick string, avatar string) error {
	c := Character{
		Nick:   nick,
		Reward: 0,
		Avatar: avatar,
	}
	return save(c)
}

func (l *Logic) AddReward(character *Character, reward int64) error {
	character.AddReward(reward)
	return save(*character)
}

func (l *Logic) ReduceReward(character *Character, reward int64) error {
	err := character.Reduce(reward)
	if err != nil {
		return err
	}
	return save(*character)
}

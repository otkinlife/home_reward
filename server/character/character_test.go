package character

import (
	"fmt"
	"home-reward/server/base"
	"home-reward/server/config"
	"testing"
)

func TestCharacter(t *testing.T) {
	config.InitConfig()
	base.InitMySQL()
	create(t)
	list := list(t)
	if len(list) == 0 {
		return
	}
	c := Character{}
	for _, v := range list {
		c = v
		break
	}
	add(t, c)
	reduce(t, c)
}

func create(t *testing.T) {
	fmt.Println("test create...")
	err := Create("贾凯超", "")
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test create pass")
}

func list(t *testing.T) map[int64]Character {
	fmt.Println("test list...")
	data, err := List()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(data)
	fmt.Println("test list pass")
	return data
}

func add(t *testing.T, c Character) {
	fmt.Println("test add...")
	err := AddReward(c, 100)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println("test add pass")
}

func reduce(t *testing.T, c Character) {
	fmt.Println("test add...")
	err := ReduceReward(c, 10)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println("test add pass")
}

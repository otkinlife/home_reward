package character

import (
	"fmt"
	"home-reward/server/base"
	"home-reward/server/config"
	"testing"
)

var logic *Logic

func TestCharacter(t *testing.T) {
	config.InitConfig()
	base.InitMySQL()
	logic = NewLogic("1.1.1.1")
	fmt.Println(logic.CurrentCharacter)
	//create(t)
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
	err := logic.Create("狗蛋", "")
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test create pass")
}

func list(t *testing.T) map[int64]Character {
	fmt.Println("test list...")
	data, err := logic.List()
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
	err := logic.AddReward(c, 100)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println("test add pass")
}

func reduce(t *testing.T, c Character) {
	fmt.Println("test add...")
	err := logic.ReduceReward(c, 10)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println("test add pass")
}

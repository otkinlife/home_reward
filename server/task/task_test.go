package task

import (
	"fmt"
	"home-reward/server/base"
	"home-reward/server/config"
	"testing"
)

var logic *Logic

func TestTask(t *testing.T) {
	config.InitConfig()
	base.InitMySQL()

	logic = NewLogic("1.1.1.1")
	fmt.Println(logic.CurrentCharacter)
	create(t)
	list := list(t)
	if len(list) == 0 {
		return
	}
	id := int64(0)
	for k, _ := range list {
		id = k
		break
	}
	get(t, id)
	finish(t, id)
	cancelFinish(t, id)
	cancelGet(t, id)
	del(t, id)
}

func create(t *testing.T) {
	fmt.Println("test create...")
	err := logic.Create("拖地", 100)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test create pass")
}

func list(t *testing.T) map[int64]Task {
	fmt.Println("test list...")
	list, err := logic.List()
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println(list)
	fmt.Println("test list pass")
	return list
}

func get(t *testing.T, id int64) {
	fmt.Println("test get ...")
	err := logic.Get(id)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test get pass")
}

func cancelGet(t *testing.T, id int64) {
	fmt.Println("test cancel get ...")
	err := logic.CancelGet(id)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test cancel get pass")
}

func finish(t *testing.T, id int64) {
	fmt.Println("test finish ...")
	err := logic.Finish(id)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test finish pass")
}

func cancelFinish(t *testing.T, id int64) {
	fmt.Println("test cancel finish ...")
	err := logic.CancelFinish(id)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test cancel finish pass")
}

func del(t *testing.T, id int64) {
	fmt.Println("test del ...")
	err := logic.Delete(id)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
	fmt.Println("test del pass")
}

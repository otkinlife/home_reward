package api

import (
	"encoding/json"
	"fmt"
	"home-reward/server/task"
	"net/http"
	"strconv"
)

func GetTaskList(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	data, err := task.List()
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = data
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	name := r.URL.Query().Get("name")
	if name == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("任务名不存在！")
		return
	}
	reward := r.URL.Query().Get("reward")
	if reward == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("赏金不存在！")
		return
	}
	rewardNumber, err := strconv.Atoi(reward)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = task.Create(name, int64(rewardNumber))
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "任务发布成功"
}

func DelTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("任务不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = task.Delete(int64(ID))
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "任务删除成功"
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("任务不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}

	err = task.Get(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "任务领取成功"
}

func CancelGetTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	TaskIDString := r.URL.Query().Get("id")
	if TaskIDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("任务不存在！")
		return
	}
	TaskID, err := strconv.Atoi(TaskIDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = task.CancelGet(int64(TaskID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "任务取消领取成功"
}

func FinishTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	TaskIDString := r.URL.Query().Get("id")
	if TaskIDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("任务不存在！")
		return
	}
	TaskID, err := strconv.Atoi(TaskIDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = task.Finish(int64(TaskID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}

	resp.Data = "任务完成"
}

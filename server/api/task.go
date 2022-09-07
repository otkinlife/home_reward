package api

import (
	"encoding/json"
	"fmt"
	"home-reward/server/object"
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
	data := map[int64][]*object.Task{}
	for _, task := range object.TaskList {
		data[task.Status] = append(data[task.Status], task)
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
		resp.Data = fmt.Sprint("悬赏名不存在！")
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
	err = object.CreateTask(name, int64(rewardNumber))
	if err != nil {
		resp.Data = "悬赏发布成功，但文件未更新"
		return
	}
	resp.Data = "悬赏发布成功"
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
		resp.Data = fmt.Sprint("悬赏不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = object.DeleteTask(int64(ID))
	if err != nil {
		resp.Data = "悬赏删除成功，但文件未更新"
		return
	}
	resp.Data = "悬赏删除成功"
}

func DoTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("悬赏不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}

	err = object.DoTask(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "悬赏领取成功"
}

func UnDoTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	TaskIDString := r.URL.Query().Get("id")
	if TaskIDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("悬赏不存在！")
		return
	}
	TaskID, err := strconv.Atoi(TaskIDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = object.UnDoTask(int64(TaskID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "悬赏取消领取成功"
}

func DoneTask(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	TaskIDString := r.URL.Query().Get("id")
	if TaskIDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("悬赏不存在！")
		return
	}
	TaskID, err := strconv.Atoi(TaskIDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = object.DoneTask(int64(TaskID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}

	resp.Data = "悬赏完成"
}

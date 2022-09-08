package api

import (
	"encoding/json"
	"fmt"
	"home-reward/server/wish"
	"net/http"
	"strconv"
)

func GetWishList(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()

	ip := getClientIP(r)
	logic := wish.NewLogic(ip)
	if logic.CurrentCharacter == nil {
		resp.ErrNo = 2
		resp.Data = "该设备还未绑定角色"
		return
	}

	data, err := logic.List()
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = data
}

func CreateWish(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()

	ip := getClientIP(r)
	logic := wish.NewLogic(ip)
	if logic.CurrentCharacter == nil {
		resp.ErrNo = 2
		resp.Data = "该设备还未绑定角色"
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("愿望名不存在！")
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
	other := r.URL.Query().Get("other")
	err = logic.Create(name, int64(rewardNumber), other)
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "愿望发布成功"
}

func DelWish(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()

	ip := getClientIP(r)
	logic := wish.NewLogic(ip)
	if logic.CurrentCharacter == nil {
		resp.ErrNo = 2
		resp.Data = "该设备还未绑定角色"
		return
	}

	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("愿望不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = logic.Delete(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "愿望删除成功"
}

func FinishWish(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()

	ip := getClientIP(r)
	logic := wish.NewLogic(ip)
	if logic.CurrentCharacter == nil {
		resp.ErrNo = 2
		resp.Data = "该设备还未绑定角色"
		return
	}

	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("愿望不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}

	err = logic.Finish(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "愿望购买成功"
}

func CancelFinishWish(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()

	ip := getClientIP(r)
	logic := wish.NewLogic(ip)
	if logic.CurrentCharacter == nil {
		resp.ErrNo = 2
		resp.Data = "该设备还未绑定角色"
		return
	}

	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("愿望不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = logic.Finish(int64(ID))
	if err != nil {
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "愿望取消完成成功"
}

package api

import (
	"encoding/json"
	"fmt"
	"home-reward/server/character"
	"net/http"
	"strconv"
)

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	ip := getClientIP(r)
	logic := character.NewLogic(ip)
	if logic.CurrentCharacter == nil {
		resp.ErrNo = 2
		resp.Data = "该设备还未绑定角色"
		return
	}
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("角色不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	data, err := logic.One(int64(ID))
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = data
}

func BindCharacterWithIP(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	ip := getClientIP(r)
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("角色不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = character.Bind(ip, int64(ID))
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "角色绑定设备成功"
}

func UnBindCharacterWithIP(w http.ResponseWriter, r *http.Request) {
	resp := new(Resp)
	defer func() {
		setupCORS(&w)
		resString, _ := json.Marshal(resp)
		_, _ = w.Write(resString)
	}()
	ip := getClientIP(r)
	IDString := r.URL.Query().Get("id")
	if IDString == "" {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint("角色不存在！")
		return
	}
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	err = character.UnBind(ip, int64(ID))
	if err != nil {
		resp.ErrNo = 1
		resp.Data = fmt.Sprint(err)
		return
	}
	resp.Data = "角色解绑设备成功"
}

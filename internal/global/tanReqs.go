// internal/global/tanReqs.go

package global

import (
	"OpenTan/config"
	"OpenTan/internal/global/model"
	"OpenTan/internal/global/response"
	"OpenTan/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func Login(id, password string) (string, error) {
	body := `{"identity":"` + id + `","password":"` + password + `","remember":true}`
	req, err := utils.NewTanPostRequest(model.API_BASE+"/users/login", utils.JsonString2Body(body))
	if err != nil {
		return "", err
	}
	utils.AddHeader(req, "accept", "application/json, text/plain, */*")
	utils.AddHeader(req, "content-type", "application/json")
	utils.AddHeader(req, "origin", "https://mytan.maiseed.com.cn")
	utils.AddHeader(req, "referer", "https://mytan.maiseed.com.cn/login")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var msg model.LoginMsg
	err = json.Unmarshal(respBody, &msg)
	if err != nil {
		return "", err
	}
	return msg.Data.Token.Token, nil
}

func GetModels() func(c *gin.Context) {
	req, err := utils.NewTanGetRequest(model.API_BASE + "/models")
	if err != nil {
		return response.NewServerError(500, "Internal Server Error")
	}
	utils.AddHeader(req, "accept", "application/json, text/plain, */*")
	utils.AddHeader(req, "authorization", config.Get().API_KEY)
	utils.AddHeader(req, "priority", "u=1, i")
	utils.AddHeader(req, "referer", "https://mytan.maiseed.com.cn/chat")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response.NewServerError(500, "Internal Server Error")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.NewServerError(500, "Internal Server Error")
	}

	return func(c *gin.Context) {
		c.Data(200, "application/json", respBody)
	}
}

const (
	NeedRefCode = "user_forced_login_limit"
	NeedRefMsg  = "登录设备发生变更。为保障账号安全，请重新登录。"
)

func TryRefresh() bool {
	req, err := utils.NewTanGetRequest(model.API_BASE + "/models")
	if err != nil {
		utils.PanicOnErr(err)
	}
	utils.AddHeader(req, "accept", "application/json, text/plain, */*")
	utils.AddHeader(req, "authorization", config.Get().API_KEY)
	utils.AddHeader(req, "priority", "u=1, i")
	utils.AddHeader(req, "referer", "https://mytan.maiseed.com.cn/chat")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.PanicOnErr(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.PanicOnErr(err)
	}

	var msg model.NeedRefreshMsg
	_ = json.Unmarshal(respBody, &msg)
	if len(msg.Errors) < 1 {
		return false
	}
	if msg.Errors[0].Code == NeedRefCode && msg.Errors[0].Message == NeedRefMsg {
		c := config.Get()
		token, err := Login(c.ID, c.Password)
		if err != nil {
			utils.PanicOnErr(err)
		}
		log.Println("Token before: ", c.API_KEY)
		c.API_KEY = "Bearer " + token
		config.Set(c)
		log.Println("Token after: ", c.API_KEY)
		return true
	} else {
		return false
	}
}

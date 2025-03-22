// internal/global/model/loginMsg.go

package model

type LoginMsg struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
	Data    struct {
		User struct {
			ID            string `json:"_id"`
			IDInt         int    `json:"id"`
			Name          string `json:"name"`
			Avatar        string `json:"avatar"`
			Email         string `json:"email"`
			EmailVerified bool   `json:"email_verified"`
			Phone         string `json:"phone"`
			PhoneVerified bool   `json:"phone_verified"`
			PasswordBind  bool   `json:"password_bind"`
			Gender        string `json:"gender"`
			CreatedTime   string `json:"created_time"`
		} `json:"user"`
		Token struct {
			Token      string `json:"token"`
			ExpireTime string `json:"expire_time"`
		} `json:"token"`
	} `json:"data"`
}

type NeedRefreshMsg struct {
	Success bool `json:"success"`
	Errors  []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Data map[string]interface{} `json:"data"`
}

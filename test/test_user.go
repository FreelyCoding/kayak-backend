package test

import (
	"github.com/go-playground/assert/v2"
	"kayak-backend/api"
	"net/http"
	"testing"
)

func testUserInfo(t *testing.T) {
	loginRes := api.LoginResponse{}
	code := Post("/login", "", &api.LoginInfo{
		UserName: initUser[4].Name,
		Password: initUser[4].Password,
	}, &loginRes)
	assert.Equal(t, code, http.StatusOK)
	assert.NotEqual(t, loginRes.Token, "")

	token := loginRes.Token
	userInfoRes := api.UserInfoResponse{}
	code = Get("/user/info", token, map[string][]string{}, &userInfoRes)
	//fmt.Println(time.Now().Local())
	//fmt.Println(userInfoRes.CreateAt)
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, userInfoRes.UserId, 5)
	assert.Equal(t, userInfoRes.UserName, initUser[4].Name)
	assert.Equal(t, userInfoRes.Email, initUser[4].Email)
	assert.Equal(t, userInfoRes.Phone, initUser[4].Phone)
}

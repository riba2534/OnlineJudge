package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riba2534/OnlineJudge/backend/models"
	"github.com/riba2534/OnlineJudge/backend/pkg/app"
	"github.com/riba2534/OnlineJudge/backend/pkg/e"
	"github.com/riba2534/OnlineJudge/backend/pkg/util"
)

func Login(c *gin.Context) {
	var user models.User
	appG := app.Gin{C: c}
	if err := c.ShouldBind(&user); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	username := user.UserName
	password := user.PassWord
	if username == "" || password == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	isExist, err := models.CheckAuth(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	// 生成token
	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	var user models.User
	// 绑定获取json参数
	if err := c.ShouldBind(&user); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	userName := user.UserName
	if userName == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	isExist, err := models.CheckUser(userName)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	// 用户已经存在
	if isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_EXIST_USER, nil)
		return
	}
	// 创建用户
	err = models.CreateUser(user)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func UserData(c *gin.Context) {
	appG := app.Gin{C: c}
	var userDataParam models.UserData
	if err := c.ShouldBind(&userDataParam); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	userName := userDataParam.UserName
	if userName == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	userData, err := models.GetUserDataByUserName(userName)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, userData)
}

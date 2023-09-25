package controller

import (
	"encoding/json"
	"fmt"
	"github.com/RaymondCode/simple-demo/form"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
)

type IUserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	UserInfo(c *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(service service.IUserService) IUserController {
	return &UserController{service}
}

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.UserResp `json:"user"`
}

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (u *UserController) Register(c *gin.Context) {
	// 1.解析参数
	zap.L().Info("Register controller")
	username := c.Query("username")
	rawpassword := c.Query("password")
	password := utils.PwdEncode(rawpassword)
	fmt.Println("Register controller", password)
	userId, err := u.userService.Register(username, password)
	if userId <= 0 || err != nil {
		//zap.L().Info(err.Error())
		zap.L().Info("Register controller" + err.Error())
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		token, err := utils.GenToken(username, password)
		if err != nil {
			log.Println("GenToken 出错了", err)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "服务器出错了"},
			})
		}
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 0},
			UserId:   userId,
			Token:    token,
		})
	}
}

func (u *UserController) Login(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")
	//zap.L().Info("Login controller" + username + "-" + password)
	//if username == "" || password == "" {
	//	zap.L().Debug("Login controller 用户名和密码不能为空")
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "用户名和密码不能为空"},
	//	})
	//}
	var loginForm form.Login
	err := c.ShouldBind(&loginForm)
	if err != nil {
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
			})
		}
		zap.L().Debug("Login controller 用户名或密码格式错误" + err.Error())
		dataType, _ := json.Marshal(err.Translate(global.Trans))
		errString := string(dataType)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: errString},
		})
	}
	userId, err := u.userService.Login(loginForm.Username, loginForm.Password)
	fmt.Println("Login", userId)
	if userId <= 0 || err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		token, err := utils.GenToken(loginForm.Username, loginForm.Password)
		if err != nil {
			log.Println("GenToken 出错了", err)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "服务器出错了"},
			})
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userId,
			Token:    token,
		})
	}
}

func (u *UserController) UserInfo(c *gin.Context) {
	//token := c.Query("token")
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	user, err := u.userService.User(userId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *user,
		})
	}
}

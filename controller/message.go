package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IMessageController interface {
	MessageAction(c *gin.Context)
	MessageChat(c *gin.Context)
}

type MessageController struct {
	message service.IMessageService
}

func NewRMessageController(message service.IMessageService) IMessageController {
	return &MessageController{message}
}

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []model.MessageResp `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func (m *MessageController) MessageAction(c *gin.Context) {
	toUserIdStr := c.Query("to_user_id")
	content := c.Query("content")
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		log.Println("toUserId 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "toUserId 格式错误"})
	}
	err = m.message.MessageAction(username.(string), password.(string), content, toUserId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// MessageChat all users have same follow list
func (m *MessageController) MessageChat(c *gin.Context) {
	toUserIdStr := c.Query("to_user_id")
	preMsgTimeStr := c.Query("pre_msg_time")
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	//fmt.Println(username, password)
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		log.Println("toUserId 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "toUserId 格式错误"})
	}
	preMsgTime, _ := strconv.ParseInt(preMsgTimeStr, 10, 64)
	msgs, err := m.message.MessageChat(username.(string), password.(string), preMsgTimeStr, toUserId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "MessageAction 出错了"})
	} else {
		// 1.preMsgTime = 0，获取全部聊天记录
		if preMsgTime == 0 {
			c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: msgs})
		} else if len(msgs) == 0 {
			c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}})
		} else {
			i := len(msgs) - 1
			for ; i > 0; i-- {
				if msgs[i].CreateTime == preMsgTime {
					break
				}
			}
			c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: msgs[i+1:]})
		}
		// 2.已经获取过所有聊天记录了，追加
		// 3.正在聊天且最新的一条消息是我发的，不追加
	}

	//else if preMsgTime == msgs[len(msgs)-1].CreateTime || (preMsgTime != 0 && msgs[len(msgs) - 1].ToUserId == toUserId){
	//	//fmt.Println("MessageChat", msgs)
	//	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}})
	//} else {
	//	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: msgs})
	//}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}

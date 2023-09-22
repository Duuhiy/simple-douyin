package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/repository/redis"
	"log"
	"time"
)

//
type IMessageService interface {
	MessageAction(username, password, content string, toUserId int64) error
	MessageChat(username, password string, toUserId int64) ([]model.MessageResp, error)
}

type MessageService struct {
	db  mysql.IUserRepository
	rdb redis.IRedis
}

func (m *MessageService) MessageAction(username, password, content string, toUserId int64) error {
	//TODO implement me
	// 把消息插入数据库
	userId, _ := m.db.FindOneByToken(username, password)
	message := model.Message{
		ToUserId:   toUserId,
		FromUserId: userId,
		Content:    content,
		CreateAt:   time.Now(),
	}
	err := m.db.MessageInsert(&message)
	return err
}

func (m *MessageService) MessageChat(username, password string, toUserId int64) ([]model.MessageResp, error) {
	//TODO implement me
	//log.Println("MessageChat")
	userId, _ := m.db.FindOneByToken(username, password)
	//fmt.Println(user)
	msgs, err := m.db.MessageFindByUserToUser(userId, toUserId)
	var msgResp []model.MessageResp
	for _, msg := range msgs {
		//fmt.Println(msg.CreateAt)
		//createTime, err := strconv.ParseInt(msg.CreateAt.Format("2006-01-02 03:04:05"), 10, 64)
		if err != nil {
			log.Println("MessageChat", err)
		}
		msgr := model.MessageResp{
			Id:         msg.Id,
			ToUserId:   msg.ToUserId,
			FromUserId: msg.FromUserId,
			Content:    msg.Content,
			CreateTime: msg.CreateAt.Unix(),
		}
		msgResp = append(msgResp, msgr)
	}
	//fmt.Println(msgResp[0].CreateTime)
	return msgResp, err
}

func NewMessageService(db mysql.IUserRepository, rdb redis.IRedis) IMessageService {
	return &MessageService{db, rdb}
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

//var chatConnMap = sync.Map{}

//func RunMessageServer() {
//	listen, err := net.Listen("tcp", "127.0.0.1:8888")
//	if err != nil {
//		fmt.Printf("Run message sever failed: %v\n", err)
//		return
//	}
//
//	for {
//		conn, err := listen.Accept()
//		if err != nil {
//			fmt.Printf("Accept conn failed: %v\n", err)
//			continue
//		}
//
//		go process(conn)
//	}
//}
//
//func process(conn net.Conn) {
//	defer conn.Close()
//
//	var buf [256]byte
//	for {
//		n, err := conn.Read(buf[:])
//		if n == 0 {
//			if err == io.EOF {
//				break
//			}
//			fmt.Printf("Read message failed: %v\n", err)
//			continue
//		}
//
//		var event = MessageSendEvent{}
//		_ = json.Unmarshal(buf[:n], &event)
//		fmt.Printf("Receive Message：%+v\n", event)
//
//		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
//		if len(event.MsgContent) == 0 {
//			chatConnMap.Store(fromChatKey, conn)
//			continue
//		}
//
//		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
//		writeConn, exist := chatConnMap.Load(toChatKey)
//		if !exist {
//			fmt.Printf("User %d offline\n", event.ToUserId)
//			continue
//		}
//
//		pushEvent := MessagePushEvent{
//			FromUserId: event.UserId,
//			MsgContent: event.MsgContent,
//		}
//		pushData, _ := json.Marshal(pushEvent)
//		_, err = writeConn.(net.Conn).Write(pushData)
//		if err != nil {
//			fmt.Printf("Push message failed: %v\n", err)
//		}
//	}
//}

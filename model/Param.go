package model

type DouyinUserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DouyinUserRegisterResp struct {
	Status_code int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	Status_msg  string `json:"status_msg"`  // 返回状态描述
	User_id     int64  `json:"user_id"`     // 用户id
	Token       string `json:"token"`       // 用户鉴权token
}

type DouyinUserLoginReq struct {
	Username string `json:"username"` // 登录用户名
	Password string `json:"password"` // 登录密码
}

type DouyinUserLoginResp struct {
	Status_code int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	Status_msg  string `json:"status_msg"`  // 返回状态描述
	User_id     int64  `json:"user_id"`     // 用户id
	Token       string `json:"token"`       // 用户鉴权token
}

type DouyinUserReq struct {
	User_id int64  `json:"user_id"` // 用户id
	Token   string `json:"token"`   // 用户鉴权token
}

type DouyinUserResp struct {
	Status_code int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	Status_msg  string `json:"status_msg"`  // 返回状态描述
	User        User   `json:"user"`
}

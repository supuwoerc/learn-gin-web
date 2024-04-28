package response

var code2message = map[int]string{
	SUCCESS:        "操作成功",
	ERROR:          "操作失败",
	INVALID_PARAMS: "参数错误",
	INVALID_TOKEN:  "鉴权失败",
	UNKNOWN_ERROR:  "未知错误",
	RECOVERY_ERROR: "严重错误，请联系管理员",
}

func GetMessage(code int) string {
	if msg, exist := code2message[code]; exist {
		return msg
	}
	return code2message[UNKNOWN_ERROR]
}

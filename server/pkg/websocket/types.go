package websocket

const (
	KindJoin    = "Join"
	KindLeft    = "Left"
	KindMessage = "Message"
	KindTyping  = "Typing"
)

type ISocketEvent struct {
	Kind    string      `json:"kind"`
	Payload interface{} `json:"payload"`
}

type IUserJoinData struct {
	Kind   string `json:"kind"`
	UserId string `json:"user_id"`
}

type IUserLeaveData struct {
	Kind   string `json:"kind"`
	UserId string `json:"user_id"`
}

type IUserTypingData struct {
	Kind    string `json:"kind"`
	OUserId string `json:"user_id"`
}

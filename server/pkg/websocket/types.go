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

type IUserTypeData struct {
	Kind    string `json:"kind"`
	OuserId string `json:"user_id"`
}

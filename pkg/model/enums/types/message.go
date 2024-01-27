package types

type MessageType int32

const (
	MessageType__UNKNOWN MessageType = iota
	//
	MessageType__Hi

	// 文字
	MessageType__Text

	// 系統訊息
	MessageType__SystemMsg
)

type MessageStatus int32

const (
	MessageStatus__UNKNOWN MessageStatus = iota
	MessageStatus__Initial
	MessageStatus__Recycle
	MessageStatus__Edit
)

type ContentType int32

const (
	ContentType__UNKNOWN ContentType = iota
	ContentType__Hi
	ContentType__Text
)

// 訊息發送者標籤
type MsgSenderTagType int

const (
	//
	MsgSenderTagType_Unknown MsgSenderTagType = iota
	// 自動消息
	MsgSenderTagType_AutoReply
)

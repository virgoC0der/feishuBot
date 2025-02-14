package messages

import "encoding/json"

const (
	EventTypeReceiveMsg = "im.message.receive_v1"
)

type Verification struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
}

type LarkEvent struct {
	Schema string `json:"schema"`
	Header struct {
		EventId    string `json:"event_id"`
		Token      string `json:"token"`
		CreateTime string `json:"create_time"`
		EventType  string `json:"event_type"`
		TenantKey  string `json:"tenant_key"`
		AppId      string `json:"app_id"`
	} `json:"header"`
	Event json.RawMessage `json:"event"`
}

type ReceiveMsgV2 struct {
	Sender struct {
		SenderId struct {
			UnionId string `json:"union_id"`
			UserId  string `json:"user_id"`
			OpenId  string `json:"open_id"`
		} `json:"sender_id"`
		SenderType string `json:"sender_type"`
		TenantKey  string `json:"tenant_key"`
	} `json:"sender"`
	Message struct {
		MessageId   string `json:"message_id"`
		RootId      string `json:"root_id"`
		ParentId    string `json:"parent_id"`
		CreateTime  string `json:"create_time"`
		UpdateTime  string `json:"update_time"`
		ChatId      string `json:"chat_id"`
		ThreadId    string `json:"thread_id"`
		ChatType    string `json:"chat_type"`
		MessageType string `json:"message_type"`
		Content     string `json:"content"`
		Mentions    []struct {
			Key string `json:"key"`
			Id  struct {
				UnionId string `json:"union_id"`
				UserId  string `json:"user_id"`
				OpenId  string `json:"open_id"`
			} `json:"id"`
			Name      string `json:"name"`
			TenantKey string `json:"tenant_key"`
		} `json:"mentions"`
		UserAgent string `json:"user_agent"`
	} `json:"message"`
}

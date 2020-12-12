package webhook

import "fmt"

type WebhookEventEnum int

const (
	GetAllUsersEvent WebhookEventEnum = iota
	GetUserEvent
	RegisterUserEvent
	DeleteUserEvent
	GetCurrentUsersEvent
	PushCurrentUserEvent
	DeleteCurrentUserEvent
)

func (e WebhookEventEnum) String() string {
	return toString[e]
}

var toString = map[WebhookEventEnum]string{
	GetAllUsersEvent:       "GetAllUsersEvent",
	GetUserEvent:           "GetUserEvent",
	RegisterUserEvent:      "RegisterUserEvent",
	DeleteUserEvent:        "DeleteUserEvent",
	GetCurrentUsersEvent:   "GetCurrentUsersEvent",
	PushCurrentUserEvent:   "PushCurrentUserEvent",
	DeleteCurrentUserEvent: "DeleteCurrentUserEvent",
}

var toID = map[string]WebhookEventEnum{
	"GetAllUsersEvent":       GetAllUsersEvent,
	"GetUserEvent":           GetUserEvent,
	"RegisterUserEvent":      RegisterUserEvent,
	"DeleteUserEvent":        DeleteUserEvent,
	"GetCurrentUsersEvent":   GetCurrentUsersEvent,
	"PushCurrentUserEvent":   PushCurrentUserEvent,
	"DeleteCurrentUserEvent": DeleteCurrentUserEvent,
}

func (e *WebhookEventEnum) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data interface{}
	if err := unmarshal(&data); err != nil {
		return err
	}

	if wee, ok := toID[data.(string)]; ok {
		*e = wee
	} else {
		return fmt.Errorf("invalid WebhookEventEnum value '%s'", data.(string))
	}

	return nil
}

type WebhookRule struct {
	Name   string                 `yaml:"name"`
	Event  WebhookEventEnum       `yaml:"event"`
	Url    string                 `yaml:"url"`
	Params map[string]interface{} `yaml:"params"`
}

var WebhookRules []WebhookRule

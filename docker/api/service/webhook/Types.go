package webhook

import "fmt"

type WebhookEventEnum int

const (
	GetAllUsersEventEnum WebhookEventEnum = iota
	GetUserEventEnum
	RegisterUserEventEnum
	DeleteUserEventEnum
	GetCurrentUsersEventEnum
	PushCurrentUserEventEnum
	DeleteCurrentUserEventEnum
)

func (e WebhookEventEnum) String() string {
	return toString[e]
}

var toString = map[WebhookEventEnum]string{
	GetAllUsersEventEnum:       "GetAllUsersEvent",
	GetUserEventEnum:           "GetUserEvent",
	RegisterUserEventEnum:      "RegisterUserEvent",
	DeleteUserEventEnum:        "DeleteUserEvent",
	GetCurrentUsersEventEnum:   "GetCurrentUsersEvent",
	PushCurrentUserEventEnum:   "PushCurrentUserEvent",
	DeleteCurrentUserEventEnum: "DeleteCurrentUserEvent",
}

var toID = map[string]WebhookEventEnum{
	"GetAllUsersEvent":       GetAllUsersEventEnum,
	"GetUserEvent":           GetUserEventEnum,
	"RegisterUserEvent":      RegisterUserEventEnum,
	"DeleteUserEvent":        DeleteUserEventEnum,
	"GetCurrentUsersEvent":   GetCurrentUsersEventEnum,
	"PushCurrentUserEvent":   PushCurrentUserEventEnum,
	"DeleteCurrentUserEvent": DeleteCurrentUserEventEnum,
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

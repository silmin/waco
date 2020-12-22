package webhook

import "fmt"

type WebhookRule struct {
	Name   string                 `yaml:"name"`
	Event  WebhookEventEnum       `yaml:"event"`
	Url    string                 `yaml:"url"`
	Method WebhookMethodEnum      `yaml:"method"`
	Params map[string]interface{} `yaml:"params"`
}

var WebhookRules []WebhookRule

type WebhookEventEnum int

const (
	RegisterUserEvent WebhookEventEnum = iota
	DeleteUserEvent
	PushCurrentUserEvent
	PopCurrentUserEvent
)

func (e WebhookEventEnum) String() string {
	return toEventString[e]
}

var toEventString = map[WebhookEventEnum]string{
	RegisterUserEvent:    "RegisterUserEvent",
	DeleteUserEvent:      "DeleteUserEvent",
	PushCurrentUserEvent: "PushCurrentUserEvent",
	PopCurrentUserEvent:  "PopCurrentUserEvent",
}

var toEventID = map[string]WebhookEventEnum{
	"RegisterUserEvent":    RegisterUserEvent,
	"DeleteUserEvent":      DeleteUserEvent,
	"PushCurrentUserEvent": PushCurrentUserEvent,
	"PopCurrentUserEvent":  PopCurrentUserEvent,
}

func (e *WebhookEventEnum) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data interface{}
	if err := unmarshal(&data); err != nil {
		return err
	}

	if wee, ok := toEventID[data.(string)]; ok {
		*e = wee
	} else {
		return fmt.Errorf("invalid WebhookEventEnum value '%s'", data.(string))
	}

	return nil
}

type WebhookMethodEnum int

const (
	WebhookGET WebhookMethodEnum = iota
	WebhookPOST
)

func (e WebhookMethodEnum) String() string {
	return toMethodString[e]
}

var toMethodString = map[WebhookMethodEnum]string{
	WebhookGET:  "GET",
	WebhookPOST: "POST",
}

var toMethodID = map[string]WebhookMethodEnum{
	"GET":  WebhookGET,
	"POST": WebhookPOST,
}

func (e *WebhookMethodEnum) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data interface{}
	if err := unmarshal(&data); err != nil {
		return err
	}

	if wme, ok := toMethodID[data.(string)]; ok {
		*e = wme
	} else {
		return fmt.Errorf("invalid WebhookMethodEnum value '%s'", data.(string))
	}

	return nil
}

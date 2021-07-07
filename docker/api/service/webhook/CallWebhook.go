package webhook

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"eleuth/waco/service/room_user"
)

func CallWebhook(webhookEvent WebhookEventEnum, user room_user.User) {
	var werrs []WebhookError
	for _, rule := range WebhookRules {
		if rule.Event == webhookEvent {
			log.Println("Call Webhook :", rule.Name)
			switch rule.Method {
			case WebhookGET:
				werrs = append(werrs, CallWebhookGET(rule, user))
			case WebhookPOST:
				werrs = append(werrs, CallWebhookPOST(rule, user))
			}
		}
	}

	if len(werrs) != 0 {
		for _, werr := range werrs {
			if werr.Err != nil {
				log.Println("Webhook Error :", werr.Name, ":", werr.Err)
			}
		}
	}
}

func CallWebhookGET(rule WebhookRule, user room_user.User) WebhookError {
	url := markToUserParams(rule.Url, user)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return WebhookError{rule.Name, err}
	}

	values := req.URL.Query()
	for k, v := range rule.Params {
		key := markToUserParams(k, user)
		value := markToUserParams(v.(string), user)
		values.Set(key, value)
	}
	req.URL.RawQuery = values.Encode()

	if _, err := http.DefaultClient.Do(req); err != nil {
		return WebhookError{rule.Name, err}
	}

	return WebhookError{rule.Name, nil}
}

func CallWebhookPOST(rule WebhookRule, user room_user.User) WebhookError {
	values := make(url.Values)
	for k, v := range rule.Params {
		key := markToUserParams(k, user)
		value := markToUserParams(v.(string), user)
		values.Set(key, value)
	}

	url := markToUserParams(rule.Url, user)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))
	if err != nil {
		return WebhookError{rule.Name, err}
	}

	if _, err := http.DefaultClient.Do(req); err != nil {
		return WebhookError{rule.Name, err}
	}

	return WebhookError{rule.Name, nil}
}

func markToUserParams(s string, user room_user.User) string {
	s = strings.ReplaceAll(s, "<card_no>", user.CardNo)
	s = strings.ReplaceAll(s, "<display_name>", user.DisplayName)
	s = strings.ReplaceAll(s, "<full_name>", user.FullName)
	s = strings.ReplaceAll(s, "<pronunciation>", user.Pronunciation)
	s = strings.ReplaceAll(s, "<playlist>", user.Playlist)
	s = strings.ReplaceAll(s, "<email>", user.Email)
	return s
}

package webhook

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"eleuth/waco/service/room_user"
)

func CallWebhook(webhookEvent WebhookEventEnum, user room_user.User) {
	var errs []error
	for idx, rule := range WebhookRules {
		if rule.Event == webhookEvent {
			log.Println("Call Webhook", idx, ":", rule)
			switch rule.Method {
			case WebhookGET:
				errs = append(errs, CallWebhookGET(rule, user))
			case WebhookPOST:
				errs = append(errs, CallWebhookPOST(rule, user))
			}
		}
	}

	if len(errs) != 0 {
		for idx, err := range errs {
			if err != nil {
				log.Println("Webhook Error", idx, ":", err)
			}
		}
	}
}

func CallWebhookGET(rule WebhookRule, user room_user.User) error {
	url := markToUserParams(rule.Url, user)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	values := req.URL.Query()
	for k, v := range rule.Params {
		key := markToUserParams(k, user)
		value := markToUserParams(v.(string), user)
		values.Set(key, value)
	}
	req.URL.RawQuery = values.Encode()

	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	return nil
}

func CallWebhookPOST(rule WebhookRule, user room_user.User) error {
	values := make(url.Values)
	for k, v := range rule.Params {
		key := markToUserParams(k, user)
		value := markToUserParams(v.(string), user)
		values.Set(key, value)
	}

	url := markToUserParams(rule.Url, user)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	return nil
}

func markToUserParams(s string, user room_user.User) string {
	s = strings.ReplaceAll(s, "<card_no>", user.CardNo)
	s = strings.ReplaceAll(s, "<display_name>", user.DisplayName)
	s = strings.ReplaceAll(s, "<full_name>", user.FullName)
	s = strings.ReplaceAll(s, "<pronunciation>", user.Pronunciation)
	s = strings.ReplaceAll(s, "<email>", user.Email)
	return s
}

package webhook

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

func getRuleFilePaths(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			path, err := getRuleFilePaths(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, path...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, nil
}

func ImportWebhookRules() ([]WebhookRule, error) {
	log.Println("--import webhook config--")

	ruleFilePaths, err := getRuleFilePaths("/app/webhook_rules/")
	var webhookRules []WebhookRule

	for _, filename := range ruleFilePaths {
		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var rules []WebhookRule
		err = yaml.Unmarshal([]byte(buf), &rules)
		if err != nil {
			return nil, err
		}

		webhookRules = append(webhookRules, rules...)
	}

	return webhookRules, err
}

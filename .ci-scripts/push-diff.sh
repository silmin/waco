#!/bin/bash

git diff --exit-code $CI_COMMIT_SHA

if [ $? = 0 ]; then
    exit 0
else
    eval "$(ssh-agent -s)"
    echo "$SSH_PRIVATE_KEY" | ssh-add - > /dev/null

    git config --global user.name "horisan"
    git config --global user.email "horisan@example.com"

    git remote set-url --push origin git@$CI_SERVER_HOST:$CI_PROJECT_PATH.git
    git add .
    git commit -m "formatted by horisan with GitLab runner"
    git -c core.sshCommand="ssh -oStrictHostKeyChecking=no" push origin HEAD:$CI_COMMIT_REF_NAME

	if [ $? = 0 ]; then
		curl --request POST \
			--header "PRIVATE-TOKEN: $HORISAN_API_TOKEN" \
			"$CI_API_V4_URL/projects/$CI_PROJECT_ID/pipelines/$CI_PIPELINE_ID/cancel"
	fi
fi


#!/bin/sh

argocd login 192.168.10.21 --username "admin" --password "qwer1234" --insecure

argocd app get $CI_PROJECT_NAME > /dev/null 2>&1

if [ "$?" -ne "0" ]; then
	argocd app create $CI_PROJECT_NAME \
		--repo $CI_PROJECT_URL \
		--path .k8s \
		--dest-namespace $CI_PROJECT_NAME \
		--dest-server https://kubernetes.default.svc \
		--sync-policy automatic \
		--sync-option CreateNamespace=true \
		--kustomize-image $REMOTE_REGISTRY/$REGISTRY_REPO:$CI_COMMIT_SHA
else
	argocd app set $CI_PROJECT_NAME --kustomize-image $REMOTE_REGISTRY/$REGISTRY_REPO:$CI_COMMIT_SHA
fi


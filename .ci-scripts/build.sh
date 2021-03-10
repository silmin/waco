mkdir -p /kaniko/.docker

# docker build & push
/kaniko/executor --context $CI_PROJECT_DIR/docker/api/ \
	--dockerfile $CI_PROJECT_DIR/docker/api/Dockerfile \
	--destination $INTERNAL_REGISTRY/$REGISTRY_REPO:$CI_COMMIT_TAG \
	--destination $INTERNAL_REGISTRY/$REGISTRY_REPO:$CI_COMMIT_SHA \
	--insecure --skip-tls-verify


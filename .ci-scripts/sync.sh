#!/bin/bash

STATUS=`curl -u $HARBOR_USER_INFO -X POST http://beatrice.eleuth/api/v2.0/replication/executions -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"policy_id\": 1}" | tee curl_result | jq -r ". | length"`

cat ./curl_result

exit $STATUS


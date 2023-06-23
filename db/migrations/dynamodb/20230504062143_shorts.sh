#!/usr/bin/env sh
args="--table-name=shorts"
args="${args} --attribute-definitions AttributeName=key,AttributeType=S"
args="${args} --key-schema AttributeName=key,KeyType=HASH"
args="${args} --billing-mode PAY_PER_REQUEST"
endpoint=${AWS_ENDPOINT:-${AWS_DEFAULT_ENDPOINT}}

if [ "${endpoint}" != "" ]; then
  args="${args} --endpoint-url ${endpoint}"
fi;

result=$(aws dynamodb create-table $args)

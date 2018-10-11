#!/bin/bash

env | sort

waitforit -address=http://localstack/test:4572 -status=404 -timeout=30
aws --endpoint-url=http://localstack:4572 s3api create-bucket --bucket test

sam local start-api --docker-volume-basedir "${APPLICATION}/dist" --host 0.0.0.0 --template application/template.yaml --docker-network container_default

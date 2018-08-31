#!/bin/bash

env | sort
sam local start-api --docker-volume-basedir "${APPLICATION}/dist" --host 0.0.0.0 --template application/template.yaml

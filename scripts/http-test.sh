#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'

echo_msg() {
    echo
    echo
	echo -e "$GREEN============================== $1$NC"
}

echo_msg "Create a task"
curl -i -XPUT http://localhost:8000/v1/task -d "{\"title\": \"Update at $(date)\"}"

echo_msg "Get the task with ID 1"
curl -i -XGET http://localhost:8000/v1/task/1

echo_msg "Update the task with ID 1"
curl -i -XPUT http://localhost:8000/v1/task/1 -d "{\"title\": \"Update at $(date)\"}"

echo_msg "List all task"
curl    -XGET http://localhost:8000/v1/task | jq .

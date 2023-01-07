#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'

echo_msg() {
    echo
    echo
	echo -e "$GREEN============================== $1$NC"
}

echo_msg "Create a Todo"
curl -i -XPUT http://localhost:8000/v1/todo -d "{\"title\": \"Update at $(date)\"}"

echo_msg "Get the todo with ID 1"
curl -i -XGET http://localhost:8000/v1/todo/1

echo_msg "Update the todo with ID 1"
curl -i -XPUT http://localhost:8000/v1/todo/1 -d "{\"title\": \"Update at $(date)\"}"

echo_msg "List all Todo"
curl -i -XGET http://localhost:8000/v1/todo

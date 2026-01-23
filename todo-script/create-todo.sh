#!/bin/sh

# Exit immediately if a command fails
set -e

# Check required environment variables
if  [ -z "$TODO_URL" ]; then
  echo "Error: TODO_TEXT and TODO_URL environment variables must be set" >&2
  exit 1
fi
if [ -z "$TODO_TEXT" ]; then
  TODO_TEXT=$(curl -sI "https://en.wikipedia.org/wiki/Special:Random" \
  | awk -F': ' 'tolower($1)=="location" { print $2 }')
fi

# Send POST request with x-www-form-urlencoded data
response=$(curl -s -w "\n%{http_code}\n" -X POST \
    -H "Content-Type: application/x-www-form-urlencoded" \
    --data-urlencode "todo=$TODO_TEXT" \
    "$TODO_URL/todos")

http_body=$(echo "$response" | sed '$d')
http_code=$(echo "$response" | tail -n1)

if [ "$http_code" == "200" ]; then
	echo "Created todo: $http_body"
	exit 0
else
	echo "Failed to create todo: $http_body"
	exit 1
fi

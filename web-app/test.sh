#!/bin/sh

addr="localhost"

printf '%s\n' "[get]: json"
curl \
	--header "Content-Type: application/json" \
	--request GET \
	--data '{"username":"abc","password":"xyz", "age": 30}' \
	${addr}:1224/redis/val

printf '%s\n' "[post]: json"
curl \
	--header "Content-Type: application/json" \
	--request POST \
	--data '{"username":"abc","password":"xyz"}' \
	${addr}:1224/redis/val


printf '%s\n' "[get]: key"
curl \
	--request GET \
	${addr}:1224/redis/val

printf '%s\n' "[post]: key"
curl \
	--request POST \
	${addr}:1224/redis/val

echo


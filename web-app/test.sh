#!/bin/sh

port="9090"
addr="192.168.49.2:${port}"

# filestorage {{{
test_func() {
	section="${1}"
	arg="${2}"
	data="${3}"

	if [ -z "${data}" ] || [ "${data}" = "" ]; then
		data='{"car": "Mitsubishi", "tuner": true}'
	fi

	curl \
		--header "Content-Type: application/json" \
		--request GET \
		${addr}/${section}/${arg}

	curl \
		--header "Content-Type: application/json" \
		--request POST \
		--data "${data}" \
		${addr}/${section}/${arg}


	curl \
		--request GET \
		${addr}/${section}/${arg}

	curl \
		--request POST \
		${addr}/${section}/${arg}

}
# }}}

test_func "redis" "key-string" '{"data": "true", "type": "json", "testing": true}'

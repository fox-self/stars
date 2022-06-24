# shellcheck shell=bash

task.run() {
	go run .
}

task.run-python() {
	local token="$GITHUB_TOKEN"
	[ -z "$token" ] && token="$(<.env)"

	pipx run starred --username hyperupcall --token "$token" --sort "$@" > './README.md'
}

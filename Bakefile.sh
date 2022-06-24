# shellcheck shell=bash

task.run() {
	go run .
}

task.run-python() {
	token="$(<.env)"
	
	pipx run starred --username hyperupcall --token "$token" --sort "$@" > './README.md'
}

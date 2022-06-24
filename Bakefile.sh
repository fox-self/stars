# shellcheck shell=bash

task.run() {
	go run .
}

task.run-python() {
	pipx run starred --username hyperupcall --token "$(<.env)" --sort "$@" > './README.md'
}

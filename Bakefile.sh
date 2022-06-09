# shellcheck shell=bash

task.run() {
	starred --username hyperupcall --token "$TOKEN" --sort "$@" > './README.md'
}

task.rungo() {
	go run .
}
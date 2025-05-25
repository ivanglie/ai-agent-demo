include .env

.PHONY: run

run:
	DEEPSEEK_API_KEY=$(DEEPSEEK_API_KEY) go run .
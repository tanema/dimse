default: help

run: ## run the test program connecting to dicomserver.co.uk
	@go run cmd/main.go

lcl: ## run the test program with local orthanc. requires make up
	@go run cmd/main.go -host=localhost

gen: ## generate string methods for data types
	@go generate ./...

clean:
	@rm -rf ./tmp

up:
	@mkdir -p ./tmp
	@docker-compose \
		--file ./test/pacs/compose.yml \
		up -dy

down:
	@docker-compose \
		--file ./test/pacs/compose.yml \
		down

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+%?:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

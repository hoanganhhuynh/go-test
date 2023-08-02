BUILD_SCRIPTS_PATH	:= ./build/scripts

build:
	go build .
.PHONY: build

run:
	docker-compose up sdk
.PHONY: run

create-db:
	${BUILD_SCRIPTS_PATH}/create-db.sh
.PHONY: create-db
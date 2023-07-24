BUILD_SCRIPTS_PATH	:= ./build/scripts

build:
	${BUILD_SCRIPTS_PATH}/build.sh
.PHONY: build

run:
	docker-compose up
.PHONY: run

create-db:
	${BUILD_SCRIPTS_PATH}/create-db.sh
.PHONY: create-db
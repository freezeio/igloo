
IMGTAG ?= latest
REPONAME ?= freezevicente
IMAGE := ${REPONAME}/igloo:${IMGTAG}


all: build push

generate-dockerfile:
	go build -o bin/igloo
	bin/igloo

build: generate-dockerfile
	docker build -t ${IMAGE} .

push:
	docker push ${IMAGE}

run:
	docker run -it --rm ${IMAGE} bash

clean:
	rm -rf bin
	rm -f Dockerfile
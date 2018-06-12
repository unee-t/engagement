NAME=engagement
REPO=uneet/$(NAME)

all:
	go build

dev:
	@echo $$AWS_ACCESS_KEY_ID
	jq '.profile |= "uneet-dev" |.stages.staging |= (.domain = "e.dev.unee-t.com" | .zone = "dev.unee-t.com")' up.json.in > up.json
	up

demo:
	@echo $$AWS_ACCESS_KEY_ID
	jq '.profile |= "uneet-demo" |.stages.staging |= (.domain = "e.demo.unee-t.com" | .zone = "demo.unee-t.com")' up.json.in > up.json
	up

prod:
	@echo $$AWS_ACCESS_KEY_ID
	jq '.profile |= "uneet-prod" |.stages.staging |= (.domain = "e.unee-t.com" | .zone = "unee-t.com")' up.json.in > up.json
	up

clean:
	rm -f engagement gin-bin

build:
	docker build -t $(REPO) --build-arg COMMIT=$(shell git describe --always) .

start:
	docker run -d --name $(NAME) -p 9000:9000 $(REPO)

stop:
	docker stop $(NAME)
	docker rm $(NAME)

sh:
	docker exec -it $(NAME) /bin/sh

.PHONY: dev demo prod

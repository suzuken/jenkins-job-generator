.PHONY: test jenkins browse ssh

test:
	go test -v ./...

# launch jenkins container for testing purpose
jenkins:
	docker-compose up -d

HOST:=$(shell docker-machine ip $(shell docker-machine active))
browse:
	open http://$(HOST):8080

ssh:
	docker exec -it $(shell docker-compose ps -q) bash

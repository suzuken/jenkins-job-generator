test:
	go test -v

# launch jenkins container for testing purpose
jenkins: .jenkins

.jenkins:
	docker run -p 8080:8080 -p 50000:50000 jenkins && touch .jenkins&

HOST:=$(shell docker-machine ip $(shell docker-machine active))
browse:
	open http://$(HOST):8080

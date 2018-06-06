build:
	dep ensure

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/events/create/main.go
	mkdir -p bin/events
	zip bin/events/create.zip main
	mv main bin/events/create

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/subscriptions/create/main.go
	mkdir -p bin/subscriptions
	zip bin/subscriptions/create.zip main
	mv main bin/subscriptions/create

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/subscriptions/delete/main.go
	mkdir -p bin/subscriptions
	zip bin/subscriptions/delete.zip main
	mv main bin/subscriptions/delete
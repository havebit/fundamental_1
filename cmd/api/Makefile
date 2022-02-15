run:
	go run main.go
test:
	go test ./... -v
start-db:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=myapp -d postgres
image:
	docker build -t todo:test -f Dockerfile .
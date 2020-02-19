build:
	GOOS=linux go build -ldflags="-s -w" -o bin/serverless main.go

deploy: clean build
	sls deploy --verbose

clean:
	rm -rf ./bin
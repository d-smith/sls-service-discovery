compile: sample

sample:
	GOOS=linux go build -o bin/sample sample.go

bin/rerd: main.go y.go
	go build -o bin/rerd main.go y.go

y.go: rerd.y
	goyacc rerd.y

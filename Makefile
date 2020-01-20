bin/rerd: *.go parser/y.go parser/l.go
	go build -o bin/rerd cmd/rerd/main.go

parser/y.go: rerd.y
	goyacc -o parser/y.go rerd.y

parser/l.go: rerd.l
	golex -o parser/l.go rerd.l

bin/rerd: main.go parser/y.go parser/l.go
	go build -o bin/rerd main.go util.go

parser/y.go: rerd.y
	goyacc -o parser/y.go rerd.y

parser/l.go: rerd.l
	golex -o parser/l.go rerd.l

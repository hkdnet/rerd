bin/rerd: main.go y.go l.go
	go build -o bin/rerd main.go y.go util.go l.go

y.go: rerd.y
	goyacc rerd.y

l.go: rerd.l
	golex -o l.go rerd.l


GOBUILDFLAGS += -ldflags -s

OUTPUT_DIR=./_output

# build: main.go util.go model.go
build:
	go build ${GOBUILDFLAGS} -o ${OUTPUT_DIR}/inception ./

product:
	env GOOS=linux GOARCH=amd64 go build ${GOBUILDFLAGS} -o ${OUTPUT_DIR}/inception.linux ./

test:
	go test $(GOBUILDFLAGS)

clean:
	rm -rf ./${OUTPUT_DIR}

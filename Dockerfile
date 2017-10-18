FROM beekman9527/tensorflowgo

# COPY ./inceptionPredict ${GOPATH}/src/inceptionPredict

RUN mkdir ${GOPATH}/src/inceptionPredict
COPY ./imgs ${GOPATH}/src/inceptionPredict/imgs
COPY ./model-data ${GOPATH}/src/inceptionPredict/model-data
COPY ./main.go ${GOPATH}/src/inceptionPredict/main.go
COPY ./util.go ${GOPATH}/src/inceptionPredict/util.go
COPY ./model.go ${GOPATH}/src/inceptionPredict/model.go
COPY ./Makefile ${GOPATH}/src/inceptionPredict/Makefile
COPY ./run.sh ${GOPATH}/src/inceptionPredict/run.sh

WORKDIR ${GOPATH}/src/inceptionPredict
RUN go get . && make build
RUN sh run.sh

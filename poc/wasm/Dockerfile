FROM tinygo/tinygo

ADD . /plugin
WORKDIR /plugin

RUN go mod download
RUN tinygo build -o ./main.go.wasm -scheduler=none -target=wasi ./main.go
OUT=/tmp/gophish_linux

all:
	rm -rf $(OUT)
	mkdir $(OUT)
	CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" .
	npm install --only=dev
	gulp
	cp -r static $(OUT)
	cp -r templates $(OUT)
	cp -r db $(OUT)
	cp VERSION $(OUT)
	cp config.json $(OUT)
	cp ./gophish $(OUT)

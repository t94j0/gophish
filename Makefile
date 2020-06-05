OUT=/tmp/gophish_linux

all:
	rm -rf $(OUT)
	mkdir $(OUT)
	go build .
	npm install --only=dev
	gulp
	cp -r static $(OUT)
	cp -r templates $(OUT)
	cp -r db $(OUT)
	cp VERSION $(OUT)
	cp config.json $(OUT)
	cp ./gophish $(OUT)

PROJ_NAME = hypr-kbd-layout

MAIN_PATH = *.go
BUILD_PATH = build/package/

INSTALL_PATH = /usr/bin/

install:
	make build-default
	sudo cp $(BUILD_PATH)$(PROJ_NAME) $(INSTALL_PATH)$(PROJ_NAME)

uninstall:
	sudo rm $(INSTALL_PATH)$(PROJ_NAME)

build-default: clean
	go build --ldflags '-extldflags "-static"' -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

build-arm: clean
	GOOS=linux GOARCH=arm GOARM=7 make build-default

build-static: clean
	go build -ldflags "-w -linkmode external -extldflags "-static -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

build-debug: clean
	go build -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

clean:
	rm -rf $(BUILD_PATH)*

tests:
	go test -coverpkg=./... ./... -parallel=2
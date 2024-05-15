DIST_DIR=dist

# ------------------------------------------

# create the dist directory
.PHONY: dist build linux windows darwin build-all
dist:
	mkdir -p $(DIST_DIR)

# clean the dist directory
.PHONY: clean-dist
clean-dist:
	rm -rf $(DIST_DIR)

# ------------------------------------------

# build the Go project
build:
	go build -o $(DIST_DIR)/pkm

# ------------------------------------------

# build the Go project for linux amd64
linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/pkm-linux-amd64

# build the Go project for linux arm64
linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/pkm-linux-arm64

# build the Go project for linux all
linux: linux-amd64 linux-arm64

# ------------------------------------------

# build the Go project for windows
windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/pkm.exe

# build the Go project for windows all
windows: windows-amd64

# ------------------------------------------

# build the Go project for darwin amd64
darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/pkm-darwin-amd64

# build the Go project for darwin arm64
darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/pkm-darwin-arm64

# build the Go project for darwin all
darwin: darwin-amd64 darwin-arm64

# ------------------------------------------

# build the Go project for all OS
build-all: clean-dist dist linux windows darwin

# ------------------------------------------

DOCKER_IMAGE_NAME=linkedinglearning/pkm
DOCKER_IMAGE_TAG=latest

# ------------------------------------------

# build the Go project as a Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

# ------------------------------------------

# run the Go project as a Docker container
docker-run:
	docker run --rm -it $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) get

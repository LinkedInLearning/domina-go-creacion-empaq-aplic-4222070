DOCKER_IMAGE_NAME=mdelapenya/pkmcli
DOCKER_IMAGE_TAG=latest
DIST_DIR=dist
CLI=pkmcli

# ------------------------------------------

# build the Go project as a Docker image
build:
	goreleaser release --snapshot --rm-dist -f .goreleaser.yml

# ------------------------------------------

# run the Go project from the binary (only for MacOS ARM)
run-cli:
	@./$(DIST_DIR)/$(CLI)_darwin_arm64/$(CLI) get

# ------------------------------------------

# run the Go project as a Docker container
run-docker:
	docker run --rm -it $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) get

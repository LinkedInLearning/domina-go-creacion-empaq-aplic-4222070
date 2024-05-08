DOCKER_IMAGE_NAME=linkedinglearning/pkm
DOCKER_IMAGE_TAG=latest

# ------------------------------------------

# build the Go project as a Docker image
build:
	goreleaser release --snapshot --rm-dist -f .goreleaser.ym

# ------------------------------------------

# run the Go project as a Docker container
run:
	docker run --rm -it $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) get

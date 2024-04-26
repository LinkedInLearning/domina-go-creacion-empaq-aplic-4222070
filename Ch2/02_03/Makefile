DOCKER_IMAGE_NAME=linkedinglearning/pkm
DOCKER_IMAGE_TAG=latest

# ------------------------------------------

# build the Go project as a Docker image
build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

# ------------------------------------------

# run the Go project as a Docker container
run:
	docker run --rm -it $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) get

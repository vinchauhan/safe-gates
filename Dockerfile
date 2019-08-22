FROM golang
WORKDIR /workspace
# Assuming the source code is collocated to this Dockerfile, copy the whole
# directory into the container that is building the Docker image.
COPY . .
RUN go build -o /myapp
EXPOSE 8080
# When a container is run from this image, run the binary
CMD /myapp
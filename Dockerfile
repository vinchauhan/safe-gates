FROM golang
WORKDIR /workspace
# Assuming the source code is collocated to this Dockerfile, copy the whole
# directory into the container that is building the Docker image.
COPY . .
RUN go build -o /myapp
EXPOSE 3000
# When a container is run from this image, run the binary
CMD /myapp

# FROM golang as builder
# WORKDIR /workspace
# COPY . .
# RUN go build -o /myapp

# FROM alpine:3.10
# COPY --from=builder /myapp .
# CMD ["./myapp"]

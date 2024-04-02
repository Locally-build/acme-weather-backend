############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

ARG TAG_VERSION
ARG API_TOKEN

# Enable this if you want to have private repositories to be collected
# ARG GITHUB_USER
# ARG GITHUB_TOKEN

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Adding the CA certificates
RUN apk --no-cache add ca-certificates
# Adding the sed tool for versioning
RUN apk add sed

# Enable this if you want to have private repositories to be collected
# RUN git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/acme_weatherservice
COPY . .
WORKDIR /go/src/acme_weatherservice/src

# Updating the main variable.
# RUN sed -i "s/^var ver = \"[[:digit:]]\+\.[[:digit:]]\+\.[[:digit:]]\+\"/var ver = \"${TAG_VERSION}\"/g" main.go

# Using go get.
RUN go get -d -v

# Build the binary.

RUN GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/acme_weatherservice
RUN touch /go/bin/.env
RUN echo "OPENWEATHERMAP_API_KEY=${API_TOKEN}" >> /go/bin/.env
RUN echo "PORT=80" >> /go/bin/.env
COPY ./src/images /go/bin/images

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy SSL Certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy our static executable.
WORKDIR /go/bin
COPY --from=builder /go/bin/acme_weatherservice /go/bin/acme_weatherservice
COPY --from=builder /go/bin/.env /go/bin/.env
COPY --from=builder /go/bin/images /go/bin/images
# RUN cat /go/bin/.env

# Run the project binary.
EXPOSE 80

ENTRYPOINT ["/go/bin/acme_weatherservice"]
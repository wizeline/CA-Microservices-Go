#  BUILD STAGE
ARG DEBIAN_VERSION
ARG GO_VERSION
FROM golang:${GO_VERSION}-${DEBIAN_VERSION} AS builder

ARG WORKDIR
WORKDIR ${WORKDIR}

ADD . .

RUN set -eux \
        && go mod download \
        && go mod verify \
        && go mod tidy -v

RUN go build -o http-api ./cmd/http-rest-api/...

#################
#  RUN STAGE
####
FROM debian:${DEBIAN_VERSION}-slim
LABEL org.name="Wizeline"
LABEL org.image.title="CA-Microservices-GO"
LABEL org.image.description="Debian image to execute a GO Rest API service"

ARG USER="http-api"
ARG PORT=8080
ARG WORKDIR

WORKDIR ${WORKDIR}

RUN  set -eux \
        && useradd \
                --comment "http api user" \
                --create-home \
                ${USER}

COPY --from=builder ${WORKDIR}/http-api /usr/local/bin/http-api

EXPOSE ${PORT}

# USER MODE
USER ${USER}

ENTRYPOINT ["/usr/local/bin/http-api"]
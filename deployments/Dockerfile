#################
#  BUILD STAGE
####
ARG ALPINE_VERSION
ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

ARG APP_NAME
ARG WORKDIR

WORKDIR ${WORKDIR}

ADD . .

RUN set -eux \
        && go mod download \
        && go mod verify \
        && go mod tidy -v

RUN go build -o ${APP_NAME} ./cmd/http-rest-api/...

#################
#  RUN STAGE
####
FROM alpine:${ALPINE_VERSION}
LABEL org.name="Wizeline"
LABEL org.image.title="CA-Microservices-GO"
LABEL org.image.description="Alpine image to execute a GO Rest API service"

ARG APP_NAME
ARG APP_BIN
ARG USER="${APP_NAME}user"
ARG PORT=8080
ARG WORKDIR

WORKDIR ${WORKDIR}

RUN  set -eux \
        && adduser \
                -g "${APP_NAME} api user" \
                --disabled-password \
                ${USER}

COPY --from=builder ${WORKDIR}/${APP_NAME} ${APP_BIN}

EXPOSE ${PORT}

# USER MODE
USER ${USER}

ENTRYPOINT ["/usr/local/bin/camgo"]
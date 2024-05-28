#  BUILD STAGE
ARG GO_VERSION
ARG ALPINE_VERSION
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

ARG APP_NAME
ARG APP_WORKDIR

WORKDIR ${APP_WORKDIR}

ADD . .

RUN set -eux \
   && go mod download \
   && go mod verify \
   && go mod tidy -v \
   && go build -o ${APP_NAME} ./cmd/http-rest-api/...


#  RUN STAGE
FROM alpine:${ALPINE_VERSION}
LABEL org.name="Wizeline"
LABEL org.image.title="CA-Microservices-GO"
LABEL org.image.description="Alpine image to execute a GO Rest API service"

ARG APP_NAME
ARG APP_PORT
ARG APP_WORKDIR
ARG USER="${APP_NAME}_user"
ARG DB_MIGRATIONS_DIR

WORKDIR ${APP_WORKDIR}

RUN  set -eux \
        && adduser \
                -g "${APP_NAME} api user" \
                --disabled-password \
                ${USER}

COPY --from=builder ${APP_WORKDIR}/${APP_NAME} /usr/local/bin/${APP_NAME}
COPY --from=builder ${APP_WORKDIR}/${DB_MIGRATIONS_DIR} ${DB_MIGRATIONS_DIR}

EXPOSE ${APP_PORT}


# USER MODE
USER ${USER}

ENTRYPOINT ["/usr/local/bin/camgo"]
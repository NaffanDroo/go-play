# https://medium.com/tourradar/lean-golang-docker-images-using-multi-stage-builds-1015a6b4d1d1

# Prepare the env for building
ARG GO_VERSION=1.16.3
FROM golang:${GO_VERSION}-alpine AS dev

ENV APP_NAME="main" \
    APP_PATH="/var/app" \
    APP_PORT=3000

ENV APP_BUILD_NAME="${APP_NAME}"

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"

EXPOSE ${APP_PORT}
ENTRYPOINT ["sh"]

# Build
FROM dev as build

RUN (([ ! -d "${APP_PATH}/vendor" ] && go mod download && go mod vendor) || true)
RUN go build -ldflags="-s -w" -mod vendor -o ${APP_BUILD_NAME}
RUN chmod +x ${APP_BUILD_NAME}

# Production build
FROM scratch AS prod

ENV APP_BUILD_PATH="/var/app" \
    APP_BUILD_NAME="main"
WORKDIR ${APP_BUILD_PATH}
COPY --from=build ${APP_BUILD_PATH}/${APP_BUILD_NAME} ${APP_BUILD_PATH}/

EXPOSE ${APP_PORT}
ENTRYPOINT ["/var/app/main"]
CMD ""
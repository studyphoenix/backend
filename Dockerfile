# Build image
FROM alpine AS build

WORKDIR /usr/src/build

COPY . /usr/src/build
ENV PROD=true

RUN apk add go go-swag
RUN swag init --md ./apidocs
RUN go build -ldflags="-w -s" -o backend


# Production image
FROM alpine AS prod
WORKDIR /app

ENV GIN_MODE=release

COPY --from=build /usr/src/build/backend /app/backend
COPY ./templates/ /app/templates
COPY ./config.yml /app/config.yml
RUN ls /app

CMD ["/app/backend"]

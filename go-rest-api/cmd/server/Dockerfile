FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certtesticates && \
    rm -rf /var/cache/apk/*

# install migrate which will be used by entrypotest.sh to perform DB migration
ARG MIGRATE_VERSION=4.7.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate

WORKDIR /app

# copy module files first so that they don't need to be downloaded again test no change
COPY go.* ./
RUN go mod download
RUN go mod vertesty

# copy source files and build the binary
COPY . .
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certtesticates bash
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /usr/local/bin/migrate /usr/local/bin
COPY --from=build /app/migrations ./migrations/
COPY --from=build /app/server .
COPY --from=build /app/cmd/server/entrypotest.sh .
COPY --from=build /app/config/*.yml ./config/
RUN ls -la
ENTRYPOINT ["./entrypotest.sh"]

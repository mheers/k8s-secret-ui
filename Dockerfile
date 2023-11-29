ARG nodego="mheers/node-go:18.17.1-1.21-bookworm"
ARG go="golang:1.21.4-alpine3.18"
ARG base="alpine:3.18.4"

# first stage where we build the frontend
FROM --platform=$BUILDPLATFORM ${nodego} as frontendBuilder
WORKDIR /usr/src/app

# copy the dependency files and prefetch dependencies
COPY frontend/package.json frontend/package-lock.json ./frontend/
# install npm dependencies
RUN cd ./frontend/ && npm install

# copy the source code, ci-scripts, git and VERSION and build the package
COPY frontend/ ./frontend/

# build the frontend
RUN cd ./frontend/ && npm run build

# third stage for building the backend
FROM --platform=$BUILDPLATFORM ${go} AS backendBuilder

RUN apk add --no-cache bash git gcc libc-dev

# Copy the code from the host and compile it
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

ADD . ./

# copy built frontend / storybook
RUN rm -r ./frontend/*
COPY --from=frontendBuilder /usr/src/app/frontend/ ./frontend/

ARG TARGETPLATFORM
ARG BUILDPLATFORM

RUN [ "$(uname)" = Darwin ] && system=darwin || system=linux; \
    ./ci/go-build.sh --os ${system} --arch $(echo $TARGETPLATFORM  | cut -d/ -f2)

# final stage
FROM ${base}
RUN apk add --no-cache curl
WORKDIR /app
COPY --from=backendBuilder /go/src/app/goapp /bin/backend

ENTRYPOINT ["/bin/backend"]
CMD ["server"]
EXPOSE 8000

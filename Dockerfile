FROM node:18.17.1-bullseye AS frontend
WORKDIR /workspace

COPY frontend/yarn.lock frontend/package.json ./

RUN yarn --frozen-lockfile

COPY frontend ./

RUN yarn build

FROM --platform=$BUILDPLATFORM golang:1.20.6-bullseye as backend

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY --from=frontend /workspace/dist/ ./pkg/infrastructure/router/static/app/

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
  go build \
  -a \
  -tags netgo -installsuffix netgo \
  -ldflags="-s -w -extldflags \"-static\"" \
  -o short-url \
  main.go \
  && chmod +x /workspace/short-url

FROM gcr.io/distroless/static:nonroot
COPY --from=backend --chown=nonroot:nonroot /workspace/short-url /usr/local/bin/short-url
ENV TZ=Asia/Tokyo
USER 65532:65532

ENTRYPOINT ["short-url"]

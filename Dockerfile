FROM hub.chinaopen.ai/library/go-builder:1.13-alpine AS builder

ARG COMMIT_ID
ARG VERSION=""
ARG VCS_BRANCH=""
ARG GRPC_STUB_REVISION=""
ARG PROJECT_NAME=boot
ARG DOCKER_PROJECT_DIR=/build
ARG EXTRA_BUILD_ARGS=""
ARG GOCACHE=""

WORKDIR $DOCKER_PROJECT_DIR
COPY . $DOCKER_PROJECT_DIR

ENV GOPRIVATE=git.chinaopen.ai
ENV GOSUMDB=sum.golang.google.cn

RUN git config --global url."git@git.chinaopen.ai:".insteadOf "https://git.chinaopen.ai/" \
    && mkdir -p /output \
    && make build -e GOCACHE=$GOCACHE \
    -e COMMIT_ID=$COMMIT_ID -e OUTPUT_FILE=/output/boot \
    -e VERSION=$VERSION -e VCS_BRANCH=$VCS_BRANCH -e EXTRA_BUILD_ARGS=$EXTRA_BUILD_ARGS

FROM alpine:3.12
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache --update add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai

COPY --from=builder /output/boot /app/boot
COPY config/config.sample.yml /app
COPY gen/apidoc.html /app/

EXPOSE 8080
CMD ["/app/boot", "runserver"]
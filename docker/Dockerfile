FROM ubuntu:22.04 AS builder
LABEL authors="dalibor"

USER root

RUN apt update && apt install wget g++ unzip zip curl git python3 python3-pip -y
RUN apt full-upgrade -y

# Install go
RUN wget https://go.dev/dl/go1.22.2.linux-arm64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xvf go1.22.2.linux-arm64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install bazelisk and bazel
ENV BAZELISK_VERSION 1.19.0
ENV BAZELISK_ARCH arm64

RUN wget -O bazelisk https://github.com/bazelbuild/bazelisk/releases/download/v1.19.0/bazelisk-linux-arm64 && \
    chmod +x ./bazelisk && \
    mv bazelisk /usr/bin/bazel

RUN git clone https://github.com/bazelbuild/bazel-watcher.git /opt/bazel-watcher && \
    cd /opt/bazel-watcher && \
    bazel build //cmd/ibazel

ENV PATH=$PATH:/opt/bazel-watcher/bazel-bin/cmd/ibazel/ibazel_

RUN mkdir /workspace
WORKDIR /workspace

COPY . .

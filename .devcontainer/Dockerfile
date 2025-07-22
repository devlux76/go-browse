# syntax=docker/dockerfile:1
FROM ubuntu:24.04

# Install basic dev tools and dependencies
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y \
    build-essential \
    curl \
    wget \
    git \
    ca-certificates \
    software-properties-common \
    lsb-release \
    gnupg2 \
    unzip \
    sudo && \
    rm -rf /var/lib/apt/lists/*

# Install Go (latest stable)
ENV GO_VERSION=1.24.5
RUN wget -q https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

RUN echo 'ubuntu ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER ubuntu
WORKDIR /workspace

# Default shell
SHELL ["/bin/bash", "-c"]

# Display Go version and Chrome version for verification
RUN go version 

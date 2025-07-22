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
ENV GO_VERSION=1.22.4
RUN wget -q https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Google Chrome (not Chromium)
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | gpg --dearmor > /usr/share/keyrings/google-chrome.gpg && \
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/google-chrome.gpg] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

USER ubuntu
WORKDIR /workspace

# Default shell
SHELL ["/bin/bash", "-c"]

# Display Go version and Chrome version for verification
RUN go version && google-chrome --version

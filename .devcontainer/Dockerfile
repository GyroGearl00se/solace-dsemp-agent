# Devcontainer Dockerfile for solace-dsemp-agent
FROM mcr.microsoft.com/devcontainers/go:1.24

# Install make and any other build tools
RUN apt-get update && apt-get install -y make && rm -rf /var/lib/apt/lists/*

# Set default shell to zsh
SHELL ["/bin/zsh", "-c"]

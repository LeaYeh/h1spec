# Start from the latest Ubuntu version
FROM ubuntu:latest

# Update the system and install utils
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    sudo \
    gcc \
    zsh \
    git \
    wget \
    curl \
    nginx \
    tree


# Detect architecture and download the appropriate Go binary
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz; \
    elif [ "$ARCH" = "aarch64" ]; then \
        wget https://go.dev/dl/go1.20.6.linux-arm64.tar.gz; \
    else \
        echo "Unsupported architecture: $ARCH"; exit 1; \
    fi && \
    tar -C /usr/local -xzf go1.20.6.linux-*.tar.gz && \
    rm go1.20.6.linux-*.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"

# Install oh-my-zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" || true
RUN echo "source $HOME/.oh-my-zsh/oh-my-zsh.sh" >> $HOME/.zshrc

# Set the working directory in the container to /workspace
WORKDIR /workspace

CMD ["/bin/zsh"]

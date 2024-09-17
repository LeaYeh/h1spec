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


# Install the stable version of Go
RUN wget https://go.dev/dl/go1.20.6.linux-arm64.tar.gz && \
    tar -C /usr/local -xzf go1.20.6.linux-arm64.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"

# Install oh-my-zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" || true
RUN echo "source $HOME/.oh-my-zsh/oh-my-zsh.sh" >> $HOME/.zshrc

# Set the working directory in the container to /workspace
WORKDIR /workspace

CMD ["/bin/zsh"]

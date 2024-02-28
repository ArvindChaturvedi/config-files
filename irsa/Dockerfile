FROM ubuntu:latest

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    unzip \
    sudo \
    && apt-get clean

# Install AWS CLI v2
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf awscliv2.zip ./aws

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" \
    && install -o 1001 -g 1001 -m 0755 kubectl /usr/local/bin/kubectl \
    && rm kubectl

# Create a non-root user and group with UID and GID 1001
RUN groupadd -g 1001 appuser \
    && useradd -r -u 1001 -g appuser appuser \
    && mkdir -p /home/appuser \
    && chown -R appuser:appuser /home/appuser

# Grant sudo access to the appuser (optional, adjust as needed)
RUN echo "appuser ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/appuser

# Switch to the non-root user
USER appuser

# Set the working directory
WORKDIR /home/appuser

# Set entrypoint to bash for interactive use
CMD ["/bin/bash"]

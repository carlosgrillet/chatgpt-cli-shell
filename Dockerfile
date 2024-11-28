FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o gpt_cli

FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y \
    openssh-server \
    && apt-get clean
RUN useradd -rm -d /home/chatgpt -s /chatgpt-shell -u 1001 chatgpt
COPY --from=builder /gpt_cli /usr/bin/gpt_cli
RUN chmod +x /usr/bin/chatgpt-shell
RUN mkdir /var/run/sshd \
    && echo 'chatgpt:password' | chpasswd \
    && sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config \
    && sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin no/' /etc/ssh/sshd_config \
    && echo "ForceCommand /usr/bin/gpt_cli" >> /etc/ssh/sshd_config
EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]


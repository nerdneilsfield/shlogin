FROM alpine:latest

COPY shlogin /usr/bin/shlogin
COPY config_example.toml /etc/shlogin/config.toml

ENTRYPOINT ["shlogin", "cron", "/etc/shlogin/config.toml"]
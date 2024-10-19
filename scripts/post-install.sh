#! /bin/sh

mkdir -p /etc/shlogin

mkdir -p /var/log/shlogin

/usr/bin/shlogin gen /etc/shlogin/example.toml
/usr/bin/shlogin gen /etc/shlogin/example.json
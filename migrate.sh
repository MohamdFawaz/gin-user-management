#!/bin/bash

sm=$(which "sql-migrate")
if [ -z $sm ]; then
    go get -u github.com/go-sql-driver/mysql
    go get github.com/rubenv/sql-migrate/...
fi


case "$1" in
    "new")
    $sm new $2
    ;;
    "up")
    $sm up
    ;;
    "redo")
    $sm redo
    ;;
    "status")
    $sm status
    ;;
    "down")
    $sm down
    ;;
    *)
    echo "Usage: $(basename "$0") new | up | status | down"
    exit 1
esac
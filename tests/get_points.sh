#!/usr/bin/env sh

if [ -z "$1" ]; then
    echo "get_points ID"
    exit
fi

curl -X "GET" -H "Content-Type: application/json" "http://localhost:8080/receipts/$1/points"

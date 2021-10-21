#!/bin/sh

if [[ -z "$TG_TOKEN" || -z "$CHAT_ID" ]]; then
  echo "env missing"
  exit 1
fi

./pusher
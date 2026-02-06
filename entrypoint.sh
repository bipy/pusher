#!/bin/sh

if [ -z "$TG_TOKEN" ] || [ -z "$CHAT_ID" ]; then
  echo "Error: TG_TOKEN and CHAT_ID environment variables are required"
  exit 1
fi

exec ./pusher
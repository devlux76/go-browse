#!/bin/bash
# Launch Chrome/Chromium with remote debugging enabled

CHROME_PATH="${CHROME_PATH:-google-chrome}"
PORT="${DEVTOOLS_PORT:-9222}"

"$CHROME_PATH" \
  --remote-debugging-port=$PORT \
  --user-data-dir="$HOME/.config/google-chrome" \
  --no-first-run \
  --no-default-browser-check \
  "$@"

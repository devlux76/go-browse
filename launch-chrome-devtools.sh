#!/bin/bash
# Attempt to launch the default browser with DevTools Protocol enabled.
# Only works if the default browser is Chrome, Edge, or Brave.

BROWSER_CMD="${BROWSER:-google-chrome}"
PORT="${DEVTOOLS_PORT:-9222}"

"$BROWSER_CMD" \
  --remote-debugging-port=$PORT \
  --user-data-dir="$HOME/.config/remote-browser-profile" \
  --no-first-run \
  --no-default-browser-check \
  "$@"
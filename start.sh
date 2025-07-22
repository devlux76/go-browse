
#!/bin/bash
# start.sh: Check for Chrome DevTools and launch browser with remote debugging

CHROME_DEVTOOLS_URL="http://localhost:9222/json/version"
BROWSER_CMD="${BROWSER:-google-chrome}"
PORT="${DEVTOOLS_PORT:-9222}"


# Check if Chrome DevTools is available on the host; if not, launch Chrome with remote debugging
if ! curl -sf "$CHROME_DEVTOOLS_URL" > /dev/null; then
  echo "[INFO] Chrome DevTools is not running on the host. Launching Chrome with remote debugging on port $PORT..."
  "$BROWSER_CMD" \
    --remote-debugging-port=$PORT \
    --no-first-run \
    --no-default-browser-check \
    --disable-popup-blocking \
    about:blank &
  # Wait for Chrome DevTools to become available
  for i in {1..10}; do
    if curl -sf "$CHROME_DEVTOOLS_URL" > /dev/null; then
      echo "[INFO] Chrome DevTools is now available."
      break
    fi
    sleep 1
  done
fi


# Launch browser with remote debugging enabled
"$BROWSER_CMD" \
  --remote-debugging-port=$PORT \
  --user-data-dir="$PROFILE_DIR" \
  --no-first-run \
  --no-default-browser-check \
  "$@"


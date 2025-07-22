
#!/bin/bash
# start.sh: Check for Chrome DevTools and launch browser with remote debugging

CHROME_DEVTOOLS_URL="http://localhost:9222/json/version"
BROWSER_CMD="${BROWSER:-google-chrome}"
PORT="${DEVTOOLS_PORT:-9222}"


# Check if Chrome DevTools is available on the host; if not, launch Chrome with remote debugging
if ! curl -sf "$CHROME_DEVTOOLS_URL" > /dev/null; then
  echo "[INFO] Chrome DevTools is not running on the host. Preparing Chrome profile for remote debugging..."

  # Detect default Chrome user data directory (Linux)
  DEFAULT_PROFILE="$HOME/.config/google-chrome"
  if [ ! -d "$DEFAULT_PROFILE" ]; then
    DEFAULT_PROFILE="$HOME/.config/chromium"
  fi

  # Create a temp directory for the Chrome profile copy
  TEMP_PROFILE="/tmp/chrome-profile-$USER"
  rm -rf "$TEMP_PROFILE"
  mkdir -p "$TEMP_PROFILE"

  # Copy the Default profile (may take a few seconds)
  if [ -d "$DEFAULT_PROFILE/Default" ]; then
    echo "[INFO] Copying Chrome Default profile to temp directory..."
    cp -a "$DEFAULT_PROFILE/Default" "$TEMP_PROFILE/Default"
  else
    echo "[WARN] Default Chrome profile not found. Chrome will start with a fresh profile."
  fi

  echo "[INFO] Launching Chrome with remote debugging on port $PORT using temp profile..."
  "$BROWSER_CMD" \
    --remote-debugging-port=$PORT \
    --user-data-dir="$TEMP_PROFILE" \
    --no-first-run \
    --no-default-browser-check \
    --disable-popup-blocking &

  # Wait for Chrome DevTools to become available
  for i in {1..10}; do
    if curl -sf "$CHROME_DEVTOOLS_URL" > /dev/null; then
      echo "[INFO] Chrome DevTools is now available."
      break
    fi
    sleep 1
  done
fi



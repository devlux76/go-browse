#!/bin/bash
# preflight.sh: Check for Chrome DevTools and launch if available

CHROME_DEVTOOLS_URL="http://host.docker.internal:9222/json/version"

# Check if Chrome DevTools is available
if curl -sf "$CHROME_DEVTOOLS_URL" > /dev/null; then
  echo "Chrome DevTools detected on host. Launching DevTools..."
  ./launch-chrome-devtools.sh
else
  echo "[WARNING] Chrome DevTools is not running on the host (port 9222 not available)."
  echo "Please start Chrome with remote debugging enabled:"
  echo "  google-chrome --remote-debugging-port=9222 &"
fi

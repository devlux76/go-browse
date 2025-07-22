# go-browse

## Getting Started

This project requires a browser running on your host machine with the Chrome DevTools Protocol enabled. The browser is not included or run inside this container or app.

### 1. Start Your Browser with DevTools Protocol

Before using this app, you must start your browser (Chrome, Edge, or Brave) with remote debugging enabled. Use the provided script for your platform:

- **Linux/macOS:**
  ```sh
  ./launch-chrome-devtools.sh
  ```
- **Windows:**
  ```bat
  launch-chrome-devtools.bat
  ```

You may set the `CHROME_PATH` environment variable to specify your browser executable (e.g., `google-chrome`, `microsoft-edge`, `brave-browser`).

### 2. Why?

This app connects to your browser via the DevTools Protocol to enable advanced automation and AI-assisted browsing. The browser must be running on your host for this to work.

### 3. Troubleshooting

- Make sure your browser is started with the script above before using this app.
- The app connects to `host.docker.internal:9222` by default.
- Only Chrome, Edge, and Brave (Chromium-based) are supported.

---

For more details, see the scripts in the project root.

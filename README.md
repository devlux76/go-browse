# go-browse

## This project exposes Chrome Developer Tools (CDP) to an AI by wrapping them in a model context protocol server.

It relies on https://github.com/modelcontextprotocol/go-sdk to provide the MCP paradigms and semantics. It also relies on https://github.com/chromedp/chromedp to provide the connection to a chrome browser.


## Getting Started

This project requires a browser running on your host machine with the Chrome DevTools Protocol enabled. The browser is not included or run inside this container or app.

There is a "start.sh" that will launch the user's browser with the correct settings so that our MCP can connect to it.

There is a binary (along with a docker container if you wish), called go-browse that provides the CDP interface wrapped up in an MCP for AI consumption.

### 2. Why?

This app connects to your browser via the DevTools Protocol to enable advanced automation and AI-assisted browsing. The browser must be running on your host for this to work.

### 3. Troubleshooting

- Make sure your browser is started with the script above before using this app.
- When running in docker the server connects to `host.docker.internal:9222` by default, otherwise it will connect to localhost:9222.

- Only Chrome, Edge, and Brave (Chromium-based) are supported.

---

For more details, see the scripts in the project root.

[![Deploy to Firebase](https://github.com/Mahiro-T/tenho-challenge/actions/workflows/firebase-hosting-merge.yml/badge.svg)](https://github.com/Mahiro-T/tenho-challenge/actions/workflows/firebase-hosting-merge.yml)

# 天和チャレンジ(tenho-challenge)
This project provides a web application that is developed using Go programming language. 
It allows users to randomly arrange Mahjong tiles for a Tenho(Blessing of Heaven) challenge which is a unique role in the Mahjong game. 
Additionally, the application gives the functionality for users to share their generated results on X (Twitter).

## Building and Running

This project uses Go and WebAssembly, compiled with TinyGo.

### Prerequisites

- **Go**: Version 1.22.1. You can download it from [https://golang.org/dl/](https://golang.org/dl/). Make sure Go is installed and the `go` command is in your PATH.
- **TinyGo**: Version 0.34.0. You can download it from [https://tinygo.org/getting-started/install/](https://tinygo.org/getting-started/install/). Make sure TinyGo is installed and the `tinygo` command is in your PATH.

### Building

1.  **Set up Go environment (if not already done)**:
    Ensure your `GOROOT` is set correctly (usually handled by the Go installation) and that `$GOPATH/bin` (or `$HOME/go/bin` by default) is in your `PATH`.
2.  **Download Go module dependencies**:
    Navigate to the project root directory and run:
    ```bash
    go mod download
    ```
3.  **Copy TinyGo's `wasm_exec.js`**:
    The `wasm_exec.js` file required by TinyGo needs to be present in the `public` directory.
    ```bash
    cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" public/
    ```
4.  **Build the Wasm module**:
    Run the build script from the project root:
    ```bash
    ./build.sh
    ```
    This script uses TinyGo to compile `main.go` into `public/main.wasm`.

### Running

1.  **Start an HTTP Server**:
    You need to serve the `public` directory. An easy way to do this is with Python's built-in HTTP server. From the project root directory, run:
    ```bash
    python3 -m http.server --directory public
    ```
    This will typically start a server on `localhost:8000`.
2.  **Open in Browser**:
    Open your web browser and navigate to `http://localhost:8000`.

## Contributors
<a href="https://github.com/mahiro-t/tenho-challenge/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=mahiro-t/tenho-challenge" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
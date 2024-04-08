[![Deploy to Firebase](https://github.com/Mahiro-T/tenho-challenge/actions/workflows/firebase-hosting-merge.yml/badge.svg)](https://github.com/Mahiro-T/tenho-challenge/actions/workflows/firebase-hosting-merge.yml)

# 天和チャレンジ(tenho-challenge)
This project provides a web application that is developed using Go programming language. 
It allows users to randomly arrange Mahjong tiles for a Tenho(Blessing of Heaven) challenge which is a unique role in the Mahjong game. 
Additionally, the application gives the functionality for users to share their generated results on X (Twitter).

## How to Run
1. Move `wasm_exec.js` from the Go root directory.
    ```bash
    $ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" public/.
    ```
1. To compile the Go code to wasm.
    ```bash
    $ GOOS=js GOARCH=wasm go build -o public/main.wasm
    ```
1. Start HTTP Server. this is python example.
    ```bash
    $ python3 -m http.server --directory public
    ```
1. Open your web browser and navigate to `localhost:8000`.

## Contributors
<a href="https://github.com/mahiro-t/tenho-challenge/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=mahiro-t/tenho-challenge" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `[{"name": "wsdl2kotlin","project": "libs","description": "Generate Kotlin DTO's and function from WSDL spec"},{"name": "go-macos-keychain","project": "libs","description": "go library for macos keychain integration"},{"name": "go-oauth-cli-client","project":"libs","description":"go library for oauth flow"},{"name":"neovim","project":"editors","description": "Vim-fork focused on extensibility and usability"},{"name":"javazone-demo","project":"demos","description":"This code!"},{"name":"linux","project":"torvalds","description":"Linux kernel source tree"}]`)
}

func main() {
    http.HandleFunc("/github/projects", handler)
		fmt.Println("Running server...")
    log.Fatal(http.ListenAndServe(":9017", nil))
}

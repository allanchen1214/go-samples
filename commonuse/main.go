package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	// URL encode/decode
	urlStr := "https://www.google.com.hk/?redirect_url=https://www.baidu.com/"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Printf("escapeUrl: %s\n", escapeUrl)

	rawUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Printf("rawUrl: %s\n", rawUrl)

	// Base64 encode/decode
	str := "Hello, World! 你好，世界。"
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("base64 encoded: %s\n", encodedStr)
	decodedBytes, _ := base64.StdEncoding.DecodeString(encodedStr)
	fmt.Printf("base64 decoded: %s\n", string(decodedBytes))
}

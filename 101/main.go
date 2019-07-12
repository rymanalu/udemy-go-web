package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	c := getCode("john@beatles.com")
	fmt.Println(c)
	c = getCode("jon@beatles.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha1.New, []byte("custom_key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

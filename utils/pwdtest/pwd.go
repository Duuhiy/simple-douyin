package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
)

func main() {
	// Using the default options
	salt, encodedPwd := password.Encode("generic password", nil)
	check := password.Verify("generic password", salt, encodedPwd, nil)
	fmt.Println(check) // true

	// Using custom options
	options := &password.Options{10, 10000, 32, sha512.New}
	salt, encodedPwd = password.Encode("generic password", options)
	finalPwd := fmt.Sprintf("%s$%s", salt, encodedPwd)
	fmt.Println(len(finalPwd))
	check = password.Verify("generic password", salt, encodedPwd, options)
	fmt.Println(check) // true
}

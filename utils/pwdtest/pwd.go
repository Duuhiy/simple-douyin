package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
)

func main() {
	// Using the default options
	//salt, encodedPwd := password.Encode("123456", nil)
	//check := password.Verify("1", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	options := &password.Options{10, 10000, 32, sha512.New}
	salt, encodedPwd := password.Encode("123456", options)
	finalPwd := fmt.Sprintf("%s$%s", salt, encodedPwd)
	fmt.Println(len(finalPwd), finalPwd)
	fmt.Println(salt)
	check := password.Verify("123456", salt, encodedPwd, options)
	fmt.Println(check) // true
}

package main

import (
	"fmt"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	res := EncryptPassword("123456")
	fmt.Println(res)
	pass := DecryptPassword(res)
	fmt.Println(pass)
}

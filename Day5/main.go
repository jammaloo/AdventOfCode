package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
)

func main() {
	doorId := "ojvtpuvg"


	i := 0
	password := ""
	for found := 0; found < 8; {
		i++
		hashTarget := []byte(doorId + strconv.Itoa(i))

		md5Hash := fmt.Sprintf("%x", md5.Sum(hashTarget))

		if  md5Hash[0:5]== "00000" {
			fmt.Println("Found char")
			password += md5Hash[5:6]
			found++
			continue
		}
	}
	fmt.Println(password)
}

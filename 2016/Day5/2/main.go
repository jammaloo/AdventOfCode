package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
	"strings"
)

func main() {
	doorId := "ojvtpuvg"

	i := 0
	password := []string{"-", "-", "-", "-", "-", "-", "-", "-"}
	for found := 0; found < 8; {
		i++
		hashTarget := []byte(doorId + strconv.Itoa(i))

		md5Hash := fmt.Sprintf("%x", md5.Sum(hashTarget))

		if  md5Hash[0:5] != "00000" {
			continue
		}

		position, error := strconv.Atoi(md5Hash[5:6])

		//6th character must be be less than 8
		if error != nil || position > 7 {
			continue
		}

		//we should only replace a character in the password once
		if password[position] != "-" {
			continue
		}

		password[position] = md5Hash[6:7]

		fmt.Println(strings.Join(password,""))

		found++
		continue
	}
	fmt.Println("Password: " + strings.Join(password,  ""))
}

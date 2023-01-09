package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	secret := "ckczppom"
	fmt.Println(part1(secret))
	fmt.Println(part2(secret))
}
func getMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func part1(secret string) string {
	for i := 0; i < 100000000; i++ {
		hash := getMd5Hash(secret + fmt.Sprint(i))
		if hash[:5] == "00000" {
			return fmt.Sprint(i)
		}
	}
	return "Not found"
}

func part2(secret string) string {
	for i := 0; i < 100000000; i++ {
		hash := getMd5Hash(secret + fmt.Sprint(i))
		if hash[:6] == "000000" {
			return fmt.Sprint(i)
		}
	}
	return "Not found"
}

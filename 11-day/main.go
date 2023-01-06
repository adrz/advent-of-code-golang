package main

import "fmt"

func main() {
	input := "hepxcrrq"
	new_password := part1(input)
	fmt.Println(new_password)
	fmt.Println(part1(new_password))
}

func hasOneIncreasing(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if int(password[i+1]) == int(password[i])+1 && int(password[i+2]) == int(password[i])+2 {
			return true
		}
	}
	return false
}

func hasNoForbiddenLetters(password string) bool {
	for _, char := range password {
		if char == 'i' || char == 'o' || char == 'l' {
			return false
		}
	}
	return true
}

func hasTwoPairs(password string) bool {
	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func checkPassword(password string) bool {
	return hasOneIncreasing(password) && hasNoForbiddenLetters(password) && hasTwoPairs(password)
}

func increment(password string) string {
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] == 'z' {
			password = password[:i] + "a" + password[i+1:]
		} else {
			password = password[:i] + string(password[i]+1) + password[i+1:]
			return password
		}
	}
	return password
}

func part1(data string) string {
	for {
		data = increment(data)
		if checkPassword(data) {
			return data
		}
	}
}

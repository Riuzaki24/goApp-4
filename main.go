package main

import "fmt"

type account struct {
	login string
	password string
	url string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	outputPassword(login, password, url)

	// account1 := account{
	// 	login,
	// 	password,
	// 	url,
	// }
	account1 := account{
		password: password,
		url: url,
		login: login,
	}
}


func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func outputPassword(login, password, URL string) {
	fmt.Println()
}
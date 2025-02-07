package main

import "fmt"

type account struct {
	login string
	password string
	url string
}

func main() {
	str := []rune("Привет!:)")
	for _, ch := range string(str){
		fmt.Println(ch)
	}



	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{
		password: password,
		url: url,
		login: login,
	}

	outputPassword(&myAccount)
}


func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var result string
	fmt.Scan(&result)
	return result
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
	fmt.Println(*acc)
}
package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login string
	password string
	url string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGJIKLMNOPQRSRUVWXYZ1234567890-*!")

func main() {
	fmt.Println(generatePassword(12))


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


// Данная функция создает массив rune result, длины n, цикл for проходится по result, получая индекс(i) и записывая туда на каджый шаг символ letterRunes[rand.IntN(len(letterRunes))](берет случайный индекс из letterRunes)  
func generatePassword(n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.IntN(len(letterRunes))]
	} 
	return string(result)
}
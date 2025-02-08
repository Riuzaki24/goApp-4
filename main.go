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

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

// Данная функция создает массив rune result, длины n, цикл for проходится по result, получая индекс(i) и записывая туда на каджый шаг символ letterRunes[rand.IntN(len(letterRunes))](берет случайный индекс из letterRunes)  
func (acc *account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.IntN(len(letterRunes))]
	} 
	acc.password = string(result)
}

// Конструктор новой функции 
func newAccount(login, passoword, url string) *account{
	return &account{
		url: url,
		login: login,
		password: password,
	}
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGJIKLMNOPQRSRUVWXYZ1234567890-*!")

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount := newAccount(login, password, url)
	myAccount.generatePassword(12)
	myAccount.outputPassword()
	fmt.Println(myAccount)
}


func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var result string
	fmt.Scan(&result)
	return result
}


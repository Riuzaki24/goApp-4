package main

import (
	"app-4/account"
	"app-4/files"
	"fmt"
)

func main() {
	fmt.Println("Приложение паролей")
	Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}

}

func getMenu() int {
	var variant int
	fmt.Println("Выберете вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount() {

}

func deleteAccount() {

}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	vault.
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}

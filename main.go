package main

import (
	"app-4/account"
	"app-4/cloud"
	"app-4/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Приложение паролей")
	// vault := account.NewVault(files.NewJsonDb("data.json"))
	vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Login")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("удалено")
	} else {
		output.PrintError("Не найдено")
	}

}

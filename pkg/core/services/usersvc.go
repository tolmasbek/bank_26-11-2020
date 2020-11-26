package services

import (
	"bank/models"
	"database/sql"
	"fmt"
)

const AuthorizationOperation=`1.Авторизация
2.Выйти `

const LoginOperation =`Введите логин и пароль: `

func Authorization()(login, password string){
	fmt.Println(AuthorizationOperation)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(LoginOperation)
		fmt.Println("login: ")
		fmt.Scan(&login)
		fmt.Println("password: ")
		fmt.Scan(&password)
		return login, password
	case 2:
		fmt.Println("Good Bye )")
	default:
		fmt.Println("Не корректны ввод попробуйте еще раз")
	}
	return
}

func Login (database *sql.DB, login, password string){
	var User models.Users
	_ = database.QueryRow(`Select *From users Where login=($1) and password=($2)`, login, password).Scan(
		&User.Id,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Remove)
	fmt.Println(User)
}
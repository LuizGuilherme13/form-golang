package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

type usuario struct {
	id        int16
	nome      string
	sobrenome string
	idade     uint8
	email     string
	fone      string
}

func main() {
	var user [3]usuario
	var opc uint8

	fmt.Println("-----------")
	fmt.Println("CADASTRE-SE")
	fmt.Println("-----------")
	for count := 1; count <= 2; count++ {
		user[count].id = int16(count)

		fmt.Println("Digite seu nome: User", count)
		fmt.Scan(&user[count].nome)
		if user[count].nome == "" {
			fmt.Println("Nome não pode ser vazio!")
			count = int(user[count].id)

		}

		fmt.Println("Digite sua sobrenome: User", count)
		fmt.Scan(&user[count].sobrenome)
		if user[count].sobrenome == "" {
			fmt.Println("Sobrenome não pode ser vazio!")
			count = int(user[count].id)
		}

		fmt.Println("Digite seu idade: User", count)
		fmt.Scan(&user[count].idade)
		if user[count].idade < 0 {
			fmt.Println("Idade não pode ser vazio!")
			count = int(user[count].id)
		}

		fmt.Println("Digite seu email: User", count)
		fmt.Scan(&user[count].email)
		if user[count].nome == "" {
			fmt.Println("Email não pode ser vazio!")
			count = int(user[count].id)
		} else {
			checkEmail := checkmail.ValidateFormat(user[count].email)
			if checkEmail != nil {
				fmt.Println("Formato de E-mail inválido!")
				count = int(user[count].id)

			}
		}

		fmt.Println("Digite seu fone: User", count)
		fmt.Scan(&user[count].fone)
		if user[count].nome == "" {
			fmt.Println("Fone não pode ser vazio!")
			count = int(user[count].id)
		}
		fmt.Println("----------------")
	}

	fmt.Println("Qual registro deseja ver?")
	for i := 1; i <= 2; i++ {
		fmt.Println(user[i].id, "-", user[i].nome)
	}
	fmt.Scan(&opc)
	fmt.Println(user[opc].nome, user[opc].sobrenome)
	fmt.Println(user[opc].idade)
	fmt.Println(user[opc].email)
	fmt.Println(user[opc].fone)
}

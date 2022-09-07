package main

import (
	"fmt"
)

type Person struct {
	name, lastName, phone string
	id                    int
}

func main() {
	var p [10]Person
	var nome, sobrenome, fone, continuar string
	var menu int

	for i := 0; i <= 10; i++ {
		fmt.Println("Cadastrando pessoa:", i)
		fmt.Println("Nome:")
		fmt.Scan(&nome)
		fmt.Println("Sobrenome:")
		fmt.Scan(&sobrenome)
		fmt.Println("Fone:")
		fmt.Scan(&fone)

		p[i].id = i
		p[i].name = nome
		p[i].lastName = sobrenome
		p[i].phone = fone
		fmt.Println("==============")

		fmt.Println("Continuar? s/n")
		fmt.Scan(&continuar)

		if continuar == "s" {
			continue
		} else {

			break
		}

	}
	fmt.Println("=====Cadastrados=====")
	for _, value := range p {
		if value.id == 0 && value.name == "" {
			continue
		}
		fmt.Println(value.id, " - ", value.name)
	}

	fmt.Println("O que deseja fazer?:")
	fmt.Println("1 - Editar")
	fmt.Println("2 - Excluir")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		var x int

		fmt.Println("Quem deseja editar?:")
		for _, value := range p {
			if value.id == 0 && value.name == "" {
				continue
			}
			fmt.Println(value.id, " - ", value.name)
		}
		fmt.Scanln(&x)
		fmt.Println("Nome:")
		fmt.Scan(&nome)
		fmt.Println("Sobrenome:")
		fmt.Scan(&sobrenome)
		fmt.Println("Fone:")
		fmt.Scan(&fone)
		p[x].name = nome
		p[x].lastName = sobrenome
		p[x].phone = fone

	case 2:
		var x int

		fmt.Println("Quem deseja excluir?:")
		for _, value := range p {
			if value.id == 0 && value.name == "" {
				continue
			}
			fmt.Println(value.id, " - ", value.name)
		}
		fmt.Scanln(&x)
		p[x].id = 0
		p[x].name = ""
		p[x].lastName = ""
		p[x].phone = ""

	}
	for _, value := range p {
		if value.id == 0 && value.name == "" {
			continue
		}
		fmt.Println(value.id, " - ", value.name)
	}

}

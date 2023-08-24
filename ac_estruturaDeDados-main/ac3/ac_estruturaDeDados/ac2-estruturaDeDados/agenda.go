package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionaContato(contato Contato, contatos *[5]Contato) {
	for i := range contatos {
		if contatos[i] == (Contato{}) {
			contatos[i] = contato
			break
		}
	}
}

func excluiContato(contatos *[5]Contato) {
	for i := len(contatos) - 1; i >= 0; i-- {
		if contatos[i] != (Contato{}) {
			contatos[i] = Contato{}
			break
		}
	}
}

func main() {
	var contatos [5]Contato

	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Adicionar contato")
		fmt.Println("2. Excluir último contato")
		fmt.Println("3. Mostrar contatos")
		fmt.Println("4. Sair")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			fmt.Println("Nome:")
			nome, _ := leitor.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Println("E-mail:")
			email, _ := leitor.ReadString('\n')
			email = strings.TrimSpace(email)

			contato := Contato{
				Nome:  nome,
				Email: email,
			}

			adicionaContato(contato, &contatos)
			fmt.Println("Contato adicionado!")

		case 2:
			excluiContato(&contatos)
			fmt.Println("Último contato excluído!")

		case 3:
			fmt.Println("Lista de contatos:")
			for _, c := range contatos {
				if c != (Contato{}) {
					fmt.Printf("Nome: %s\nE-mail: %s\n", c.Nome, c.Email)
				}
			}

		case 4:
			fmt.Println("Encerrando o programa...")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida!")
		}
		fmt.Println("-------------------------------")
	}
}

var leitor = bufio.NewReader(os.Stdin)


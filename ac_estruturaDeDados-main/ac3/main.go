package main

import (
	"ac3/contato"
	"ac3/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var contatos [5]contato.Contato

	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Adicionar contato")
		fmt.Println("2. Excluir último contato")
		fmt.Println("3. Mostrar contatos")
		fmt.Println("4. Editar contato")
		fmt.Println("5. Sair")

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

			contato := contato.Contato{
				Nome:  nome,
				Email: email,
			}

			utils.AdicionaContato(contato, &contatos)
			fmt.Println("Contato adicionado!")

		case 2:
			utils.ExcluiContato(&contatos)
			fmt.Println("Último contato excluído!")

		case 3:
			exibeContatos(&contatos)

		case 4: //implementar
			exibeContatos(&contatos)
			fmt.Println("Escolha o contato que deseja editar: ")
			var editar int
			fmt.Scanln(&editar)

			if editar >= 0 && editar < len(contatos) {
				fmt.Println("Novo E-mail:")
				novoEmail, _ := leitor.ReadString('\n')
				novoEmail = strings.TrimSpace(novoEmail)

				// Edita o e-mail do contato selecionado
				contatos[editar].Email = novoEmail

				fmt.Println("E-mail do contato atualizado!")
			} else {
				fmt.Println("Índice de contato inválido!")
			}

		case 5:
			fmt.Println("Encerrando o programa...")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida!")
		}
		fmt.Println("-------------------------------")
	}
}

var leitor = bufio.NewReader(os.Stdin)

func exibeContatos(contatos *[5]contato.Contato) {
	fmt.Println("Lista de contatos:")
	for i, c := range contatos {
		if c != (contato.Contato{}) {
			fmt.Printf("[%d] - Nome: %s - E-mail: %s\n", i, c.Nome, c.Email)
		}
	}
}

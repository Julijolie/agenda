package utils

import (
	"ac3/contato"
)

func AdicionaContato(ctt contato.Contato, contatos *[5]contato.Contato) {
	for i := range contatos {
		if contatos[i] == (contato.Contato{}) {
			contatos[i] = ctt
			break
		}
	}
}

func ExcluiContato(contatos *[5]contato.Contato) {
	for i := len(contatos) - 1; i >= 0; i-- {
		if contatos[i] != (contato.Contato{}) {
			contatos[i] = contato.Contato{}
			break
		}
	}
}

func EditaEmail(emailEditado string, ind int, a *[5]contato.Contato) {
	a[ind].AlterarEmail(emailEditado)
}

package main

import "fmt"

type Carro struct{
	Modelo string
	Cor string
	Velocidade int
	EstaLigado bool
}

func (c *Carro) Ligar(){
	c.EstaLigado = true
}
func (c *Carro) Desligar(){
	c.EstaLigado = false
}
func (c *Carro) Acelerar(valor int){
	c.Velocidade += valor
}

func main(){
	carro := Carro{
		Modelo: "Aurora",
		Cor: "cinza",
		Velocidade: 20,
		EstaLigado: true,
	}
	carro.Ligar()
	fmt.Println(carro)

	carro.Acelerar(15)
	fmt.Println(carro)

	carro.Desligar()
	fmt.Println(carro)
}




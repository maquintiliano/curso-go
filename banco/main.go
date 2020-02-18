package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {

	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteMaria := clientes.Titular{"Maria", "123.123.123.12", "Desenvolvedora"}
	contaDaMaria := contas.ContaCorrente{Titular: clienteMaria, NumeroAgencia: 123, NumeroConta: 123456}

	fmt.Println(contaDaMaria.ObterSaldo())
	fmt.Println(contaDaMaria.Depositar(100))
	PagarBoleto(&contaDaMaria, 60)
	fmt.Println(contaDaMaria.ObterSaldo())

	clienteJoao := clientes.Titular{"João", "321.321.321.21", "Estudante"}
	contaDoJoao := contas.ContaPoupanca{Titular: clienteJoao, NumeroAgencia: 123, NumeroConta: 654321}

	fmt.Println(contaDoJoao.ObterSaldo())
	fmt.Println(contaDoJoao.Depositar(100))
	PagarBoleto(&contaDoJoao, 70)
	fmt.Println(contaDoJoao.ObterSaldo())

}

package contas

import (
	"curso-go/banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

//chamando a conta corrente para a func
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente!"
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", c.saldo
	} else {

		return "Valor do depósito menor que 0.", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {

	if c.saldo > valorDaTransferencia && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {

	return c.saldo
}

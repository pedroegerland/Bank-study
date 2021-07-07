package conta

import (
	"fmt"
	"math/rand"
)

// Conta -
type Conta struct {
	Titular       string
	Banco         int
	NumeroAgencia int
	NumeroConta   int
	Documento     string
	TipoConta     string
	Saldo         float64
	PacoteTarifas CestaDeTarifas
	Extrato       []SalvarTransacao
}

// Ext -
type Ext struct {
	Titular       string
	NumeroBanco   int
	NumeroAgencia int
	NumeroConta   int
	Documento     string
	TipoConta     string
}

// CriarConta -
func CriarConta(titular, documento string) Conta {
	fmt.Println("")
	conta := Conta{}
	conta.Titular = titular
	conta.Banco = 851
	conta.NumeroAgencia = 001
	conta.NumeroConta = rand.Intn(999999)
	conta.Documento = documento
	conta.TipoConta = "Conta_Corrente"
	conta.PacoteTarifas.PacoteTed = 0
	conta.PacoteTarifas.PacoteP2p = 0
	conta.PacoteTarifas.TedContaCorrente = 1
	conta.PacoteTarifas.P2pContaCorrente = 2
	conta.PacoteTarifas.ValorContaCorrente = 3.5

	fmt.Println(conta)
	return conta
}

// CriarExterna -
func CriarExterna(titular, documento string) Ext {
	fmt.Println("")
	conta := Ext{}
	conta.Titular = titular
	conta.NumeroBanco = 451
	conta.NumeroAgencia = 256
	conta.NumeroConta = rand.Intn(999999)
	conta.Documento = documento
	conta.TipoConta = "Conta_Corrente"

	fmt.Println(conta)
	return conta
}

// CriarContaPoupanca -
func CriarContaPoupanca(titular, documento string) Conta {
	fmt.Println("")
	conta := Conta{}
	conta.Titular = titular
	conta.Banco = 851
	conta.NumeroAgencia = 001
	conta.NumeroConta = rand.Intn(999999)
	conta.Documento = documento
	conta.TipoConta = "Conta_Poupanca"
	conta.PacoteTarifas.PacoteTed = 0
	conta.PacoteTarifas.PacoteP2p = 0
	conta.PacoteTarifas.TedContaPoupanca = 0
	conta.PacoteTarifas.P2pContaPoupanca = 0
	conta.PacoteTarifas.ValorContaPoupanca = 5.

	fmt.Println(conta)
	return conta
}

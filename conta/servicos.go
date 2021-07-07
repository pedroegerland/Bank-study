package conta

import (
	"fmt"
	"time"
)

// Extrato -
type Extrato struct {
	extrato []SalvarTransacao
}

// SalvarTransacao -
type SalvarTransacao struct {
	data          string
	valor         string
	tipoTransacao string
	conta         int
	descricao     string
	idTransacao   string
}

// CestaDeTarifas -
type CestaDeTarifas struct {
	PacoteTed          int
	PacoteP2p          int
	TedContaCorrente   int
	P2pContaCorrente   int
	ValorContaCorrente float64
	P2pContaPoupanca   int
	TedContaPoupanca   int
	ValorContaPoupanca float64
}

// Salvar -
func Salvar(valor float64, tipoTransacao, descricao string, idTransacao string, c *Conta) {
	d := time.Now().Format("02 Jan 06 15:04:05")
	stringValor := fmt.Sprintf("%g", valor)
	var newValor string
	if tipoTransacao == "ted-out" || tipoTransacao == "p2p-out" || tipoTransacao == "tarifa p2p-out" || tipoTransacao == "tarifa ted-out" || tipoTransacao == "cash-out" {
		newValor = "-" + stringValor
	} else {
		newValor = stringValor
	}
	s := SalvarTransacao{}
	s.data = d
	s.descricao = descricao
	s.idTransacao = idTransacao
	s.valor = newValor
	s.tipoTransacao = tipoTransacao
	s.conta = c.NumeroConta

	c.Extrato = append(c.Extrato, s)
}

// MostrarExtrato -
func (e *Extrato) MostrarExtrato(c *Conta) {
	fmt.Print("\n============================================================")
	fmt.Print("\nExtrato ", c.Titular)
	fmt.Print("\n============================================================\n")

	for i, line := range c.Extrato {
		if line.conta == c.NumeroConta {
			e.extrato = append(e.extrato, c.Extrato[i])
			fmt.Print("\nData: "+c.Extrato[i].data,
				"\nDescrição: "+c.Extrato[i].descricao,
				"\nid: "+c.Extrato[i].idTransacao+"\nValor: ",
				c.Extrato[i].valor,
				"\nTipo da Transação: "+c.Extrato[i].tipoTransacao+"\nNúmero da conta: ",
				c.Extrato[i].conta,
				"\n")
		}
	}
}

// ObterSaldo -
func ObterSaldo(conta *Conta) float64 {
	return conta.Saldo
}

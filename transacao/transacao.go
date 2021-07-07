package transacao

import (
	"alura/banco/conta"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

//DepositoAtm -
func DepositoAtm(valorDoDeposito float64, contaRecebedora *conta.Conta) {
	uuid := uuid.Must(uuid.NewV4()).String()
	if valorDoDeposito > 0 {
		contaRecebedora.Saldo += valorDoDeposito
		conta.Salvar(valorDoDeposito, "cash-in", "Deposito via ATM", uuid, contaRecebedora)
		fmt.Print("\nDeposito realizado com sucesso\n")
	} else {
		fmt.Print("\nNão foi possível realizar o depósito")
	}
}

//Depositar -
func Depositar(valorDoDeposito float64, tipoTransacao, descricao, uuid string, contaRecebedora *conta.Conta) (bool, string) {
	if valorDoDeposito > 0 {
		contaRecebedora.Saldo += valorDoDeposito
		conta.Salvar(valorDoDeposito, tipoTransacao, descricao, uuid, contaRecebedora)
		return true, ""
	}
	return false, "Não foi possível realizar o depósito"
}

// Saque -
func Saque(valorSaque float64, contaPagadora *conta.Conta) {
	if contaPagadora.Saldo > valorSaque {
		contaPagadora.Saldo -= valorSaque
		uuid := uuid.Must(uuid.NewV4()).String()
		conta.Salvar(valorSaque, "cash-out", "Saque", uuid, contaPagadora)
		fmt.Print("\nSaque realizado com sucesso\n")
	} else {
		fmt.Print("\nNão foi possível realizar o saque")
	}
}

// Retirar -
func Retirar(valor float64, tipoTransacao, descricao, uuid string, contaPagadora *conta.Conta) (bool, string) {
	acum := valor + contaPagadora.PacoteTarifas.ValorContaCorrente
	acum2 := valor + contaPagadora.PacoteTarifas.ValorContaPoupanca
	if tipoTransacao == "ted-out" || tipoTransacao == "p2p-out" {
		if valor <= contaPagadora.Saldo {
			if contaPagadora.TipoConta == "Conta_Corrente" {
				if contaPagadora.PacoteTarifas.PacoteP2p > contaPagadora.PacoteTarifas.P2pContaCorrente {
					if contaPagadora.Saldo >= acum {
						contaPagadora.Saldo -= valor
						conta.Salvar(valor, tipoTransacao, descricao, uuid, contaPagadora)
						TaxaCestaDeTarifas(contaPagadora, tipoTransacao, descricao)
						return true, ""
					}
					fmt.Println("Saldo insuficente")
				} else {
					contaPagadora.Saldo -= valor
					conta.Salvar(valor, tipoTransacao, descricao, uuid, contaPagadora)
					TaxaCestaDeTarifas(contaPagadora, tipoTransacao, descricao)
					return true, ""
				}

				if contaPagadora.PacoteTarifas.PacoteTed > contaPagadora.PacoteTarifas.TedContaCorrente {
					if contaPagadora.Saldo >= acum {
						contaPagadora.Saldo -= valor
						conta.Salvar(valor, tipoTransacao, descricao, uuid, contaPagadora)
						TaxaCestaDeTarifas(contaPagadora, tipoTransacao, descricao)
						return true, ""
					}
					fmt.Println("Saldo insuficente")
				} else {
					contaPagadora.Saldo -= valor
					conta.Salvar(valor, tipoTransacao, descricao, uuid, contaPagadora)
					TaxaCestaDeTarifas(contaPagadora, tipoTransacao, descricao)
					return true, ""
				}
			} else {
				if contaPagadora.Saldo >= acum2 {
					contaPagadora.Saldo -= valor
					conta.Salvar(valor, tipoTransacao, descricao, uuid, contaPagadora)
					TaxaCestaDeTarifas(contaPagadora, tipoTransacao, descricao)
					return true, ""
				}
				fmt.Println("Saldo insuficente")
			}
		}
	} else {
		if valor <= contaPagadora.Saldo {
			contaPagadora.Saldo -= valor
			conta.Salvar(valor, tipoTransacao, descricao, uuid, contaPagadora)
		}
	}
	return false, "Não foi possível realizar a retirada do valor, saldo insuficiente"
}

// P2p -
func P2p(valor float64, descricao string, contaPagadora, contaRecebedora *conta.Conta) {
	fmt.Println("")
	uuid := uuid.Must(uuid.NewV4()).String()
	statusRetirar, errRetirar := Retirar(valor, "p2p-out", descricao, uuid, contaPagadora)
	if statusRetirar {
		statusDepositar, errDepositar := Depositar(valor, "p2p-in", descricao, uuid, contaRecebedora)
		if statusDepositar {
			fmt.Println("Transferência entre contas realizada com sucesso")
		} else {
			fmt.Println(errDepositar)
		}
	} else {
		fmt.Println(errRetirar)
	}

}

// Transferir -
func Transferir(valor float64, descricao string, contaPagadora *conta.Conta, Externa *conta.Ext) {
	fmt.Println("")
	uuid := uuid.Must(uuid.NewV4()).String()
	statusRetirar, errRetirar := Retirar(valor, "ted-out", descricao, uuid, contaPagadora)
	if statusRetirar {
		fmt.Println("TED realizada com sucesso")
	} else {
		fmt.Println(errRetirar)
	}
}

// TaxaCestaDeTarifas -
func TaxaCestaDeTarifas(contaPagadora *conta.Conta, tipoTransacao, descricao string) {
	uuid := uuid.Must(uuid.NewV4()).String()
	tipoTransacao2 := "tarifa " + tipoTransacao
	if contaPagadora.TipoConta == "Conta_Poupanca" {

		Retirar(contaPagadora.PacoteTarifas.ValorContaPoupanca, tipoTransacao2, descricao, uuid, contaPagadora)

		fmt.Print("Tarifa " + tipoTransacao + " Aplicada com sucesso\n\n")
	} else {
		if tipoTransacao == "p2p-out" {
			if contaPagadora.PacoteTarifas.PacoteP2p > contaPagadora.PacoteTarifas.P2pContaCorrente {

				Retirar(contaPagadora.PacoteTarifas.ValorContaCorrente, tipoTransacao2, descricao, uuid, contaPagadora)

				fmt.Print("Tarifa P2p Aplicada com sucesso\n\n")

				contaPagadora.PacoteTarifas.PacoteP2p++
			} else {
				contaPagadora.PacoteTarifas.PacoteP2p++
			}
		}
		if tipoTransacao == "ted-out" {
			if contaPagadora.PacoteTarifas.PacoteTed > contaPagadora.PacoteTarifas.TedContaCorrente {

				Retirar(contaPagadora.PacoteTarifas.ValorContaCorrente, tipoTransacao2, descricao, uuid, contaPagadora)

				fmt.Print("Tarifa Ted Aplicada com sucesso\n\n")

				contaPagadora.PacoteTarifas.PacoteTed++
			} else {
				contaPagadora.PacoteTarifas.PacoteTed++
			}
		}
	}
}

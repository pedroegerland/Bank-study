package main

import (
	"alura/banco/conta"
	"alura/banco/transacao"
)

func main() {

	e := conta.Extrato{}

	contaDoPedro := conta.CriarConta("Pedro Egerland", "749.456.990-11")
	contaDaKaren := conta.CriarConta("Karen Clemente", "086.602.440-95")
	contaDaCindy := conta.CriarConta("Cindy Shitzu", "049.729.200-92")
	contaDaBianca := conta.CriarContaPoupanca("Bianca Clemente Egerland", "876.460.090-40")
	contaDoExternaPedro := conta.CriarExterna("Pedro Egerland", "749.456.990-11")

	transacao.DepositoAtm(2500., &contaDoPedro)
	transacao.Saque(500., &contaDoPedro)

	transacao.P2p(250., "Para meu amor", &contaDoPedro, &contaDaKaren)
	transacao.P2p(250., "Para meu Dog", &contaDoPedro, &contaDaCindy)
	transacao.P2p(124., "Para minha princesa", &contaDoPedro, &contaDaBianca)
	transacao.P2p(131., "Para minha princesa", &contaDoPedro, &contaDaBianca)

	transacao.Transferir(200., "Para meu banco", &contaDoPedro, &contaDoExternaPedro)
	transacao.Transferir(300., "Para meu banco", &contaDoPedro, &contaDoExternaPedro)
	transacao.Transferir(300., "Para meu banco", &contaDoPedro, &contaDoExternaPedro)

	transacao.P2p(250., "Para meu pai", &contaDaBianca, &contaDoPedro)

	transacao.P2p(9999999., "Para minha princesa", &contaDoPedro, &contaDaBianca)
	transacao.Transferir(9999999., "Para meu banco", &contaDoPedro, &contaDoExternaPedro)
	transacao.P2p(9999999., "Para meu pai", &contaDaBianca, &contaDoPedro)

	e.MostrarExtrato(&contaDoPedro)
	e.MostrarExtrato(&contaDaKaren)
	e.MostrarExtrato(&contaDaBianca)
}

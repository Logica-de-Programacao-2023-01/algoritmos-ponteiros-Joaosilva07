type Conta struct {
    saldo float64
}

func adicionarValorAoSaldo(conta *Conta, valor float64) {
    conta.saldo += valor
}

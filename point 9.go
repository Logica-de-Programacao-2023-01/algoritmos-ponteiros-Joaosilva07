type Livro struct {
    Título string
    Autor  string
    Preço  float64
}

func aplicarDesconto(livro *Livro) {
    desconto := livro.Preço * 0.1
    livro.Preço -= desconto
}

type Livro struct {
    Título string
    Autor  string
}

func alterarTituloSeAutorDesconhecido(livro *Livro) {
    if livro.Autor == "Anônimo" {
        livro.Título = "Desconhecido"
    }
}

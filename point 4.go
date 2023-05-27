func somaUltimosDigitos(ptr *int) {
    num := *ptr % 100
    digit1 := num / 10
    digit2 := num % 10
    *ptr = digit1 + digit2
}

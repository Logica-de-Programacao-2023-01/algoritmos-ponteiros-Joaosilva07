func verificarParImpar(ptr *int) {
    if *ptr%2 == 0 {
        *ptr = 0 // par
    } else {
        *ptr = 1 // ímpar
    }
}

func somaPrimeirosNaturais(ptr *int, n int) {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    *ptr = sum
}

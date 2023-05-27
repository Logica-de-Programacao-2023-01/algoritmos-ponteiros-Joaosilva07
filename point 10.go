func preencherPrimos(ptr *[]int, n int) {
    isPrimo := func(num int) bool {
        if num < 2 {
            return false
        }
        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                return false
            }
        }
        return true
   

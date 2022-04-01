package fibo

func Fibo() func() int {
	nb1 := 0
	nb2 := 1

	return func() int {
		ret_val := nb1
		nb1, nb2 = nb2, nb2+nb1
		return ret_val
	}
}

package module01

func GCD(a, b int) int {
	gcd := 0
	if a >= b {
		gcd = b
	} else {
		gcd = a
	}
	for gcd > 1 {
		if a%gcd == 0 && b%gcd == 0 {
			return gcd
		}
		gcd -= 1

	}
	return gcd
}

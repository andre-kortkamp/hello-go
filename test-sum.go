package main

func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x+y, true
	}
	return x+y, false
}

func main() {
	resultado, status := soma(11, 20)
	resultado, status = soma(10, 20)
	println(resultado, status)
}

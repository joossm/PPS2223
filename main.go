package main

func main() {
	println(f(1, 1))

}
func f(x int, y int) int {
	if x == 0 {
		return y
	} else {
		if y == 0 {
			return x
		} else {
			xp := f(x-1, y)
			yp := f(x, y-1)
			return 1 + f(xp, yp)
		}
	}
}

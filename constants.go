package main

import ("fmt")
import ("math")

const s string = "constants"
func main(){

	fmt.Println(s)
	const n = 50000000
	const d = 3e20/n
	fmt.Println(d)
	fmt.Println(math.Sin(n))

}
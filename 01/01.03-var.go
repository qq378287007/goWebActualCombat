package main

func main() {

	//var (
	//	age     int = 28
	//	name    string = "shirdon"
	//	balance float32 = 999999.99
	//)

	//var (
	//	age int
	//	name string
	//	balance float32
	//)

	var language string = "Go"
	//var language = "Go"
	//language := "Go"

	//var age, name, balance = 30, "liao", 99.99
	age, name, balance := 30, "shirdon", 999999.99

	d, c := "D", "C"
	c, d = d, c

	println(c)
	println(d)

	println(age)
	println(name)
	println(balance)
	println(language)
}

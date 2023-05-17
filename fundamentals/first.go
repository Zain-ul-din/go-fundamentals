package main

const url = "https://www.example.com"

func main() {
	print("hello world\n")
	// variables
	var message string = "Hello from GO!"
	var price = 14.4 // Implicit type

	foo := "it's foo" // initializing shortcut

	print(message, price, foo, "\n")

	println("url: ", url)

	printData()
	fmtPackageDemo()
}

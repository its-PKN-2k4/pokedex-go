package main

func main() {
	//print("Hello, World!\n")
	processed := cleanInput("  hello    world    ")
	for _, word := range processed {
		print(word)
		print("\n")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	print("Enter text: ")
	text, _ := reader.ReadString('\n')
	println("Your text is", text)

	fmt.Print("Enter text: ")
	text2 := ""
	fmt.Scanf("%s", &text2)
	fmt.Println("Your second text is: ", text2)

}

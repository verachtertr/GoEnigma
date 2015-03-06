package main

import "os"
import "fmt"
import "GoEnigma/enigma"

func main() {

	if len(os.Args) == 1 {
		println("Usage:", os.Args[0], "input.json")
		return
	}

	enigma_config, err := enigma.ReadConfig(os.Args[1])

	if err != nil {
		println("Error:", err)

	}

	fmt.Println(enigma_config)

}

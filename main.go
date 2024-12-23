package main

import (
	"findMe/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Find Me!")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

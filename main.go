package main

import (
	"fmt"
	"os"
)

func main() {	
	data, err := os.ReadFile("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	
	y := 30
	fmt.Println("Valor de y:", y)
	
	fmt.Println("Conteúdo:", string(data))
	fmt.Println("Executando código normalmente")

	z := 50
	fmt.Println("Valor de z:", z)

	if err := os.Chdir("/caminho/invalido"); err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
	}
}

// main.go
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
	
	// Atribuição ineficiente
	var y int
	y = 20
	y = 30

	fmt.Println("Conteúdo:", string(data))
	
	// Condição que sempre avalia para true
	if 1 == 1 {
		fmt.Println("Esta condição sempre é verdadeira")
	}

	z := 50
	fmt.Println(z, y)

	// Retorno de função ignorado
	os.Chdir("/caminho/invalido")
}

// Função não utilizada
func unusedFunction() string {
	return "Esta função nunca é chamada"
}

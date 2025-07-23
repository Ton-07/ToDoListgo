// Inicializa o servidor, conecta ao banco, configura as rotas e dá start na API.

package main

import (
	"fmt"
)

func main() {

	todos := todos{}
	todos.add("Estudar Go")
	todos.add("Levar o Thor para passear")
	todos.add("Terminar o to-do-list em go")

	fmt.Printf("%+v\n\n", todos[2])
	todos.print()

	todos.remover(0)

	todos.print()

	fmt.Printf("%+v\n\n", todos[1])

	todos.alternarStatus(0)

	todos.print()

}

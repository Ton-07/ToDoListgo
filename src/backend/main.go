package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Conectar no banco
	var err error
	dsn := "postgres://ton:1234@localhost:5433/todolist"
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(fmt.Errorf("erro ao conectar no banco: %w", err))
	}
	defer DB.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(`
--- MENU ---
1. Adicionar tarefa
2. Listar tarefas
3. Remover tarefa
4. Marcar como feita
5. Desmarcar tarefa
6. Sair
Escolha uma opção:`)

		scanner.Scan()
		opcao := scanner.Text()

		switch opcao {
		case "1":
			fmt.Print("Digite o título da tarefa: ")
			scanner.Scan()
			titulo := scanner.Text()
			err := InserirTodo(titulo)
			if err != nil {
				fmt.Println("Erro ao adicionar tarefa:", err)
			} else {
				fmt.Println("✅ Tarefa adicionada!")
			}

		case "2":
			todos, err := BuscarTodos()
			if err != nil {
				fmt.Println("Erro ao listar tarefas:", err)
				continue
			}
			PrintTodos(todos)

		case "3":
			fmt.Print("Digite o ID da tarefa para remover: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}
			err = RemoverTodo(id)
			if err != nil {
				fmt.Println("Erro ao remover tarefa:", err)
			} else {
				fmt.Println("🗑️ Tarefa removida.")
			}

		case "4":
			fmt.Print("Digite o ID da tarefa para marcar como feita: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}
			todos, err := BuscarTodos()

			if id < 1 || id > len(todos) {
				fmt.Printf("posição inválida: %d", id)
				continue

			}

			if todos[id-1].Feito {
				fmt.Println("A tarefa já está marcada como feita.")
				continue
			}

			err = AlternarStatus(id)
			if err != nil {
				fmt.Println("Erro ao marcar como feita:", err)
			} else {
				fmt.Println("✅ Tarefa marcada como feita.")
			}

		case "5":
			fmt.Print("Digite o ID da tarefa para desmarcar: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}
			todos, err := BuscarTodos()

			if id < 1 || id > len(todos) {
				fmt.Printf("posição inválida: %d", id)
				continue

			}

			if !todos[id-1].Feito {
				fmt.Println("A tarefa já está marcada como pendente.")
				continue
			}

			err = AlternarStatus(id)
			if err != nil {
				fmt.Println("Erro ao desmarcar tarefa:", err)
			} else {
				fmt.Println("❌ Tarefa desmarcada.")
			}

		case "6":
			fmt.Println("👋 Saindo...")
			return

		default:
			fmt.Println("Opção inválida! 😅")
		}
	}
}

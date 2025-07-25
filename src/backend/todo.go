package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jedib0t/go-pretty/table"
)

var DB *pgxpool.Pool

type todo struct {
	ID         int        `db:"id"`
	Titulo     string     `db:"titulo"`
	Feito      bool       `db:"feito"`
	Criado     time.Time  `db:"criado_em"`
	Finalizado *time.Time `db:"finalizado_em"`
}

func BuscarTodos() ([]todo, error) {
	rows, err := DB.Query(context.Background(), `
		SELECT id, titulo, feito, criado_em, finalizado_em 
		FROM todos 
		ORDER BY criado_em ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []todo
	for rows.Next() {
		var t todo
		err := rows.Scan(&t.ID, &t.Titulo, &t.Feito, &t.Criado, &t.Finalizado)
		if err != nil {
			return nil, err
		}
		lista = append(lista, t)
	}
	return lista, nil
}

func InserirTodo(titulo string) error {
	_, err := DB.Exec(context.Background(), `INSERT INTO todos (titulo) VALUES ($1)`, titulo)
	return err
}

func RemoverTodo(posicaoVisual int) error {
	todos, err := BuscarTodos()
	if err != nil {
		return err
	}

	if posicaoVisual < 1 || posicaoVisual > len(todos) {
		return fmt.Errorf("posição inválida: %d", posicaoVisual)
	}

	id := todos[posicaoVisual-1].ID

	cmdTag, err := DB.Exec(context.Background(), `DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("tarefa com id %d não encontrada", id)
	}
	return nil
}

func AlternarStatus(posicaoVisual int) error {
	todos, err := BuscarTodos()
	if err != nil {
		return err
	}

	if posicaoVisual < 1 || posicaoVisual > len(todos) {
		return fmt.Errorf("posição inválida: %d", posicaoVisual)
	}

	id := todos[posicaoVisual-1].ID

	var feito bool
	err = DB.QueryRow(context.Background(), `SELECT feito FROM todos WHERE id = $1`, id).Scan(&feito)
	if err != nil {
		return err
	}

	if !feito {
		_, err = DB.Exec(context.Background(),
			`UPDATE todos SET feito = true, finalizado_em = NOW() WHERE id = $1`, id)
	} else {
		_, err = DB.Exec(context.Background(),
			`UPDATE todos SET feito = false, finalizado_em = NULL WHERE id = $1`, id)
	}
	return err
}

func PrintTodos(todos []todo) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Criado.Before(todos[j].Criado)
	})

	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"#", "Título", "Finalizada", "Criada em", "Concluída em"})
	for index, t := range todos {
		finalizada := "❌"
		completedAt := ""

		if t.Feito {
			finalizada = "✅"
			if t.Finalizado != nil {
				completedAt = t.Finalizado.Local().Format(time.RFC1123)

			}
		}

		tbl.AppendRow(table.Row{
			strconv.Itoa(index + 1),
			t.Titulo,
			finalizada,
			t.Criado.Local().Format(time.RFC1123),
			//t.Criado.Format(time.RFC1123),
			completedAt,
		})
	}

	tbl.Render()
}

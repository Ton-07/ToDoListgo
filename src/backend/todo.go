package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

// Add ==> ok
// Remover ==> ok
// validar se o index está dentro dos limites ==> ok
// alternar o status de conclusão de uma tarefa.
// método de edição
//  método chamado print

type todo struct {
	titulo     string
	feito      bool
	criado     time.Time
	finalizado *time.Time
}

type todos []todo // array do tipo todo

func (ts *todos) add(titulo string) {
	now := time.Now()
	task := todo{
		titulo:     titulo,
		feito:      false,
		criado:     now,
		finalizado: nil,
	}
	*ts = append(*ts, task)
}

// Precisamos verificar se o nosso index está dentro do limite do array

func (list *todos) indexvalido(index int) error {

	if index < 0 || index >= len(*list) {
		err := errors.New("index inválido")
		fmt.Println(err.Error())
		return err
	}

	return nil

}

func (ts *todos) remover(index int) error {
	t := *ts

	if err := t.indexvalido(index); err != nil {
		return err
	}

	*ts = append(t[:index], t[index+1:]...)
	return nil
}

func (ts *todos) alternarStatus(index int) error {
	t := (*ts)

	if err := t.indexvalido(index); err != nil {
		return err
	}

	// // Inverte o status 'feito' da tarefa

	estaFeita := t[index].feito
	if !estaFeita {
		feitaem := time.Now()
		t[index].finalizado = &feitaem
	}

	t[index].feito = !estaFeita

	return nil

}

func (ts *todos) edit(index int, titulo string) error {
	t := (*ts)

	if err := t.indexvalido(index); err != nil {
		return err
	}

	// Editar o index
	t[index].titulo = titulo

	return nil

}

func (todos *todos) print() {
	// table := table.New(os.Stdout)
	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"#", "Title", "Completed", "Created At", "Completed At"})
	for index, t := range *todos {
		finalizada := "❌"
		completedAt := ""

		if t.feito {
			finalizada = "✅"
			if t.finalizado != nil {
				completedAt = t.finalizado.Format(time.RFC1123)
			}
		}

		tbl.AppendRow(table.Row{
			strconv.Itoa(index),
			t.titulo,
			finalizada,
			t.criado.Format(time.RFC1123),
			completedAt,
		})
	}

	tbl.Render()
}

# 🧠 ToDo CLI — Gerenciador de Tarefas com Go + PostgreSQL + Docker

Um app de terminal simples e funcional para gerenciar tarefas direto pelo CLI. Feito em Go, com persistência em PostgreSQL e um toque de terminal bonito via `go-pretty`.

---

## 📦 Tecnologias

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [pgx](https://github.com/jackc/pgx) — driver PostgreSQL para Go
- [go-pretty/table](https://github.com/jedib0t/go-pretty) — tabelas bonitonas no terminal

---

## 📁 Estrutura do Projeto

📦 TODOLISTGO
┣ 📂db
┃ ┗ 📜 connection.go # Conexão com o banco
┣ 📂src\backend
┃ ┣ 📂migrations
┃ ┃ ┣ 📜 <timestamp>_create_todos_table.up.sql
┃ ┃ ┗ 📜 <timestamp>_create_todos_table.down.sql
┣ 📜 main.go # Roteamento e servidor
┣ 📜 todo.go # Modelo e lógica das tarefas
┣ 📜 .env # Variáveis de ambiente
┣ 📜 .env.example # Exemplo para uso
┣ 📜 docker-compose.yml # Banco via Docker
┣ 📜 go.mod / go.sum # Dependências do projeto
┗ 📜 README.md # Este arquivo


---

## 🚀 Como rodar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/seuuser/todolistgo.git
cd todolistgo

2. Configure o .env
Crie um arquivo .env baseado no .env.example:

cp .env.example .env
Edite as variáveis conforme necessário:

POSTGRES_USER=coloque_seu_usuario
POSTGRES_PASSWORD=coloque_sua_senha
POSTGRES_DB=nome_do_banco


---

3. Suba o banco com Docker

Certifique-se de ter o Docker e o Docker Compose instalados. Depois, rode:

docker-compose up -d


4. Rode as migrations (manual por enquanto)

O Docker já vem com o cliente psql instalado por padrão, então você pode interagir com o banco direto pelo terminal com:

Você pode usar o arquivo .sql que já está no projeto para criar a tabela todos.

Opção 1 — Usando o psql da sua máquina

psql -h localhost -U ton -p 5433 -d todolist -f src/backend/migrations/20250723145125_create_todos_table.up.sql

Opção 2 — Criando a tabela manualmente
Se quiser colar direto no terminal do PostgreSQL:

CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  titulo TEXT NOT NULL,
  feito BOOLEAN NOT NULL DEFAULT false,
  criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

Se funcionar, você verá algo assim:

CREATE TABLE

5. Execute a API

go run main.go todo.go
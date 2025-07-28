# 🧠 TaskNotify

TaskNotify é uma aplicação simples de gerenciamento de tarefas com backend em Go utilizando o framework Gin e GORM como ORM. Ela permite criar usuários e associar tarefas a esses usuários. Ideal para estudos ou como base para sistemas mais complexos.

---

## 🚀 Funcionalidades

- ✅ Cadastro de usuários
- ✅ Criação, listagem, atualização e finalização de tarefas
- ✅ Associação de tarefas a usuários
- ✅ Persistência com PostgreSQL
- ✅ Autenticação (em desenvolvimento)
- ✅ Suporte a variáveis de ambiente com `.env`

---

## 🛠️ Tecnologias

- [Golang](https://golang.org)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [dotenv](https://github.com/joho/godotenv)

---

## 🔧 Como rodar localmente

### Pré-requisitos

- Go 1.20+
- PostgreSQL rodando localmente
- Variáveis de ambiente configuradas

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/tasknotify.git
cd tasknotify
```

### 2. Configure o .env


```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=tasknotify_db
PORT=8080
```

### 3. Rode a aplicação

```bash
go run main.go
```
#### A API será executada em: ```http://localhost:8080```

---

## 📬 Endpoints principais

| Método | Rota                | Descrição                       |
| ------ | ------------------- | ------------------------------- |
| POST   | `/users/`           | Cria um novo usuário            |
| GET    | `/tasks`            | Lista todas as tarefas          |
| POST   | `/tasks`            | Cria uma nova tarefa            |
| PUT    | `/tasks/:id`        | Atualiza uma tarefa             |
| PUT    | `/tasks/:id/finish` | Marca uma tarefa como concluída |











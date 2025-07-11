# API de Gerenciamento de hábitos - Habit Manager

<div align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge" height="40" alt="go logo"  />
  <img width="20" />
  <img src="https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white&style=for-the-badge" height="40" alt="docker logo"  />
  <img width="20" />
  <img src="https://img.shields.io/badge/MySQL-4479A1?logo=mysql&logoColor=white&style=for-the-badge" height="40" alt="mysql logo"  />
</div>

## 👨‍💻 Tecnologias e Ferramentas

| Tecnologia | Descrição |
| ---------- | --------- |
| Go         | Linguagem de programaçâo estaticamente tipada |
| Chi        | Framework Go que facilita a criação der servidores HTTP |
| MySQL      | Banco de dados relacional |
| Docker     | Plataforma de software para implantar aplicativos em containers |
| JWT        | JSON Web Token, autenticação baseada em tokens

## 📝 Descrição do projeto

A API de gerenciamento de hábitos (Habit Manager API) é uma aplicação para o controle e gerenciamento de hábitos pessoais. Com ela podemos criar e medir o progresso dos nossos hábitos.

Toda a API foi desenvolvida em **Go** e com auxílio do framework **Chi** para acelerar o desenvolvimento do projeto. Para a persistência dos dados usei **MySQL** em conjunto com **Docker** para rodar na aplicação. Em relação a autenticação para acesso aos recursos optei pelo **JWT**.

Nesse projeto tive a oportunidade de aplicar meus conhecimentos em APIs RESTful, tratamento de erros, persistência de dados usando ferramentas como SQLC, autenticação, tratamento de JSON.

## ⚡ Funcionalidades do projeto

- CRUD de hábitos
- CRUD de usuários
- Status do hábito 
- Login e Logout de usuários com autenticação JWT


## ⚙ Endpoints

- Criar hábitos: ```/habits```
- Listar hábitos: ```/habits/list```
- Atualizar hábitos ```/habits/update{id}```
- Deletar hábitos ```/habits/delete{id}```
- Criar usuários: ```/users/signup```
- Atualizar status do hábito ```/habits/updatestatus/{id}```


## 📂 Estrutura de pastas

```shell
├───cmd                         # pontos de entrada da aplicação
│   └───api                     # API aplicada
└───internal                    # Lógica de negócio
    ├───api                     # Rotas e funcionalidades das rotas
    ├───entity                  # Modelo da entidade Habit
    ├───services                # Funções de tratamento das requisições
    ├───store                   # Funcionalidades do banco de dados
    │   └───mysqlstore          # Banco de dados da aplicação
    │       ├───migrations      # Migrações do banco de dados
    │       └───queries         # Queries para manipulação do banco de dados
    ├───usecase                 # Uso de caso
    │   ├───habit               # Funções de uso de caso para hábitos
    │   └───user                # Funções de uso de caso para usuários
    ├───utils                   # Funções de utilidade
    └───validator               # Funções de validação do JSON
```
## Rodando localmente

Você precisa ter [Go](https://go.dev/) instalado em sua máquina. Versão utilizada nesse projeto ```1.22.4```.

1. Clone este repositório
2. Instale todas as dependências com os comandos: ```go mod tidy``` e também ```go get -u ./...``` na raiz do projeto
3. Inicie o servidor com o comando ```go run ./cmd/api/main.go```
4. Acesse os endpoints fornecidos

**Obs.:** Foi utilizado Docker para subir o banco de dados MySQL localmente. Se você tem o Docker instalado em sua máquina lembre-se de subir o banco de dados com o comando ```docker compose up -d```.
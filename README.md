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

A API de gerenciamento de hábitos (Habit Manager API) é uma aplicação para o controle e gerenciamento de hábitos pessoais. Com ela podemos criar e medir o progresso dos nossos hábitos como estudar, fazer exercícios, beber água etc.

Toda a API foi desenvolvida em **Go** e com auxílio do framework **Chi** para acelerar o desenvolvimento do projeto. Para a persistência dos dados usei **MySQL** em conjunto com **Docker** para rodar na aplicação. Em relação a autenticação para acesso aos recursos optei pelo JWT.

Nesse projeto tive a oportunidade de aplicar meus conhecimentos em arquitetura, testes unitários, tratamento de erros e persistência de dados usando ferramentas como SQLC.

## ⚡ Funcionalidades do projeto

- CRUD de hábitos
- CRUD de usuários
- Login e Logout de usuários


## ⚙ Endpoints

- Criar hábitos: ```/habits```
- Listar hábitos: ```/habits/list```
- Criar usuários: ```/users/signup```

## 📂 Estrutura de pastas

```shell
├───cmd
│   └───api
└───internal
    ├───api
    ├───entity
    ├───services
    ├───store
    │   └───mysqlstore
    │       ├───migrations
    │       └───queries
    ├───usecase
    │   └───user
    ├───utils
    └───validator
```

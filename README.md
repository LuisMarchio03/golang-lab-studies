# 🦫 GoLang Lab Studies

Este repositório é um laboratório técnico de experimentação, estudos avançados, provas de conceito e simulações arquiteturais utilizando **Go (Golang)**.

> **ℹ️ Aviso:** Os conteúdos aqui reunidos foram desenvolvidos entre **2022 até hoje**, mesclando estudos pessoais, testes de performance, microserviços e exploração de bibliotecas e práticas modernas do ecossistema Go.

---

## 🎯 Objetivo

- Consolidar estudos técnicos de Go com foco em sistemas backend performáticos, escaláveis e seguros.
- Criar uma base viva para revisão de conceitos e aplicação prática em projetos reais, como a arquitetura da ALOY.
- Testar arquiteturas modernas como **Hexagonal**, **Clean Architecture**, e sistemas orientados a eventos.

---

## 🧠 Tópicos Estudados

- Fundamentos da linguagem Go
- Padrões idiomáticos e boas práticas
- Manipulação avançada de erros (`errors`, `fmt.Errorf`, `wrap`, etc.)
- Concorrência com goroutines e channels
- Servidores HTTP com `net/http`, `chi`, `gin` e `fiber`
- Arquitetura em camadas / Clean Architecture
- Microsserviços com gRPC, REST, e filas (RabbitMQ, NATS)
- Go Modules & Workspaces
- CLI com Cobra
- Integração com banco de dados (PostgreSQL com `pgx`, SQLite)
- Testes (unitários, mocks e integração)

---

## ⚙️ Tecnologias Utilizadas

- **Go** (>= v1.20)
- **Chi**, **Gin**, **Fiber**
- **GORM**, **pgx**, **sqlx**
- **Cobra CLI**
- **GoMock**, **Testify**, **httptest**
- **RabbitMQ**, **Redis**, **NATS**
- **Docker** para simulação de ambientes
- **Makefile**, **Taskfile** para automações

---

## 🧪 Executando Exemplo

```bash
# Clone o repositório
git clone https://github.com/LuisMarchio03/golang-lab-studies.git

# Entre em um projeto exemplo
cd golang-lab-studies/rest-api/chi-basic-api

# Rode
go run main.go
```

---

Para projetos com dependências externas, execute go mod tidy antes da execução.

---

## 📘 Roadmap Pessoal (em progresso)

- [] Criar APIs REST com Chi e Gin

- [] Estudo de concorrência com goroutines

- [] Testes unitários e mocks

- [] Clean Architecture com injeção de dependência

- [] Microsserviços com gRPC + RabbitMQ

- [] Ferramentas CLI com Cobra

---

## 📄 Licença

MIT — Desenvolvido por Luís Gabriel Marchió Batista

Todos os códigos são voltados para uso educacional, análise técnica e aprendizado profundo de Go.


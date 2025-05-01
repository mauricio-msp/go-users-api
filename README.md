# 🚀 API RESTful em Go – Desafio CRUD In-Memory

## 📚 Descrição

Este projeto é uma **API RESTful** desenvolvida em **Go**, construída como parte do desafio do curso para consolidar os conceitos de construção de APIs.

A API realiza **operações CRUD (Create, Read, Update, Delete)** sobre um recurso em memória, ou seja, os dados são armazenados temporariamente durante a execução da aplicação, sem persistência em banco de dados.

✅ **Principais características:**

- API totalmente **in-memory** (os dados são apagados ao reiniciar a aplicação).
- Operações seguras contra **data races** utilizando mecanismos de concorrência do Go.
- Estruturada seguindo boas práticas de organização de código, separação de responsabilidades e legibilidade (Clean Code e SOLID).
- Permite gerenciar uma lista de distribuidores (ou outro recurso definido), incluindo **criação, listagem, atualização e exclusão**.

---

## 🗂️ Estrutura de pastas/arquivos

O projeto está organizado da seguinte forma:

```bash
.
├── cmd/                            # Diretório para os executáveis da aplicação
│   └── app/                        # Aplicação principal
│       └── main.go                 # Ponto de entrada da aplicação
├── internal/                       # Código interno da aplicação (não exportado para outros módulos)
│   ├── domain/                     # Pacote de domínio (modelos e interfaces)
│   │   ├── repository.go           # Interface do repositório (contrato)
│   │   └── user.go                 # Entidade User (modelo de domínio)
│   ├── dto/                        # Data Transfer Objects (estruturas de entrada/saída)
│   │   └── user_dto.go             # Data Transfer Objects relacionados a User
│   ├── repository/                 # Implementações concretas de repositórios
│   │   └── user_repository.go      # Implementação concreta da interface de repositório (in-memory ou DB)
│   ├── service/                    # Camada de serviço (regras de negócio)
│   │   └── user_service.go         # Lógica de negócio / aplicação do caso de uso
│   └── web/                        # Camada web (interface HTTP)
│       ├── handlers/               # Handlers/Controllers HTTP
│       │   └── user_handler.go     # Controladores HTTP que recebem e respondem requisições
│       └── server/                 # Configuração do servidor HTTP
│           └── server.go           # Configuração e inicialização do servidor HTTP
├── .gitignore                      # Arquivo para ignorar arquivos/pastas no controle de versão
├── api.http                        # Arquivo de requisições para testar a API via REST Client
├── go.mod                          # Gerenciamento de dependências
├── go.sum                          # Checksums das dependências do Go
└── README.md                       # README sobre o projeto
```

---

## 🏃‍♂️ Como rodar o projeto

Certifique-se de ter o Go instalado (**versão 1.18 ou superior recomendada**).

Clone o repositório:

```bash
git clone https://github.com/mauricio-msp/go-users-api.git
cd go-users-api
```

Execute o projeto:

```bash
go run cmd/app/main.go
```

A API estará rodando por padrão em [http://localhost:3333](http://localhost:3333)

---

## 🧪 Testando a API

O projeto contém um arquivo chamado **`api.http`**, que pode ser usado com a extensão **REST Client** no VS Code para testar facilmente os endpoints da API.

👉 **Como usar:**

1. Instale a extensão **REST Client** no Visual Studio Code.
2. Abra o arquivo **`api.http`** no editor.
3. Clique no botão **"Send Request"** que aparece acima de cada requisição.

O arquivo contém exemplos de requisições HTTP para os seguintes endpoints:

| Método | Rota        | Descrição               |
| ------ | ----------- | ----------------------- |
| GET    | /users      | Lista todos os usuários |
| GET    | /users/{id} | Busca usuário por ID    |
| POST   | /users      | Cria um novo usuário    |
| PUT    | /users/{id} | Atualiza um usuário     |
| DELETE | /users/{id} | Deleta um usuário       |

Cada requisição já está pré-configurada no arquivo **`api.http`** com exemplos de payloads no formato JSON.

---

## 📝 Observações

- Como a aplicação armazena os dados em memória, todos os dados serão apagados ao parar ou reiniciar a aplicação.
- Para evitar data races, foram utilizados sync.Mutex (ou outras abordagens, dependendo da implementação) para proteger as operações concorrentes de escrita e leitura no slice/mapa de dados.
- A API está preparada para uso local e pode ser facilmente estendida para utilizar um banco de dados no futuro.

Feito com :blue_heart: by [Maurício Porfírio](https://github.com/mauricio-msp).

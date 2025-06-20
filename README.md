# Go ICMS Calculator

Este repositório contém um pequeno aplicativo escrito em **Go (Golang)** cujo objetivo principal é **servir de laboratório para aprimorar minhas habilidades na linguagem**, explorando boas práticas de organização de código, testes e distribuição de pacotes.

## Descrição do projeto

O projeto calcula a alíquota de ICMS (Imposto sobre Circulação de Mercadorias e Serviços) de uma venda interestadual a partir de regras simplificadas. Apesar do domínio fiscal ser real, o propósito do código é exclusivamente **educacional**.

## Estrutura de diretórios

- `cmd/` – ponto de entrada da aplicação.
- `internal/` – pacotes internos contendo a lógica de negócio.
  - `icms/` – regras de cálculo de ICMS.
  - `sell/` – modelo que representa uma venda.
  - `random/` – utilidades diversas.
- `messaging/` – abstrações para publicação de mensagens (ex.: Pub/Sub).

## Requisitos

- Go 1.22 ou superior – instale a partir de <https://go.dev/doc/install>.

## Como executar

```bash
# clone o repositório
$ git clone https://github.com/seu-usuario/go-icsm-calculator.git
$ cd go-icsm-calculator

# baixe as dependências
$ go mod download

# execute o binário
$ go run ./cmd
```

## Próximos passos

1. Cobrir as regras de ICMS com testes unitários.
2. Separar pacotes em módulos independentes, quando fizer sentido.
3. Adicionar integração contínua (CI) para testes automáticos.

---

> **Aviso**: este projeto não deve ser usado em produção; as regras fiscais são simplificadas e não substituem consultoria especializada.

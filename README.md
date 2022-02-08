# Backend reserva online

Projeto criado para ser o backend do aplicativo

## Padrões de desenvolvimento

Segue padrões adotados neste projeto

### Package by feature

Cada feature vai ser um package separado para facilitar testes e deixar mais simples:

- config: centralização das configurações
- server: camada principal da aplicação, vai cuidar das configurações do servidor, rotas, banco de dados, etc.
- user: pacote responsavel por funcionalidades referentes ao usuário, SignIn, Login...
- company: pacote responsavel por funcionalidades das empresas. 
- schedule: pacote responsavel por funcionalidades das atividades, cadastro, inscrição, etc.

### Testes unitários

Cada funcionalidade disponibilizada deve ter testes unitários cobrindo-a.

Para rodar os testes unitários, basta rodar os comandos abaixo:

- go test ./... (para rodar todos os testes)

### Formatação do projeto

Antes de comitar, sempre executar o comando abaixo pra ver se o projeto está todo bem itentado, seguindo boas práticas do golang

go fmt ./...

## Para rodar o projeto

Para rodar o projeto precisamos primeiro instalar as dependencias utilizando o comando do go:

- go mod tidy

Após instalar precisamos subir uma aplicação do postgres, pode-se usar o docker-compose up

- docker-compose up

Após isso só rodar o projeto com o comando abaixo:

- go run cmd/main.go

## Rodando testes unitários

Para rodar e debuggar os testes unitários podemos instalar a ferramenta dlv-dap no Visual studio code, seguindo as instruções abaixo:

    First options is to fix it by running "Go: Install/Update Tools" command from the Command Palette (Linux/Windows: Ctrl+Shift+P, Mac: ⇧+⌘+P).
    Then, mark dlv & dlv-dap from the menu, and hit ok to start install/update.

Depois de instalar só apertar F5 na sua classe de teste que ele vai rodar os testes!
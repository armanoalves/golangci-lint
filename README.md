# Golangci-lint

Este repositório demonstra como configurar e utilizar o `golangci-lint` em um projeto Go.

## Estrutura do Projeto

```bash
.
├── .golangci.yml  # Configuração do golangci-lint
├── arquivo.txt     # Arquivo de exemplo
└── main.go         # Arquivo de código Go
```

## Configuração do GolangCI-Lint

O arquivo `.golangci.yml` contém a configuração utilizada para rodar o linter. Ele permite ajustar regras e ativar/desativar linters específicos de acordo com as necessidades do projeto.

## Instalação

Para instalar o `golangci-lint`, podemos utilizar algumas formas listadas na [documentação ofcial](https://golangci-lint.run/welcome/install/), mas recomendo a via instalador do Go por ser mais direta.

### Via Go
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Executando o Linter

Para rodar o `golangci-lint` no projeto é necessário usar o comando `run` e adicionar o nome do arquivo após ele:
```bash
golangci-lint run main.go
```

Isso verificará o código-fonte baseado nas regras definidas em `.golangci.yml` e indicará possíveis problemas.

## Resultado Antes e Depois

Antes de rodar o `golangci-lint`:
[Commit antes das correções](https://github.com/armanoalves/golangci-lint/tree/2e88d999a6a19376df3af3aa73a1ac24ed002526)

Depois de rodar o `golangci-lint` e corrigir os problemas identificados:
[Commit depois das correções](https://github.com/armanoalves/golangci-lint/tree/ec4f72eff28f4a4b53a2290f91e5efd728d91d1e)


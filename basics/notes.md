# Notas

## Executando um programa

- Para compilar: `go build [file].go`
- Para executar um programa compilado: `./[file-buildado]`
- Para compilar e executar um programa: `go build [file].go`

## Convenções

- Ponto e virgula opcional
- Métodos de pacotes externos são PascalCase (primeira letra maiúscula)

## Peculiaridades

- Variáveis não inicializadas possuem um valor default, exemplo:
  - strings ficam vazias: ""
  - inteiros ficam zerados: 0
  - o float fica zerado: 0.0
- Variáveis declaradas DEVEM ser utilizadas, **o programa não compila com variáveis não utilizadas**
- Variáveis possuem inferência de tipo
- Pode ser declarado a variável utilizando a sintaxe `idade := 27`

# Planilha-MTG

Planilha-MTG

## Visão geral do projeto

Projeto escrito em GO, tem como objetivo gerar uma arquivo chamado planilha.csv com as informações das cartas de Magic: The Gathering de acordo com a edição.

## Bibliotecas utilizadas

* [Magic: The Gathering SDK](https://github.com/fabiotavarespr/mtg-sdk-go)
* [Cobra](https://github.com/spf13/cobra)

## Pré-requisitos

``` 
go get github.com/fabiotavarespr/mtg-sdk-go
go get github.com/spf13/cobra
```

## Utilização

Listar edições
``` 
go run main.go sets
```

Gerar planilha
``` 
go run main.go
```

## Autor

* **Fabio Rego** - *Trabalho inicial* - [fabiotavarespr](https://github.com/fabiotavarespr)

Veja a lista de pessoas que [contribuiram](https://github.com/fabiotavarespr/Planilha-MTG/contributors) nesse projeto.
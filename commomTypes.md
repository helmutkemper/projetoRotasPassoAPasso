# Tipos comuns

A primeira coisa a ser feita é criar um tipo comum para **origem**, **destino** e 
**preço**.

Imagine o seguinte, o preço foi definido com **int** e pode mudar a qualquer momento, para
**float**. O mesmo vale para **origem** e **destino**.

Por exemplo:
```golang
  type Price int
```

Pode mudar fácil para
```golang
  type Price float64
```

Isto fará o sistema aceitar e trabalhar com float mudando uma simples linha.

Por isto, sempre crie os tipos.

```golang
  package commomTypes

  // Local de destino do trecho ou rota
  type Destination string
```

```golang
  package commomTypes

  // Local de origem do trecho ou rota
  type Origen string
```

```golang
  package commomTypes

  // Preço do trecho ou rota
  type Price int
```

Golang é uma linguagem tipada e os tipos devem ser respeitados.

A criação de tipo também permite a orientação a objeto, onde um tipo é um objeto e aceita
funções em caso de validação ou outra necessidade qualquer.

Uma característica minha é dividir todos os tipos e funções em arquivos separados, para
facilitar encontrar o tipo/função e fara facilitar ver onde falta escrever os testes.

## Fonte de dados

Para o código funcionar, temos que ter uma fonte de dados para receber a lista de rotas
e poder fazer testes.

Lembre-se, este código não é uma competição para escrever o melhor algoritmo possível,
é sobre escrever um código portável de qualidade.

Para a fonte de dados, imaginei as seguintes funções iniciais:
```golang
  AddRoute()
  GetStretchByDestination()
  GetStretchByOrigin()
```

Não faça funções desnecessárias antes da hora, mantenha o foco apenas no necessário;
Faça funções pequenas e com responsabilidade única;
Use nomes de funções explícitas, de modo a qualquer programador cansado conseguir 
entender.

Para poder fazer a rota, imaginei mais um tipo:
```golang
package commomTypes

type RouteStretch struct {
  Origin      Origin
  Destination Destination
  Price       Price
}
```

O tipo comum deve ser um módulo separado para evitar referência cíclica:
O pacote A carrega o pacote B, o pacote B carrega o pacote A novamente. Isto é coisa que
o PHP trabalha bem, mas, para o Golang encara into como falha de arquitetura.

Para começar, vamos criar um pacote **testDataSource** dessa forma:

Arquivo: typeTestDataSource.go
```golang
package testDataSource

import (
	"commomTypes"
	"sync"
)

type TestDataSource struct {
	dataList  []commomTypes.RouteStretch
	mutex sync.Mutex
}
```

Todas as rotas vão ser arquivadas em forma de array na memória. Perder os dados não é um 
problema para o exercício, no momento.
```golang
  dataList  []commomTypes.RouteStretch
```

Mutex permite travar o código enquanto uma operação estiver em progresso, evitando que 
mais de uma thread tente acessar o mesmo dado ao mesmo tempo.
```golang
  mutex sync.Mutex
```

> **sync.Mutex** é obrigatório para tratar o tipo mapa, ou o sistema vai travar.

Arquivo: funcAddRoute.go
```golang
package testDataSource

import (
	"commomTypes"
)

// Adiciona uma nova rota a fonte de dados
func (el *TestDataSource) AddRoute(
  origin commomTypes.Origin,
  destination commomTypes.Destination,
  price commomTypes.Price,
) {

	el.mutex.Lock()
	defer el.mutex.Unlock()

	if len(el.dataList) == 0 {
		el.dataList = make([]commomTypes.RouteStretch, 0)
	}

	el.dataList = append(el.dataList, commomTypes.RouteStretch{
		Origin:      origin,
		Destination: destination,
		Price:       price,
	})
}
```

Quebrar os parâmetros de entrada em linhas distintas ajudam quando a resolução do monitor 
é baixa e deixa o código mais claro.
```golang
// Adiciona uma nova rota a fonte de dados
func (el *TestDataSource) AddRoute(
  origin commomTypes.Origin,
  destination commomTypes.Destination,
  price commomTypes.Price,
) {
```

Trava o acesso ao objeto até que a função termine a sua execução:
```golang
  el.mutex.Lock()
  defer el.mutex.Unlock()
```

Todo array deve ser inicializado antes do uso sobe pena de travamento, por isto a verificação:
```golang
  if len(el.dataList) == 0 {
    el.dataList = make([]commomTypes.RouteStretch, 0)
  }
```

Adiciona o novo dado ao array:
```golang
  el.dataList = append(el.dataList, commomTypes.RouteStretch{
    Origin:      origin,
    Destination: destination,
    Price:       price,
  })
```

Como dá para perceber, a função é pequena, fácil de entender e faz exclusivamente aquilo
que se propõe a fazer, ou seja, é uma boa função.

Arquivo: funcGetStretchByDestination.go
```golang
package testDataSource

import (
  "commomTypes"
 	"errors"
)

func (el *TestDataSource) GetStretchByDestination(
  destination commomTypes.Destination,
) (
  stretchList []commomTypes.RouteStretch,
  err error,
) {

  el.mutex.Lock()
  defer el.mutex.Unlock()

  stretchList = make([]commomTypes.RouteStretch, 0)

  for _, dataLine := range el.dataList {
    if dataLine.Destination == destination {
      stretchList = append(stretchList, dataLine)
    }
  }

  if len(stretchList) == 0 {
    err = errors.New(KErroStretchNotFound)
  }

  return
}
```

O Golang permite a você retornar qualquer quantidade de parâmetros necessários, e isto é
muito bom, mas, isto pode deixar o código difícil de ser seguido por um programador 
cansado. Por isto, todos os parâmetros de retorno são nomeados.
```golang
func (el *TestDataSource) GetStretchByDestination(
  destination commomTypes.Destination,
) (
  stretchList []commomTypes.RouteStretch,
  err error,
)
```

**stretchList** foi definido no retorno, e é equivalente à **var stretchList 
[]commomTypes.RouteStretch** com as mesmas regras se aplicando, ou seja, falta a função
**make()** e ela tem de ser usada dentro da função.
```golang
  stretchList = make([]commomTypes.RouteStretch, 0)
```

Faz uma busca em todos os itens da rota e popula o array de saída, quando os dados batem.
```golang
  for _, dataLine := range el.dataList {
    if dataLine.Destination == destination {
      stretchList = append(stretchList, dataLine)
    }
  }
```

Popula a variável de erro antes do retorno.
```golang
  if len(stretchList) == 0 {
    err = errors.New(KErroStretchNotFound)
  }

  return
```

Nesse ponto, todos os textos fixos de erro, devem ser colocados em constantes, por 
motivos práticos: 
 * Impede erro de digitação;
 * Fica mais fácil traduzir, caso necessário;
 * Fica fácil documentar quais funções podem gerar a mensagem de erro.

Por isto, há o arquivo de constantes para o módulo.

Arquivo: consts.go
```golang
package testDataSource

const (
  KErroStretchNotFound = "stretch not found"
)
```

Para finalizar, uma boa prática é criar o arquivo de nome **doc.go** e colocar nele uma
explicação sobre o pacote na forma de comentário sobre o mesmo.

Arquivo: doc.go
```golang
// data source de teste, também usada para demonstração. Pode ser facilmente substituída
// por outra fonte de dados qualquer, seja, MSSQL, MongoDB, etc.
package testDataSource
```

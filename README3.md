# Teste unitário é vida

Testes unitários são usados para garantir o funcionamento do código, função por função e
impede alterações de estragar o código.

Há vários tipos de testes, com nomes bonitos, como o teste de integração, mas, a linha de
frente é o teste unitário. Se ele falhar, todos os outros vão falhar.

Porém, quando escrevemos código no dia a dia, nem sempre temos tempo de fazer todos os 
testes, por isto, sigo a seguinte prioridade:

 * Tudo o que tenho dúvida se estou certo;
 * Tudo o que for importante para o funcionamento do sistema;
 * O que tem mais valor para o cliente;
 * O resto.
 
No Golang, há duas formas principais de escrever um teste, uma função começando com o
nome **Test** ou uma função começando com o nome **Example** dentro de um arquivo 
terminando com **_test.go**.

Arquivo: funcGetStretchByOrigin_test.go
```golang
package testDataSource

import (
  "commomTypes"
  "fmt"
)

func ExampleTestDataSource_GetStretchByOrigin() {
  var err error
  var origin commomTypes.Origin
  var destination commomTypes.Destination
  var price commomTypes.Price
  var dataList []commomTypes.RouteStretch

  ds := TestDataSource{}

  origin = commomTypes.Origin("Recife")
  destination = commomTypes.Destination("Jaboatão dos Guararapes")
  price = commomTypes.Price(20)
  ds.AddRoute(origin, destination, price)

  origin = commomTypes.Origin("Recife")
  destination = commomTypes.Destination("Cabo de Santo Agostinho")
  price = commomTypes.Price(80)
  ds.AddRoute(origin, destination, price)

  dataList, err = ds.GetStretchByOrigin("Recife")
  if err != nil {
    panic(err)
  }

  for _, stretch := range dataList {
    fmt.Printf("origin: %v\n", stretch.Origin)
    fmt.Printf("destination: %v\n", stretch.Destination)
    fmt.Printf("price: %v\n\n", stretch.Price)
  }

  // Output:
  // origin: Recife
  // destination: Jaboatão dos Guararapes
  // price: 20
  //
  // origin: Recife
  // destination: Cabo de Santo Agostinho
  // price: 80
}
```

Esse é o teste feito para virar documentação, tanto na documentação on-line quanto no 
editor.

Particularmente, recomendo o editor [**Goland**](https://www.jetbrains.com/pt-br/go/) da
[**Jetbrains**](https://www.jetbrains.com/pt-br/) pelas suas funcionalidades. Basta 
apertar **Ctr** + **Q** e os comentários da função vão aparecer seguidos da função de 
exemplo.

As regras desse tipo de função são:

Nome da função: **Example** + **nome do pacote** + **_** + **nome da função a ser 
testada**:
```golang
func ExampleTestDataSource_GetStretchByOrigin() {
```

No final do arquivo deve conter o comentário **// Output:** seguido de tudo o que foi
impresso no terminal.
```golang
  // Output:
  // origin: Recife
  // destination: Jaboatão dos Guararapes
  // price: 20
  //
  // origin: Recife
  // destination: Cabo de Santo Agostinho
  // price: 80
```

A grande vantagem desse teste é servir de documentação e você tem a obrigação de escrever
bons exemplos, tanto para você do futuro, quanto para a próxima pessoa a pegar o código.

A desvantagem desse teste é o fato da saída ser muito determinista, tipo, se o teste 
escrever a hora, ele irá falar.

Arquivo: funcAddRoute_test.go
```golang
package testDataSource

import (
  "commomTypes"
  "testing"
)

func TestDataSource_AddRoute(t *testing.T) {
  var err error
  var origin commomTypes.Origin
  var destination commomTypes.Destination
  var price commomTypes.Price
  var dataList []commomTypes.RouteStretch

  ds := TestDataSource{}

  origin = commomTypes.Origin("Recife")
  destination = commomTypes.Destination("Jaboatão dos Guararapes")
  price = commomTypes.Price(20)
  ds.AddRoute(origin, destination, price)

  if len(ds.dataList) != 1 {
    t.Fail()
  }

  if ds.dataList[0].Origin != "Recife" {
    t.Fail()
  }

  if ds.dataList[0].Destination != "Jaboatão dos Guararapes" {
    t.Fail()
  }

  origin = commomTypes.Origin("Recife")
  destination = commomTypes.Destination("Cabo de Santo Agostinho")
  price = commomTypes.Price(80)
  ds.AddRoute(origin, destination, price)

  if len(ds.dataList) != 2 {
    t.Fail()
  }

  if ds.dataList[0].Origin != "Recife" {
    t.Fail()
  }

  if ds.dataList[0].Destination != "Jaboatão dos Guararapes" {
    t.Fail()
  }

  if ds.dataList[1].Origin != "Recife" {
    t.Fail()
  }

  if ds.dataList[1].Destination != "Cabo de Santo Agostinho" {
    t.Fail()
  }

  dataList, err = ds.GetStretchByDestination("Jaboatão dos Guararapes")
  if err != nil {
    panic(err)
  }

  if dataList[0].Origin != "Recife" {
    t.Fail()
  }

  if dataList[0].Destination != "Jaboatão dos Guararapes" {
    t.Fail()
  }
}
```

Esse teste não aparece na documentação e é baseado no pacote de teste do próprio Golang.

A grande vantagem dele é o fato dele permitir um teste mais profundo, usando lógica.

Basta chamar o pacote de teste, como na chamada de função abaixo e ele já começa 
inicializado:
```golang
func TestDataSource_AddRoute(t *testing.T) {
```

Depois basta fazer os seus testes e informar a falha, como em:
```golang
  if len(ds.dataList) != 1 {
    t.Fail()
  }
```

> Em algumas empresas por onde eu passei, era costume usar **panic()** quando um teste 
> falha para derrubar a pipeline de deploy do microsserviço.

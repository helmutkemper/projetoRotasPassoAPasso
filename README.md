# Projeto Golang Passo a Passo para Programadores

Pessoal, a ideia desse texto é fazer um projeto Golang passo a passo, o mais próximo 
possível de como ele ocorre no dia a dia, com novas regras no meio do projeto e regras já 
definidas mudando do nada.

As novas regras vão criar armadilhas e um código ruim vai cair nessas armadilhas, por 
isto, vou tentar explicar as regras de código o melhor possível. 

**Regra:** 
 * Todo o projeto deve ser feito em Golang nativo, sem frameworks;
 * Ao final do projeto ele deve ser portável para qualquer banco de dados, com frameworks;

## Fase 1

Uma aplicação simples de terminal:
 * Carregar um arquivo CSV com rotas de ônibus e preços para cada rota;
 * O preço de cada rota deve ser obrigatoriamente inteiro positivo;
 * A rota é definida por nomes de cidades, uma para origem e outra para destino;
 * O arquivo CSV não deve mudar ao longo do projeto (isso não é armadilha);
 * A aplicação de terminal deve ter uma interface humana para perguntar a origem e o 
   destino, apresentando em seguida a rota **mais barata**, independente da quantidade de 
   conexões;
 * Testes unitários;

Esse é um exercício para explicar regras de código e não vamos complicar com horários ou 
outras informações, a ideia é entender como fazer um código portável e não um desafio 
de algoritmos.

Crie um repositório e apresente o seu código por 
[**issues**](https://github.com/helmutkemper/golang.solid.kiss.complexity.measure/issues/new)
.

Exemplo de arquivo CSV:
```csv
Recife,Jaboatão dos Guararapes,20
Recife,Cabo de Santo Agostinho,80
Jaboatão dos Guararapes,Cabo de Santo Agostinho,10
Recife,Moreno,30
Moreno,Vitória de Santo Antão,40
Vitória de Santo Antão,Escada,50
Escada,Ipojuca,30
Escada,Cabo de Santo Agostinho,35
Moreno,Cabo de Santo Agostinho,11
```

Exemplo de rota:
```text
De: Recife
Para: Cabo de Santo Agostinho

Opções:
Recife > Cabo de Santo Agostinho: 80
Recife > Jaboatão dos Guararapes > Cabo de Santo Agostinho: 30
Recife > Moreno > Cabo de Santo Agostinho: 41

Rota mais barata:
Recife > Jaboatão dos Guararapes > Cabo de Santo Agostinho: 30
```

Boa Sorte!

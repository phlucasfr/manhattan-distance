# Manhattan Distance

[![CI](https://github.com/phlucasfr/manhattan-distance/actions/workflows/ci.yml/badge.svg)](https://github.com/phlucasfr/manhattan-distance/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/phlucasfr/manhattan-distance)

Uma implementação eficiente do cálculo da Distância de Manhattan entre dois pontos em uma matriz bidimensional, escrita em Go.

## 📦 Instalação

```bash
go get github.com/phlucasfr/manhattan-distance
```

## 💻 Execução Local
O projeto pode ser executado diretamente em seu ambiente Go com os seguintes comandos:
```
# Executar o programa principal
make run

# Alternativamente, pode rodar manualmente:
go run ./examples/main.go
```

## 🐳 Execução via Docker
O projeto pode ser executado em containers Docker para facilitar o desenvolvimento e testes:
```
# Construir a imagem e iniciar os containers
make up_build

# Parar e remover os containers
make down
```

## 🚀 Como Usar
```bash
package main

import (
    "fmt"
    "github.com/phlucasfr/manhattan-distance/manhattan"
)

func main() {
    matrix := [][]int{
        {0, 0, 0, 0},
        {0, 1, 0, 0},
        {0, 0, 0, 1},
        {0, 0, 0, 0},
    }
    fmt.Println(manhattan.ManhattanDistance(matrix)) // Saída: 3
}
```
## Requisitos da Matriz
- Deve conter exatamente dois elementos `1`.  
- Todos os demais elementos devem ser `0`.  
- Deve ser retangular (todas as linhas com o mesmo comprimento).

## 🔧 Comandos Úteis
| Comando         | Descrição                                      |
|-----------------|------------------------------------------------|
| `make run`      | Executa o exemplo principal                    |
| `make test`     | Roda testes unitários com cobertura            |
| `make bench`    | Executa benchmarks de performance              |
| `make build`    | Builda o binário do manhattan                  |
| `make up`       | Inicia os containers Docker                    |
| `make down`     | Para e remove os containers Docker             |
| `make up_build` | Builda o binário e inicia os containers Docker |

## ✅ Testes

### Casos Cobertos
- Matrizes de diferentes tamanhos (1×2 até 100×100)  
- Pontos em posições extremas  
- Cenários de alta densidade  

## 🔄 CI/CD

O projeto utiliza GitHub Actions para:

1. **Executar testes** em push e pull requests  
2. **Verificar cobertura** de código (mínimo 90%)
3. **Verificar performance** de código (mínimo 1000 ns/op)


## ✨ Destaques

- **Eficiência**: interrompe a iteração da matriz assim que encontra os dois `1`s  
- **Performance**: processa matrizes 100×100 em aproximadamente 100 nanossegundos  
- **100% de Cobertura**: todos os caminhos lógicos são testados

## 🔍 Justificativas Técnicas

- **Iteração única pela matriz**:  
  A matriz é percorrida uma única vez até encontrar os dois elementos `1`, com retorno imediato após o segundo. Isso garante complexidade ótima (O(n·m)) com desempenho máximo e mínima alocação de memória.

- **Ausência de tratamento de erro**:  
  Como o desafio especifica que a matriz **sempre terá exatamente dois `1`s**, e todos os demais elementos serão `0`, **não foram implementadas validações de entrada**. Isso mantém o código enxuto e alinhado ao escopo definido.

- **Uso da struct `Point`**:  
  Utilizar uma struct com os campos `row` e `col` torna o código mais claro e semântico, facilitando leitura e manutenção.

- **Testes unitários com alta cobertura**:  
  Os testes cobrem uma variedade de tamanhos de matriz e localizações dos `1`s, garantindo a correção da lógica principal.

- **Benchmarks para diferentes cenários**:  
  Foram incluídos benchmarks para cenários de caso médio e pior caso (matriz 100×100), validando a eficiência e estabilidade da função em situações realistas.

- **Uso de `Makefile`**:  
  Centralizamos os comandos em um `Makefile` para facilitar a execução e padronizar o uso do projeto.

- **CI com GitHub Actions**:  
  O pipeline automatizado roda testes, benchmarks e valida a cobertura mínima de 90%, garantindo qualidade contínua nas entregas.

- **Containerização com Docker**:  
  Implementei a integração com Docker para demonstrar capacidade de criar ambientes isolados e reproduzíveis, seguindo práticas modernas de DevOps.




 

# linked-list-go (en)

This repository contains an implementation of a linked list in Go, with usage examples for storing integers and player objects, along with a set of sorting functions.

## Description

A linked list is a linear data structure in which elements are stored in nodes. Each node contains a value and a pointer to the next node in the sequence. This Go implementation includes methods to create a linked list of integers or player objects, add elements, sort the list, and perform common operations such as finding elements and filtering the list.

The `structs/ordering.go` file contains a set of useful functions for sorting linked lists. It includes sorting algorithms such as insertion sort and merge sort, implemented both sequentially and in parallel.

## Benchmarks and Results

Sorting benchmarks are performed using different algorithms and configurations to measure the performance of sorting linked lists. The benchmark results provide insights into the efficiency of the algorithms and help make decisions about their usage.

The benchmark machine:

- Processor Ryzen 3 5400U, 3.3Ghz, 4 cores and 8 threads
- 8GB, DDR4 3200MHZ

The benchmark results can be found below:

| Algorithm                                          | Average Time       |
|--------------------------------------------------- |-------------------|
|**Insertion Sort:**                                |20.31 ms            |
|**Merge Sort Serial:**                             |15.27 ms           |
|**Merge Sort Serial with Insertion Sort:**          |7.85 ms            |
|**Merge Sort Parallel:**                           |6,01 ms            |
|**Merge Sort Parallel with Insertion Sort:**        |2.64 ms            |


# linked-list-go (pt-BR)

Este repositório contém uma implementação de uma lista encadeada em Go, com exemplos de uso para armazenar números inteiros e objetos de jogador, além de um conjunto de funções para ordenação.

## Descrição

A lista encadeada é uma estrutura de dados linear na qual os elementos são armazenados em nós. Cada nó possui um valor e um ponteiro para o próximo nó na sequência. Esta implementação em Go inclui métodos para criar uma lista encadeada de inteiros ou objetos de jogador, adicionar elementos, classificar a lista e realizar operações comuns, como encontrar elementos e filtrar a lista.

O arquivo `structs/ordering.go` contém uma série de funções úteis para ordenação de listas encadeadas. Ele inclui algoritmos de ordenação como o algoritmo de ordenação por inserção e o algoritmo de ordenação de mesclagem, implementados tanto sequencial quanto paralelamente.

## Benchmarks e Resultados

Os benchmarks de ordenação são realizados utilizando diferentes algoritmos e configurações para medir o desempenho da ordenação de listas encadeadas. Os resultados dos benchmarks fornecem insights sobre a eficiência dos algoritmos e ajudam a tomar decisões sobre o uso dos mesmos.

Especificação da maquina:

- Processor Ryzen 3 5400U, 3.3Ghz, 4 cores and 8 threads
- 8GB, DDR4 3200MHZ

Os resultados dos benchmarks podem ser encontrados abaixo:


| Algoritmo                                         | Média dos tempos  |
|-----------------------                            |-------------------|
|**Insertion Sort:**                                |20,31 s            |
|**Merge Sort Serial:**                             |15,27 ms           |
|**Merge Sort Serial com Ordenação por Inserção:**  |7,85 ms            |
|**Merge Sort Paralelo:**                           |6,01 ms            |
|**Merge Sort Paralelo com Ordenação por Inserção:**|2,64 ms            |
|
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rodrigoaasm/linked-list-go/example"
	"github.com/rodrigoaasm/linked-list-go/structs"
)

func main() {
	fmt.Println("Benchmark:")
	structs.Benchmark(50000)
}

func examples() {
	list := structs.NewLinkedList[int]()
	playerList := structs.NewLinkedList[example.Player]()

	randgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		n := randgen.Intn(200)
		list.Push(n)
	}
	list.Sort(func(a int, b int) bool {
		return a < b
	}, false)

	numbers := [4]int{1, 2, 3, 4}
	for _, n := range numbers {
		if _, ok := list.FindEqual(n); ok {
			fmt.Println(n, "found")
		} else {
			fmt.Println(n, "not found")
		}
	}

	playerList.Push(example.Player{Age: 21, Name: "Roberto", Points: 21})
	playerList.Push(example.Player{Age: 29, Name: "Rodrigo", Points: 22})
	playerList.Push(example.Player{Age: 28, Name: "Robert", Points: 20})
	playerList.Push(example.Player{Age: 16, Name: "Ramires", Points: 23})
	playerList.Push(example.Player{Age: 19, Name: "Reinaldo", Points: 20})
	playerList.Push(example.Player{Age: 17, Name: "Ronaldo", Points: 19})
	playerList.Sort(func(playerA example.Player, playerB example.Player) bool {
		return playerA.Points > playerB.Points
	}, false)

	adultList := playerList.
		Map(func(player example.Player) interface{} {
			return example.Person{Age: player.Age, Name: player.Name}
		}).
		Filter(func(content interface{}) bool {
			return content.(example.Person).Age >= 18
		}).
		Sort(func(personA interface{}, personB interface{}) bool {
			return personA.(example.Person).Age > personB.(example.Person).Age
		}, false)

	fmt.Println("Number List:")
	list.Print()
	fmt.Println("Players Ranking:")
	playerList.Print()
	fmt.Println("Adult Players: (sort: age)")
	adultList.Print()

	if rodrigo, ok := adultList.Find(func(player interface{}) bool {
		return player.(example.Person).Name == "Rodrigo"
	}); ok {
		fmt.Println("found:", *rodrigo)
	}
}

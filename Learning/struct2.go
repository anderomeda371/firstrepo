package main

import (
	"fmt"
)

type player struct {
	name, sport string
	age         int
}

func main() {

	player1 := player{"vahid", "Football", 34}
	fmt.Printf("%s,%s,%d", player1.sport, player1.name, player1.age)

	player2 := player{
		name:  "ali",
		age:   23,
		sport: "soccer",
	}
	fmt.Printf("\n %s,%s", player2.sport, player1.sport)

	var player3 player
	player3.name = "vahid"
	player3.sport = "basket"
	player3.age = 89

	var player4 = new(player)
	player4.name = "ppp"
	player4.sport = "ijj"
	player4.age = 88

	fmt.Printf("\n %+v", *player4)
	p1 := &player4
	(*p1).name = "ali"

	fmt.Printf("\n %+v \n %+v", player4, *p1)

	player5 := struct {
		age    int
		name   string
		family string
	}{15, "ali", "tavallae"}

	fmt.Printf("\n %v", player5)

	player6 := struct {
		age    int
		name   string
		family string
		recievers
	}{
		age:    21,
		name:   "ahmad",
		family: "modiri",
	}
	fmt.Printf("%v", player6)
	fmt.Printf("\n %v \n", compare(&player3, &player3))

	fmt.Printf("\n %v", player3)

	q := &player{"asghar", "iman", 40}
	fmt.Printf("%v,%v", &q, *q)
}

func compare(a *player, b *player) bool {

	if a.sport == b.sport && a.age == b.age && a.name == b.name {
		a.name = "changed"
		return true
	}
	return false

}

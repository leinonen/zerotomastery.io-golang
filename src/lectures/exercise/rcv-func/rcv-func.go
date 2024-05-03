//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Add", amount, "Health ->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println(player.name, "Add", amount, "Energy ->", player.energy)
}

func (player *Player) addDamage(amount uint) {
	// using unsigned integer, check if wraps around
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "Player.health ->", player.health)
}

func (player *Player) consumeEnergy(amount uint) {
	// using unsigned integer, check if wraps around
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume", amount, "Player.energy ->", player.energy)
}

func main() {
	player := Player{
		name:   "Peter",
		health: 100, maxHealth: 100,
		energy: 500, maxEnergy: 500,
	}
	fmt.Println(player)

	player.addDamage(40)
	player.consumeEnergy(200)
	fmt.Println(player)

	player.addHealth(50)
	player.addEnergy(40)

	fmt.Println(player)

	player.consumeEnergy(99999)
	fmt.Println(player)
}

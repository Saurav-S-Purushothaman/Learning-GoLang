package main

import "fmt"
import "reflect"

type Player struct{
	health int
}

// func takeDamageFromExplosion(player *Player) {
// 	fmt.Println("Player is taking damage from explosion")
// 	explosionDamage := 10
// 	player.health -= explosionDamage
// }

// now lets make this so confusing guys 

func (player *Player) takeDamageFromExplosion(dmg int){
	fmt.Println("Player is taking damage")
	player.health -= dmg
} 

// this function is same as the below function

// func takeDamageFromExplosion(player Player, dmg int) {
// 	player.health -= dmg
// }

// this is a variadic function in go

func sums(nums ...int) int {
	total := 0
	for i:=0; i < len(nums); i++{
        total += nums[i]
	}
	return total
}

func helloWorld(){
	fmt.Println("Hello world")
}

func closure(myFunc func()int){
     fmt.Println(myFunc())
}

func main(){
	player := &Player{
		health:100,
	}
	// here player is an 8 byte long integer that reference to the player struct.

	typeof := reflect.TypeOf(player)

	fmt.Println("Type of the player", typeof)
	player.takeDamageFromExplosion(10)
	fmt.Println("Player health", player.health)

	total := sums(1,2,3,4,5)

	fmt.Println("Total", total)

	x := helloWorld

	fmt.Println(reflect.TypeOf(x))
	x()

	test := func() int{
		fmt.Println("Hello")
		return 0
	}()
	fmt.Println(test)
	fmt.Printf("Type is %T\n", test)

}

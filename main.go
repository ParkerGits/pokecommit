/*
Copyright © 2022 Parker Landon parkerjlandon@gmail.com
*/

package main

import (
	// "fmt"
	"math/rand"
	"time"

	// "github.com/ParkerGits/pokecommit/handlers"
	"github.com/ParkerGits/pokecommit/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// err := handlers.EngageRandomEncounter()
	// fmt.Println(err)
	cmd.Execute()
}
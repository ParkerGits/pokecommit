/*
Copyright © 2022 Parker Landon parkerjlandon@gmail.com
*/
package main

import (
	"math/rand"
	"time"

	"github.com/ParkerGits/pokecommit/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}

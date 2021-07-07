package main

import (
	"fmt"

	"github.com/Jonny-Burkholder/timeclock/internal/tools"
)

func main() {
	um := tools.NewUserMap()
	um.Load()

	steve, err := um.LoadUser("Steve")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(steve)
}

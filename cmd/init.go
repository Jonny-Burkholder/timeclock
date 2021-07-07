package main

import (
	"fmt"

	"github.com/Jonny-Burkholder/timeclock/internal/tools"
)

func main() {
	usermap := tools.NewUserMap()
	usermap.AddUser("Steve", "5555")
	usermap.AddUser("Bilbo", "1111")
	err := usermap.Save()
	if err != nil {
		fmt.Println(err)
	}
}

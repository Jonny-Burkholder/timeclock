package main

import (
	"fmt"
	"time"

	"github.com/Jonny-Burkholder/timeclock/internal/tools"
)

func main() {
	fmt.Println(tools.DisplayTime(time.Now()))
}

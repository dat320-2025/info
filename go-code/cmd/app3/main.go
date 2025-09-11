package main

import (
	"fmt"
	"os"

	"github.com/Nejm78/demo-DAT320-2025/pkg/utils"
)

func main() {
	fmt.Println("Running app3...")
	tm := utils.NewTM()
	name := os.Args[1]
	tm.Run(name)

}

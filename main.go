package main

import (
	"fmt"
	"os"

	"yt-getembed/helpers"
)

func main() {
	if len(os.Args) != 2 {
		helpers.PrintHelp()
		return
	}

	result, err := helpers.Convert(os.Args[1])
	if err != nil {
		helpers.PrintWarning(err)
		helpers.PrintHelp()
		return
	}

	fmt.Println(result)
}

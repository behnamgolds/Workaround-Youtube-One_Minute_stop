package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const (
	YT_WATCH_URL = `watch\?v=`
	DEMO_URL     = `https://www.youtube.com/watch?v=EngW7tLk6R8`
	DEMO_RESULT  = `https://www.youtube.com/embed/EngW7tLk6R8`
)

var watch_regex = regexp.MustCompile(YT_WATCH_URL)

func main() {
	if len(os.Args) != 2 {
		printHelp()
		return
	}
	if !watch_regex.MatchString(os.Args[1]) {
		printWarning()
		printHelp()
	}
	result := watch_regex.ReplaceAllString(os.Args[1], "embed/")
	fmt.Println(result)
}

func printHelp() {
	executableName := filepath.Base(os.Args[0])
	fmt.Println("About: ")
	fmt.Printf("    %v converts the watch Youtube URL to Embed URL\n", executableName)
	fmt.Println("\nUsage: ")
	fmt.Printf("    %v %v\n", executableName, "<youtube-url>")
	fmt.Println("\nExample:")
	fmt.Printf("    %v %v\n", executableName, DEMO_URL)
	fmt.Println("\nResult:")
	fmt.Printf("    %v\n", DEMO_RESULT)
}

func printWarning() {
	fmt.Println("Warning: ")
	fmt.Printf("    Input url is malformed!\n\n")
}

package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	DEMO_URL    = `https://www.youtube.com/watch?v=EngW7tLk6R8`
	DEMO_RESULT = `https://www.youtube.com/embed/EngW7tLk6R8`
)

func PrintHelp() {
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

func PrintWarning(err error) {
	fmt.Println("Warning: ")
	fmt.Printf("    %v\n\n", err.Error())
}

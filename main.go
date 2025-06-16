package main 

import (
	"fmt"
	"os"
  "path/filepath"
)

var (
  Name string
  Value string
  Verb string
  Subject string
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	bashrcPath := filepath.Join(homeDir, ".bashrc")
  fmt.Println("Bashrc path: ", bashrcPath)

  result, err := ParseFile(bashrcPath)
  if err != nil {
    fmt.Println("Fatal error: ", err)
    return
  }

  Verb = os.Args[1]
  argErr := HandleArgs(result, &result.RawFile)
  if argErr != nil {
    fmt.Println("Error parsing args: ", err)
    return
  }
}

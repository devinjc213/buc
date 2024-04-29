package main 

import (
	"fmt"
	"os"
  "path/filepath"
)

func main() {
  args := os.Args[1:]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	bashrcPath := filepath.Join(homeDir, ".bashrc")

  result, err := ParseFile(bashrcPath)
  if err != nil {
    fmt.Println("Fatal error: ", err)
    return
  }

  argErr := ParseArgs(args, result)
  if argErr != nil {
    fmt.Println("Error parsing args: ", err)
    return
  }
}

package main 

import (
	"fmt"
	"os"
  "path/filepath"
)

var (
  Name string
  Value string
  Cmd string
  CmdType string
)

func main() {
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

  argErr := HandleArgs(result, &result.RawFile)
  if argErr != nil {
    fmt.Println("Error parsing args: ", err)
    return
  }
}

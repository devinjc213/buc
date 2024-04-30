package main 

import (
	"fmt"
	"os"
  "path/filepath"
  "flag"
)

var (
  Name string
  Value string
  Cmd string
  CmdType string
)

func init() {
  flag.StringVar(&Name, "n", "", "Name of the alias or export")
  flag.StringVar(&Name, "name", "", "Name of the alias or export")
  flag.StringVar(&Value, "v", "", "Value of the alias or export")
  flag.StringVar(&Value, "value", "", "Value of the alias or export")
}

func main() {
  flag.Parse()

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

  argErr := HandleArgs(result)
  if argErr != nil {
    fmt.Println("Error parsing args: ", err)
    return
  }
}

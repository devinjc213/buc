package main 

import (
	"fmt"
	"os"
  "path/filepath"
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

  fmt.Println("Aliases:")
  for name, alias := range result.Aliases {
    fmt.Printf("  %s: %s\n", name, alias.Value)
  }

  fmt.Println("Exports:")
  for name, export := range result.Exports {
    fmt.Printf("  %s: %s\n", name, export.Value)
  }
}

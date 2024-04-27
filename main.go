package main 

import (
	"fmt"
	"os"
  "strings"
  "bufio"
  "path/filepath"
)

type AliasExport struct {
  Value string
  LineNum int
}


func main() {
  aliases := make(map[string]AliasExport)
  exports := make(map[string]AliasExport)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	bashrcPath := filepath.Join(homeDir, ".bashrc")

	file, err := os.Open(bashrcPath)
	if err != nil {
		fmt.Println("Error opening .bashrc file:", err)
		return
	}
	defer file.Close()

  lineNum := 1
  var multiLineValue strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)

    if strings.HasPrefix(trimmedLine, "alias") {
      parts := strings.SplitN(trimmedLine, "=", 2)
      if len(parts) != 2 {
        fmt.Println("Invalid alias line:", line)
        continue
      }

      name := strings.TrimPrefix(strings.TrimSpace(parts[0]), "alias ")
      value := strings.TrimSpace(parts[1])

      if (strings.HasPrefix(value, "'") || strings.HasSuffix(value, "\"")) &&
      (!strings.HasSuffix(value, "'") || strings.HasPrefix(value, "\"")) {
        multiLineValue.Reset()
        multiLineValue.WriteString(value)
        for scanner.Scan() {
          nextLine := scanner.Text()

          if strings.HasSuffix(nextLine, "\"") || strings.HasSuffix(nextLine, "'") {
            multiLineValue.WriteString(nextLine)
            break
          }

          multiLineValue.WriteString(nextLine)
        }


        aliases[name] = AliasExport{Value: multiLineValue.String(), LineNum: lineNum}
      } else {
        aliases[name] = AliasExport{Value: value, LineNum: lineNum}
      }
    } else if strings.HasPrefix(trimmedLine, "export") {
      parts := strings.SplitN(trimmedLine, "=", 2)
      if len(parts) != 2 {
        fmt.Println("Invalid export line:", line)
        continue
      }

      name := strings.TrimPrefix(strings.TrimSpace(parts[0]), "export ")
      value := strings.TrimSpace(parts[1])

      exports[name] = AliasExport{Value: value, LineNum: lineNum}
    }

    lineNum++
	}

    // Printing aliases
  fmt.Println("\nAliases:")
  for name, alias := range aliases {
      fmt.Printf("%s = %s (Line %d)\n", name, alias.Value, alias.LineNum)
      fmt.Println()
  }

  // Printing exports
  fmt.Println("\nExports:")
  for name, export := range exports {
      fmt.Printf("%s = %s (Line %d)\n", name, export.Value, export.LineNum)
      fmt.Println()
  }

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading .bashrc file:", err)
	}
}

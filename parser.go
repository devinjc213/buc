package main

import (
  "fmt"
  "os"
  "strings"
  "errors"
)

type AliasExport struct {
  Value string
  LineNum int
}

type ParsedRc struct {
    Aliases map[string]AliasExport
    Exports map[string]AliasExport
}

func ParseFile(filePath string) (*ParsedRc, error) {
  file, err := os.ReadFile(filePath)
  if err != nil {
    return nil, errors.New("Error reading file")
  }

  fileStr := string(file)

  lines := strings.Split(fileStr, "\n")

  aliases := make(map[string]AliasExport)
  exports := make(map[string]AliasExport)

  for i, line := range lines {
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
        var quoteReached bool
        j := 1 

        for !quoteReached {
          nextLine := lines[i + j]
          if strings.HasSuffix(nextLine, "\"") || strings.HasSuffix(nextLine, "'") {
            value += nextLine
            quoteReached = true
          } else {
            value += nextLine
          }
          j++ 
        }
      }

      aliases[name] = AliasExport{Value: value, LineNum: i}
    }
  }

  return &ParsedRc{Aliases: aliases, Exports: exports}, nil
}


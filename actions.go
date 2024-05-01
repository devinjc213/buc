package main

import (
  "fmt"
  "strings"
  "os"
)

func GetName(file *ParsedRc, name string, cmdType string) (string, error) {
  var val AliasExport
  var ok bool

  if cmdType == "alias" {
    val, ok = file.Aliases[name]
  } else if cmdType == "export" {
    val, ok = file.Exports[name]
  }

  if !ok {
    return "", fmt.Errorf("Variable not found: %s", name)
  } else {
    return val.Value, nil
  }
}

func SetVar(file *string, name string, value string, cmdType string) {
  linePrefix := fmt.Sprintf("%s %s=", cmdType, name)
  lines := strings.Split(*file, "\n")

  for i, line := range lines {
    if strings.HasPrefix(strings.TrimSpace(line), linePrefix) {
      lines[i] = fmt.Sprintf("%s %s=%s", cmdType, name, value)
      *file = strings.Join(lines, "\n")
      break
    }
  }

  err := os.WriteFile("~/.bashrc_test", []byte(*file), 0644)
  if err != nil {
    fmt.Println("Error writing to file")
  }
} 

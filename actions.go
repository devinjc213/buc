package main

import "fmt"

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

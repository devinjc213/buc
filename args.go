package main

import (
  "fmt"
  "errors"
)

func ParseArgs(args []string, file *ParsedRc) (error) {
  if len(args) == 0 {
    fmt.Println("No command provided")
  }

  switch args[0] {
    case "ga":
      if len(args) < 2 {
        fmt.Print("Which alias would you like retrieve: ")
        var input string

        _, err := fmt.Scanln(&input)
        if err != nil {
          return errors.New("Error reading input: " + err.Error())
        }

        value, err := GetName(file, input, "alias")
        if err != nil {
          return errors.New("Error getting alias: " + err.Error())
        }
        
        fmt.Println(value)
    }
    case "ge":
      if len(args) < 2 {
        fmt.Print("Which export would you like retrieve: ")
        var input string

        _, err := fmt.Scanln(&input)
        if err != nil {
          return errors.New("Error reading input: " + err.Error())
        }

        value, err := GetName(file, input, "export")
        if err != nil {
          return errors.New("Error getting export: " + err.Error())
        }
        
        fmt.Println(value)
    }
  }

  return nil
}

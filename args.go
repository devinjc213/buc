package main

import (
  "fmt"
  "errors"
  "slices"
  "strings"
  "os"
)

func HandleArgs(file *ParsedRc, writeFile *string) (error) {
  shorthands := []string{"ga", "ge", "sa", "se", "la", "le", "g", "l"}
  longforms := []string{"get", "set", "list"}

  verb := os.Args[1]

  if slices.Contains(shorthands, verb) {
    parseShorthand(verb)
  } else if slices.Contains(longforms, verb) {
    parseLongform(verb)
  } else {
    return errors.New("Invalid command")
  }
  
  switch Cmd {
    case "list":
      printVars(file)
    case "get":
      getVars(file)
  }

  return nil
}

func parseLongform(verb string) (error) {
  Cmd = verb
  if len(os.Args) == 2 && Cmd != "list" {
    fmt.Print("Variable name: ")
    _, err := fmt.Scanln(&Name)
    if err != nil {
      return errors.New("Error reading input: " + err.Error())
    }
  } else if len(os.Args) > 2 {
    if Cmd == "list" {
      if os.Args[2] == "alias" {
        CmdType = "alias"
      } else if os.Args[2] == "export" {
        CmdType = "export"
      }
    } else {
      if os.Args[2] == "alias" {
        CmdType = "alias"
        if len(os.Args) == 3 {
          Name = os.Args[3]
        }
      } else if os.Args[2] == "export" {
        CmdType = "export"
        if len(os.Args) == 3 {
          Name = os.Args[3]
        }
      } else {
        Name = os.Args[2]
      }
    }
  }

  return nil
}

func parseShorthand(verb string) (error) {
  if strings.HasPrefix(verb, "g") {
    Cmd = "get"
  } else if strings.HasPrefix(verb, "s") {
    Cmd = "set"
  } else if strings.HasPrefix(verb, "l") {
    Cmd = "list"
  } else {
    return errors.New("Invalid command")
  }

  if strings.HasSuffix(verb, "a") {
    CmdType = "alias"
  } else if strings.HasSuffix(verb, "e") {
    CmdType = "export"
  } else if verb == "l" {

  } else {
    return errors.New("Invalid command")
  }

  return nil
}

func printVars(file *ParsedRc) {
  if CmdType == "" {
    fmt.Printf("Aliases:\n")
    for name, alias := range file.Aliases {
      fmt.Printf("%s = %s\n", name, alias.Value)
    }

    fmt.Printf("\n\nExports:\n")
    for name, export := range file.Exports {
      fmt.Printf("%s = %s\n", name, export.Value)
    }
  } else if CmdType == "alias" {
    fmt.Printf("Aliases:\n")
    for name, alias := range file.Aliases {
      fmt.Printf("%s = %s\n", name, alias.Value)
    }
  } else if CmdType == "export" {
    fmt.Printf("Exports:\n")
    for name, export := range file.Exports {
      fmt.Printf("%s = %s\n", name, export.Value)
    }
  }
}

func getVars(file *ParsedRc) (error) {
  if CmdType == "alias" {
    value, err := GetName(file, Name, "alias")
    if err != nil {
      return errors.New("No alias found with that name")
    }
    fmt.Println(value)
  } else if CmdType == "export" {
    value, err := GetName(file, Name, "export")
    if err != nil {
      return errors.New("No export found with that name")
    }
    fmt.Println(value)
  } else {
    value, err := GetName(file, Name, "alias")
    if err != nil {
      value, err = GetName(file, Name, "export")
      if err != nil {
        return errors.New("No alias or export found with that name")
      } else {
        CmdType = "export"
      }
    } else {
      CmdType = "alias"
    }

    fmt.Println("CmdType: ", CmdType)
    fmt.Println(value)
  }

  return nil
}

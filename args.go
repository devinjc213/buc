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

  if slices.Contains(shorthands, Verb) {
    parseShorthand()
  } else if slices.Contains(longforms, Verb) {
    parseLongform()
  } else {
    return errors.New("Invalid command")
  }
  
  switch Verb {
    case "list":
      printVars(file)
    case "get":
      getVars(file)
    case "set":
      setVar(writeFile)
  }

  return nil
}

func setVar(file *string) {
    SetVar(file, Name, Value, Subject)
}

func parseLongform() (error) {
  if len(os.Args) == 2 && Verb != "list" {
    fmt.Print("Variable name: ")
    _, err := fmt.Scanln(&Name)
    if err != nil {
      return errors.New("Error reading input: " + err.Error())
    }
  } else if len(os.Args) > 2 {
    if Verb == "list" {
      if os.Args[2] == "alias" {
        Subject = "alias"
      } else if os.Args[2] == "export" {
        Subject = "export"
      }
    } else {
      if os.Args[2] == "alias" {
        Subject = "alias"
        if len(os.Args) > 3 {
          Name = os.Args[3]
        }
        if len(os.Args) > 4 {
          Value = os.Args[4]
        }
      } else if os.Args[2] == "export" {
        Subject = "export"
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

func parseShorthand() (error) {
  if strings.HasPrefix(Verb, "g") {
    Verb = "get"
  } else if strings.HasPrefix(Verb, "s") {
    Verb = "set"
  } else if strings.HasPrefix(Verb, "l") {
    Verb = "list"
  } else {
    return errors.New("Invalid command")
  }

  if strings.HasSuffix(Verb, "a") {
    Subject = "alias"
  } else if strings.HasSuffix(Verb, "e") {
    Subject = "export"
  } else if Verb == "l" {

  } else {
    return errors.New("Invalid command")
  }

  return nil
}

func printVars(file *ParsedRc) {
  if Subject == "" {
    fmt.Printf("Aliases:\n")
    for name, alias := range file.Aliases {
      fmt.Printf("%s = %s\n", name, alias.Value)
    }

    fmt.Printf("\n\nExports:\n")
    for name, export := range file.Exports {
      fmt.Printf("%s = %s\n", name, export.Value)
    }
  } else if Subject == "alias" {
    fmt.Printf("Aliases:\n")
    for name, alias := range file.Aliases {
      fmt.Printf("%s = %s\n", name, alias.Value)
    }
  } else if Subject == "export" {
    fmt.Printf("Exports:\n")
    for name, export := range file.Exports {
      fmt.Printf("%s = %s\n", name, export.Value)
    }
  }
}

func getVars(file *ParsedRc) (error) {
  if Subject == "alias" {
    value, err := GetName(file, Name, "alias")
    if err != nil {
      return errors.New("No alias found with that name")
    }
    fmt.Println(value)
  } else if Subject == "export" {
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
        Subject = "export"
      }
    } else {
      Subject = "alias"
    }

    fmt.Println("Subject: ", Subject)
    fmt.Println(value)
  }

  return nil
}


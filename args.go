package main

import (
  "fmt"
  "errors"
  "flag"
  "slices"
  "strings"
)

var (
  listAlias bool
  listExport bool
)

func HandleArgs(file *ParsedRc) (error) {
  shorthands := []string{"ga", "ge", "sa", "se", "la", "le", "g", "l"}
  longforms := []string{"get", "set", "list"}
  fmt.Println(Name)

  args := flag.Args()
  fmt.Println(args)
  if len(args) == 0 {
    return errors.New("No command provided")
  }

  if slices.Contains(shorthands, args[0]) {
    Cmd = args[0]
    if strings.HasSuffix(Cmd, "a") {
      CmdType = "alias"
    } else if strings.HasSuffix(Cmd, "e") {
      CmdType = "export"
    } else {
      return errors.New("Invalid command")
    }
  } else if slices.Contains(longforms, args[0]) {
    listCmd := flag.NewFlagSet("list", flag.ExitOnError)
    listCmd.BoolVar(&listExport, "a", true, "Only list aliases")
    listCmd.BoolVar(&listAlias, "e", true, "Only list exports")
    getCmd := flag.NewFlagSet("get", flag.ExitOnError)
    setCmd := flag.NewFlagSet("set", flag.ExitOnError)
    getCmd.StringVar(&Name, "n", "", "Name of alias or export to retrieve")
    setCmd.StringVar(&Name, "n", "", "Name of alias or export to set")

    Cmd = args[0]

    if len(args) < 2 && Cmd != "list" {
      if err := getCmd.Parse(args[1:]); err != nil {
        return errors.New("Error parsing get command: " + err.Error())
      }

      if Name == "" {
        fmt.Print("Which alias or export would you like retrieve: ")
        _, err := fmt.Scanln(&Name)
        if err != nil {
          return errors.New("Error reading input: " + err.Error())
        }
      }
    }
  }
  
  switch Cmd {
    case "list":
      if listAlias {
        fmt.Printf("Aliases:\n")
        for name, alias := range file.Aliases {
          fmt.Printf("%s = %s\n", name, alias.Value)
        }
      }

      if listAlias && listExport {
        fmt.Printf("\n\n")
      }

      if listExport {
        fmt.Printf("Exports:\n")
        for name, export := range file.Exports {
          fmt.Printf("%s = %s\n", name, export.Value)
        }
      }
    case "get":
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
    case "ga":
      if Name == "" {
        fmt.Print("Which alias would you like retrieve: ")
        _, err := fmt.Scanln(&Name)
        if err != nil {
          return errors.New("Error reading input: " + err.Error())
        }
      }
      
      value, err := GetName(file, Name, CmdType)
      if err != nil {
        return errors.New("Error getting alias: " + err.Error())
      }

      fmt.Println(value)
    case "ge":
      if Name == "" {
        fmt.Print("Which export would you like retrieve: ")
        _, err := fmt.Scanln(&Name)
        if err != nil {
          return errors.New("Error reading input: " + err.Error())
        }
      }

      value, err := GetName(file, Name, CmdType)
      if err != nil {
        return errors.New("Error getting export: " + err.Error())
      }

    fmt.Println(value)
  }

  return nil
}

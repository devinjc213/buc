package main

import (
  "bufio"
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

func ParseFile(filePath string) (ParsedRc, error) {
  aliases := make(map[string]AliasExport)
  exports := make(map[string]AliasExport)

  file, err := os.Open(filePath)
  if err != nil {
    return ParsedRc{}, errors.New("Error opening file: " + err.Error())
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading .bashrc file:", err)
	}

  return ParsedRc{Aliases: aliases, Exports: exports}, nil
}

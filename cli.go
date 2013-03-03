package main

import (
  "fmt"
  "os"
  "github.com/jessevdk/go-flags"
)

func main() {
  var opts struct {
    Delete bool `short:"d" long:"delete" description:"delete the specified key"`
  }

  var parserOptions flags.Options = flags.Default
  parser := flags.NewParser(&opts, parserOptions)
  parser.Usage = "[options] [key] [value]"
  args, err := parser.Parse()

  if err != nil {
    os.Exit(1)
  }

  if (opts.Delete) && len(args) >= 1 {
    fmt.Printf("Deleting key...\n")
  } else if len(args) >= 2 {
    fmt.Printf("Setting key...\n")
  } else if len(args) == 1 {
    fmt.Printf("Getting key...\n")
  } else {
    fmt.Printf("Listing keys...\n")
  }
}

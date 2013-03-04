package main

import (
  "fmt"
  "os"
  "github.com/jimmycuadra/gong/gong"
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
    gong.Delete(args[0])
  } else if len(args) >= 2 {
    gong.Set(args[0], args[1])
  } else if len(args) == 1 {
    value := gong.Get(args[0])
    if len(value) != 0 {
      fmt.Println(value)
      // TODO: Copy value to clipboard.
    }
  } else if gong.IsEmpty() {
    parser.WriteHelp(os.Stderr)
  } else {
    list := gong.List()

    for _, item := range list {
      fmt.Printf("%s\n", item)
    }
  }
}

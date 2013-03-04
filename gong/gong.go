package gong

import (
  "fmt"
  "os"
  "os/user"
  "path/filepath"
  "io/ioutil"
  "encoding/json"
)

var (
  data map[string]string
  gongPath string
)

func init() {
  usr, err := user.Current()

  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }

  gongPath = filepath.Join(usr.HomeDir, ".gong")

  fileContents, err := ioutil.ReadFile(gongPath)

  if err != nil {
    data = make(map[string]string)
    return
  }

  err = json.Unmarshal(fileContents, &data)

  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }
}

func Get(key string) string {
  return data[key]
}

func Set(key, value string) {
  data[key] = value
  save()
}

func Delete(key string) {
  delete(data, key)
  save()
}

func List() []string {
  var list []string

  for key, value := range data {
    // TODO: Left pad the string so that the keys' right edges line up.
    list = append(list, key + " " + value)
  }

  return list
}

func IsEmpty() bool {
  return len(data) == 0
}

func save() {
  fileContents, err := json.Marshal(data)

  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }

  err = ioutil.WriteFile(gongPath, fileContents, 0644)

  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }
}

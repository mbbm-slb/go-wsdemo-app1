package main

import (
  "fmt"
  "mbbm-vas.com/go/wsdemo/lib1"
)

func main() {
  fmt.Println("Hello from lib1: ", lib1.Text())
  fmt.Println("Hello from lib2: ", lib1.TextFromLib2())
}
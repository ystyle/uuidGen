package main

import (
    "fmt"
    "uuid"
    "strings"
    "flag"
)

var number int;
func init(){
  flag.IntVar(&number,"n", 1, "生成的uuid数量 ,默认值为 1 .")
}


func main() {
    flag.Parse()
    if number > 1 {
      for i := 0; i < number; i++ {
        var id uuid.UUID = uuid.Rand()
        fmt.Println(strings.Replace(id.Hex(), "-", "", -1))
      }
    }else{
      var id uuid.UUID = uuid.Rand()
      fmt.Println(id.Hex())
      fmt.Println(strings.Replace(id.Hex(), "-", "", -1))
    }
}

package main

import (
  user "github.com/Yra-A/Fusion_Go/kitex_gen/user/userservice"
  "log"
)

func main() {
  svr := user.NewServer(new(UserServiceImpl))

  err := svr.Run()

  if err != nil {
    log.Println(err.Error())
  }
}

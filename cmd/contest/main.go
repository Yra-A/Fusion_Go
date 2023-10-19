package main

import (
  contest "github.com/Yra-A/Fusion_Go/kitex_gen/contest/contestservice"
  "log"
)

func main() {
  svr := contest.NewServer(new(ContestServiceImpl))

  err := svr.Run()

  if err != nil {
    log.Println(err.Error())
  }
}

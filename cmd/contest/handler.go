package main

import (
  "context"
  contest "github.com/Yra-A/Fusion_Go/kitex_gen/contest"
)

// ContestServiceImpl implements the last service interface defined in the IDL.
type ContestServiceImpl struct{}

// ContestList implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) ContestList(ctx context.Context, req *contest.ContestListRequest) (resp *contest.ContestListResponse, err error) {
  // TODO: Your code here...
  return
}

// ContestInfo implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) ContestInfo(ctx context.Context, req *contest.ContestInfoRequest) (resp *contest.ContestInfoResponse, err error) {
  // TODO: Your code here...
  return
}

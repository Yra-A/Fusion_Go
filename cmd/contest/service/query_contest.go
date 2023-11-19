package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/contest/dal/db"
	"github.com/Yra-A/Fusion_Go/cmd/contest/rpc"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"github.com/Yra-A/Fusion_Go/kitex_gen/favorite"
)

type TaskFunc func() error

type QueryContestService struct {
	ctx context.Context
}

func NewQueryContestService(ctx context.Context) *QueryContestService {
	return &QueryContestService{ctx: ctx}
}
func (s *QueryContestService) QueryContest(user_id, contest_id int32) (*contest.Contest, error) {
	c := &contest.Contest{}
	tasks := []TaskFunc{
		func() error { return s.FetchContestInfo(contest_id, c) },
		func() error { return s.FetchContactInfo(contest_id, c) },
		func() error { return s.FetchFavoriteStatus(user_id, contest_id, c) },
	}

	for _, task := range tasks {
		err := task()
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (s *QueryContestService) FetchContestInfo(contest_id int32, c *contest.Contest) error {
	dbc, err := db.QueryContestByContestId(contest_id)
	if err != nil {
		return err
	}
	c.ContestId = dbc.ContestID
	c.Title = dbc.Title
	c.Description = dbc.Description
	c.CreatedTime = dbc.CreatedTime.Unix()
	c.Field = dbc.Field
	c.Format = dbc.Format
	c.ImageUrl = dbc.ImageURL
	c.ContestCoreInfo = &contest.ContestCoreInfo{
		Deadline: dbc.Deadline,
		Fee:      dbc.Fee,
		TeamSize: &contest.TeamSize{
			Min: dbc.TeamSizeMin,
			Max: dbc.TeamSizeMax,
		},
		ParticipantRequirements: dbc.ParticipantRequirements,
		OfficialWebsite:         dbc.OfficialWebsite,
		AdditionalInfo:          dbc.AdditionalInfo,
	}
	return nil
}
func (s *QueryContestService) FetchContactInfo(contest_id int32, c *contest.Contest) error {
	dbcontestContacts, err := db.FindContestContacts(contest_id)
	if err != nil {
		return err
	}
	var contactIds []int32
	for _, v := range dbcontestContacts {
		contactIds = append(contactIds, v.ContactID)
	}
	dbContacts, err := db.QueryContactsByContactIds(contactIds)
	if err != nil {
		return err
	}
	if c.ContestCoreInfo == nil {
		c.ContestCoreInfo = &contest.ContestCoreInfo{}
	}
	c.ContestCoreInfo.Contact = make([]*contest.Contact, len(dbContacts))
	for i, v := range dbContacts {
		c.ContestCoreInfo.Contact[i] = &contest.Contact{
			Name:  v.Name,
			Phone: v.Phone,
			Email: v.Email,
		}
	}
	return nil
}

func (s *QueryContestService) FetchFavoriteStatus(user_id int32, contest_id int32, c *contest.Contest) error {
	if user_id == 0 {
		c.IsFavorite = false
		return nil
	} else {
		kresp, err := rpc.QueryFavoriteStatusByUserId(s.ctx, &favorite.QueryFavoriteStatusByUserIdRequest{UserId: user_id, ContestId: contest_id})
		if err != nil {
			return err
		}
		c.IsFavorite = kresp.IsFavorite
		return nil
	}
}

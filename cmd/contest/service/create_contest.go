package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/contest/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"time"
)

type CreateContestService struct {
	ctx context.Context
}

func NewCreateContestService(ctx context.Context) *CreateContestService {
	return &CreateContestService{ctx: ctx}
}

func (s *CreateContestService) CreateContest(c *contest.Contest) error {
	dbc := &db.Contest{
		ContestID:               c.ContestId,
		Title:                   c.Title,
		ImageURL:                c.ImageUrl,
		Field:                   c.Field,
		Format:                  c.Format,
		Description:             c.Description,
		Deadline:                c.ContestCoreInfo.Deadline,
		Fee:                     c.ContestCoreInfo.Fee,
		TeamSizeMin:             c.ContestCoreInfo.TeamSize.Min,
		TeamSizeMax:             c.ContestCoreInfo.TeamSize.Max,
		ParticipantRequirements: c.ContestCoreInfo.ParticipantRequirements,
		OfficialWebsite:         c.ContestCoreInfo.OfficialWebsite,
		AdditionalInfo:          c.ContestCoreInfo.AdditionalInfo,
		CreatedTime:             time.Unix(c.CreatedTime, 0),
	}

	if err := db.CreateOrUpdateContest(dbc); err != nil {
		return err
	}
	con := c.ContestCoreInfo.Contact
	dbcon := make([]*db.Contact, 0, len(con))
	for _, v := range con {
		dbcon = append(dbcon, &db.Contact{
			Name:  v.Name,
			Phone: v.Phone,
			Email: v.Email,
		})
	}
	if err := db.CreateOrUpdateContact(dbcon); err != nil {
		return err
	}
	conid := make([]int32, 0, len(con))
	for _, v := range dbcon {
		conid = append(conid, v.ContactID)
	}
	if err := db.ContestAddContacts(dbc.ContestID, conid); err != nil {
		return err
	}
	return nil
}

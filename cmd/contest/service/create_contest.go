package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/contest/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"gorm.io/gorm"
	"time"
)

type CreateContestService struct {
	ctx context.Context
}

func NewCreateContestService(ctx context.Context) *CreateContestService {
	return &CreateContestService{ctx: ctx}
}

func (s *CreateContestService) CreateContest(c *contest.Contest) (int32, error) {
	var contestId int32
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		contestId, err = s.createOrUpdateContest(tx, c)
		if err != nil {
			return err
		}

		contactIDs, err := s.createOrUpdateContacts(tx, c.ContestCoreInfo.Contact)
		if err != nil {
			return err
		}

		if err = s.addContestContacts(tx, contestId, contactIDs); err != nil {
			return err
		}

		return nil
	})
	return contestId, err
}

func (s *CreateContestService) createOrUpdateContest(tx *gorm.DB, c *contest.Contest) (int32, error) {
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
	return db.CreateOrUpdateContestWithTx(tx, dbc)
}

func (s *CreateContestService) createOrUpdateContacts(tx *gorm.DB, contacts []*contest.Contact) ([]int32, error) {
	dbcons := make([]*db.Contact, len(contacts))
	for i, v := range contacts {
		dbcons[i] = &db.Contact{
			Name:  v.Name,
			Phone: v.Phone,
			Email: v.Email,
		}
	}

	if err := db.CreateOrUpdateContactWithTx(tx, dbcons); err != nil {
		return nil, err
	}

	contactIDs := make([]int32, len(dbcons))
	for i, v := range dbcons {
		contactIDs[i] = v.ContactID
	}

	return contactIDs, nil
}

func (s *CreateContestService) addContestContacts(tx *gorm.DB, contestID int32, contactIDs []int32) error {
	return db.ContestAddContactsWithTx(tx, contestID, contactIDs)
}

package service

import (
	"go-api/extparty"
	"go-api/models"

	"github.com/google/go-querystring/query"
)

type IJobService interface {
	AllListJob(string, string, string, string) (*[]models.Job, error)
	ListJobByID(string) (*models.Job, error)
}

type jobService struct {
	extDans extparty.IDans
}

type QueryOptions struct {
	Description string `url:"description,omitempty"`
	Location    string `url:"location,omitempty"`
	FullTime    string `url:"full_time,omitempty"`
	Page        string `url:"size,omitempty"`
}

func NewJobService(dansConst extparty.IDans) IJobService {
	return &jobService{
		extDans: dansConst,
	}
}

func (u *jobService) AllListJob(desc, loc, fulltime, page string) (*[]models.Job, error) {
	o := QueryOptions{
		Description: desc,
		Location:    loc,
		FullTime:    fulltime,
		Page:        page,
	}
	queryParams, _ := query.Values(o)

	allJob, err := u.extDans.GetJobList(queryParams.Encode())
	if err != nil {
		return nil, err
	}

	return allJob, nil
}

func (u *jobService) ListJobByID(id string) (*models.Job, error) {
	detailJob, err := u.extDans.GetJobListByID(id)
	if err != nil {
		return nil, err
	}

	return detailJob, nil
}

package handler

import (
	"go-api/app/job/service"
	"go-api/responses"

	"github.com/labstack/echo"
)

type JobHandler struct {
	Service service.IJobService
}

func (u *JobHandler) ListJob(c echo.Context) error {
	queryDesc := c.QueryParam("description")
	queryLoc := c.QueryParam("location")
	queryFulltime := c.QueryParam("full_time")
	queryPage := c.QueryParam("page")

	respJob, err := u.Service.AllListJob(queryDesc, queryLoc, queryFulltime, queryPage)
	if err != nil {
		return err
	}

	return responses.BaseSuccess(
		c,
		"List job has been loaded",
		respJob,
	)
}

func (u *JobHandler) ListJobByID(c echo.Context) error {
	respDetailJob, err := u.Service.ListJobByID(c.Param("id"))
	if err != nil {
		return err
	}

	return responses.BaseSuccess(
		c,
		"Job ID: "+c.Param("id"),
		respDetailJob,
	)
}

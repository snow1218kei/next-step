package selectionstatus

import (
	"github.com/yuuki-tsujimura/next-step/api/domain/company"
	"github.com/yuuki-tsujimura/next-step/api/domain/jobseeker"
)

type SelectionStatus struct {
	selectionStatusID SelectionStatusID
	status            string
	jobSeekerID       jobseeker.JobSeekerID
	companyID         company.CompanyID
}

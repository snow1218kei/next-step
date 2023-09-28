package mail

import (
	"github.com/yuuki-tsujimura/next-step/api/domain/company"
	"github.com/yuuki-tsujimura/next-step/api/domain/jobopening"
	"github.com/yuuki-tsujimura/next-step/api/domain/jobseeker"
)

type Mail struct {
	mailID       MailID
	jobOpeningID jobopening.JobOpeningID
	jobSeekerID  jobseeker.JobSeekerID
	companyID    company.CompanyID
	title        string
	body         string
}

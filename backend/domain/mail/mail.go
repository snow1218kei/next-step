package mail

import (
	"github.com/yuuki-tsujimura/next-step/domain/company"
	"github.com/yuuki-tsujimura/next-step/domain/jobopening"
	"github.com/yuuki-tsujimura/next-step/domain/jobseeker"
)

type Mail struct {
	mailID       MailID
	jobOpeningID jobopening.JobOpeningID
	jobSeekerID  jobseeker.JobSeekerID
	companyID    company.CompanyID
	title        string
	body         string
}

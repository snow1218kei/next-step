package message

import (
	"time"

	"github.com/yuuki-tsujimura/next-step/api/domain/company"
	"github.com/yuuki-tsujimura/next-step/api/domain/jobopening"
	"github.com/yuuki-tsujimura/next-step/api/domain/jobseeker"
)

type Message struct {
	messageID    MessageID
	jobOpeningID jobopening.JobOpeningID
	jobSeekerID  jobseeker.JobSeekerID
	companyID    company.CompanyID
	body         string
	sentAt       time.Time
	read         string
}

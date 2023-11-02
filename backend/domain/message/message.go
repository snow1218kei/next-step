package message

import (
	"github.com/yuuki-tsujimura/next-step/domain/company"
	"github.com/yuuki-tsujimura/next-step/domain/jobopening"
	"github.com/yuuki-tsujimura/next-step/domain/jobseeker"
	"time"
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

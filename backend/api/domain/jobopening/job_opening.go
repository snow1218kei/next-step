package jobopening

import "github.com/yuuki-tsujimura/next-step/api/domain/company"

type JobOpening struct {
	jobOpeningID   JobOpeningID
	title          string
	companyID      company.CompanyID
	location       string
	position       string
	jobDescription string
	requirements   []string
	benefits       []string
}

package jobseeker

type Company struct {
	companyID   CompanyID
	name        string
	departments []CompanyDepartment
	joiningYear int
	leavingYear int
}

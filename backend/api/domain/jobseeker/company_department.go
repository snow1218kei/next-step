package jobseeker

type CompanyDepartment struct {
	departmentID CompanyDepartmentID
	projects     []Project
	joiningYear  int
	leavingYear  int
}

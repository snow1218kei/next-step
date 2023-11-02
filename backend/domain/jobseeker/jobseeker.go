package jobseeker

type JobSeeker struct {
	jobSeekerID  JobSeekerID
	name         string
	age          int
	gender       string
	educations   []Education
	careers      []Career
	portfolioURL string
	articleURL   string
}

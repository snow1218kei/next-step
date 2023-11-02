package job

import "errors"

type JobUseCase struct{}

func NewHiringUseCase() *JobUseCase {
	return &JobUseCase{}
}

func (r *JobUseCase) ApplyJob(cmd ApplyJobCommand) error {
	// find applicant
	applicant, err := r.applicantRepository.Find(cmd.ApplicantID)
	if err != nil {
		return err
	}
	if applicant == nil {
		return errors.New("applicant not found")
	}

	// find job
	job, err := r.jobRepository.Find(cmd.JobID)
	if err != nil {
		return err
	}
	if job == nil {
		return errors.New("job not found")
	}

	// check if applicant is applicable
	if !job.EligibleFor(applicant) {
		// TODO: return domain specific error
		return errors.New("applicant is not applicable")
	}

	// apply job
	job.AddApplicant(applicant)

	// save job
	if err := r.jobRepository.Save(job); err != nil {
		return err
	}

	return nil
}

package job

type ApplyJobCommand struct{}

func NewApplyJobCommand() *ApplyJobCommand {
	// TODO: validate command
	return &ApplyJobCommand{}
}

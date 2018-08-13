package types

// RunnerDescription provides the internal representation of a Runner
type RunnerDescription struct {
	RunnerType             int8
	RunnerConnectionString string
	runnerID               []byte
}

// RunnerID returns the runnerID for the runner
func (rd RunnerDescription) RunnerID() []byte {
	return rd.runnerID
}

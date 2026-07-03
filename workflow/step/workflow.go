package step

type Workflow struct {
	name string

	steps []*Step
}

func NewWorkflow(name string) *Workflow {
	return &Workflow{
		name: name,
	}
}

func (w *Workflow) Do(
	a Activity,
) *Step {

	step := &Step{
		name: a.Name(),

		activity: a,
	}

	w.steps = append(
		w.steps,
		step,
	)

	return step
}

func (w *Workflow) Name() string {
	return w.name
}

func (w *Workflow) Steps() []*Step {
	return w.steps
}

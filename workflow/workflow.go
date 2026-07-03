package workflow

type Workflow struct {
	name      string
	activites []Activity
}

func NewWorkflow(name string) *Workflow {
	return &Workflow{
		name: name,
	}
}

func (w *Workflow) Do(a Activity) *Workflow {
	w.activites = append(w.activites, a)
	return w
}

func (w *Workflow) Name() string {
	return w.name
}

func (w *Workflow) Activities() []Activity {
	return w.activites
}

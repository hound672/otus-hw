package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func insertDone(in In, done Bi) Out {
	out := make(chan interface{})
	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case vv, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-done:
					out = nil
					return
				case out <- vv:
				}
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done Bi, stages ...Stage) Out {
	if len(stages) == 0 {
		out := make(chan interface{})
		close(out)
		return out
	}

	out := stages[0](insertDone(in, done))

	for _, stage := range stages[1:] {
		stage := stage

		out = stage(insertDone(out, done))
	}
	return out
}

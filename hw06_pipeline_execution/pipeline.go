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
		nilCh := make(chan interface{})
		close(nilCh)
		return nilCh
	}

	out := stages[0](insertDone(in, done))

	for _, stage := range stages[1:] {
		stage := stage

		out = stage(insertDone(out, done))
	}

	stageResults := make([]interface{}, 0)
	for vv := range out {
		stageResults = append(stageResults, vv)
	}
	resChannel := make(chan interface{}, len(stageResults))

	for _, vv := range stageResults {
		resChannel <- vv
	}
	close(resChannel)

	return resChannel
}

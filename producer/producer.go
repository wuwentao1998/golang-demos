package producer

import "context"

type Generator interface {
	generate() Result
}

type Result interface{}

type GenerateFunc func() Result

func (g GenerateFunc) generate() Result {
	return g()
}

func Producer(f Generator) (<-chan Result, func()) {
	res := make(chan Result)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer close(res)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				res <- f.generate()
			}
		}
	}()

	return res, cancel
}

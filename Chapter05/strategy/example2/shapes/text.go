package shapes

import (
	strategy "github.com/chyidl/Go-Design-Patterns/Chapter05/strategy/example2"
	"github.com/chyidl/go-design-patterns/behavioral/strategy/example2"
)

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}

package print

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapperPrinter(t *testing.T) {
	t.Run("printer panic1", func(t *testing.T) {
		ass := assert.New(t)
		defer func() {
			if r := recover(); r != nil {
				ass.Equal(r, "machine broken")
			}
		}()
		var buf bytes.Buffer
		p := NewPrinter("printer1")
		p.Print(&buf)
		p.Print(&buf)
		p.SetPrint(Printing)
		p.Print(&buf)
		p.Print(&buf)
		expected := "Waiting...\nWaiting...\nPrinting...\npanic: machine broken\n"
		ass.Equal(expected, buf.String())
	})

	t.Run("wrapper 1", func(t *testing.T) {
		ass := assert.New(t)
		var buf bytes.Buffer
		p := NewPrinter("printer1")
		p.SetPrint(p.Wrapper(Waiting))
		p.Print(&buf)
		p.Print(&buf)
		p.SetPrint(p.Wrapper(Printing))
		p.Print(&buf)
		p.Print(&buf)
		expected := `printer1: Waiting...
printer1: Waiting...
printer1: Printing...
printer1: printer is fixed 1 times
printer1: Printing...
`
		ass.Equal(expected, buf.String())
	})
}

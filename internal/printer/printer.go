package printer

import (
	"calculator/internal/model"
)

// TODO
type Printer struct {
}

func NewPrinter() *Printer {
	p := new(Printer)
	return p
}

func (p *Printer) Print(operations []*model.Expression) { //Possible should to print all Printable

}

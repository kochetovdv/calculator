package model

import (
	"fmt"
)

type Resultable interface {
	GetResult() Result
}

type Expression struct {
	Operators []*Operator
	Numbers   []*Number
	Result    Result
}

// TODO Make result updating without IDtoUpdate using pointers
func (e *Expression) Calculate() Result {
	var IDtoUpdate []ID

	for _, op := range e.Operators {
		var num1, num2 Number
		for _, num := range e.Numbers {
			if op.ID.PrevID == num.ID.ID {
				num1 = *num
			}
			if op.ID.NextID == num.ID.ID {
				num2 = *num
			}
		}
		e.Result = *op.OpFunc(num1, num2)
		if e.Result.Error != nil {
			return e.Result
		}
		IDtoUpdate = append(IDtoUpdate, *num1.ID)
		IDtoUpdate = append(IDtoUpdate, *num2.ID)
		for _, id := range IDtoUpdate {
			for _, num := range e.Numbers {
				if *num.ID == id {
					num.Result = &e.Result
				}
			}
		}
	}
	return e.Result
}

// TODO Depricated. For logging
func (e *Expression) Print() {
	fmt.Println("Expressions: Numbers")
	for i, el := range e.Numbers {
		fmt.Printf("i: %d\tNumber:%d\tPrevID:%d\tID:%d\tNextID:%d\n", i, el.Result.Value, el.ID.PrevID, el.ID.ID, el.ID.NextID)
	}
	fmt.Println("Expressions: Operators")
	for i, el := range e.Operators {
		fmt.Printf("i: %d\tRune:%s\tPrevID:%d\tID:%d\tNextID:%d\tPriority:%d\n", i, string(el.Symbol), el.ID.PrevID, el.ID.ID, el.ID.NextID, el.Priority)
	}
}

package logic

import (
	"github.com/jadenwangtw/math-test/common/gamestructs"
)

type (
	// slot game 1201
	SlotGame1201 struct {
	}
)

func NewSlotGame1201() *SlotGame1201 {
	return &SlotGame1201{}
}

// GetSpinResult
func (g *SlotGame1201) GetSpinResult() *gamestructs.SpinResult {
	return &gamestructs.SpinResult{
		Seq: 1002,
	}
}
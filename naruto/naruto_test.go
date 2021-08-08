package naruto

import (
	"testing"
)

func TestNaruto_Clone(t *testing.T) {
	n := CreateNaruto()
	c1 := n.Clone()

	t.Run("check address", func(t *testing.T) {
		if &n == &c1 {
			t.Error("address is same")
		}
	})

	t.Run("check value", func(t *testing.T) {
		if n.GetCloneID() == c1.GetCloneID() {
			t.Error("clone id is same")
		}
	})
}

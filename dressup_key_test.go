package go_princess_precure

import (
	"testing"
)

func TestCanUseKey(t *testing.T) {
	flora := NewCureFlora()

	actual := flora.CanUseKey(TransformFlora)

	if !actual {
		t.Errorf("expect flora.canUseKey(TransformFlora) == true")
	}

	actual = flora.CanUseKey(TransformMermaid)

	if actual {
		t.Errorf("expect flora.canUseKey(TransformMermaid) == false")
	}
}

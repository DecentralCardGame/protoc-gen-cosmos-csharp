package descriptor_test

import (
	"testing"

	"github.com/DecentralCardGame/protoc-gen-cosmos-csharp/descriptor"
	"gotest.tools/v3/assert"
)

func TestDescriptorCasesAreCorrect(t *testing.T) {
	d := descriptor.FromTypeUrl(".Abc.def.oneNiceType").String()
	want := "Abc.Def.OneNiceType"

	assert.Equal(t, d, want)
}

func TestDescriptorCutsMatch(t *testing.T) {
	d := descriptor.FromTypeUrl(".abc.def.oneNiceType").CutNameSpace(descriptor.FromTypeUrl(".abc.def")).String()
	want := "OneNiceType"

	assert.Equal(t, d, want)
}

func TestDescriptorCutsNothingWithNoMath(t *testing.T) {
	d := descriptor.FromTypeUrl(".abc.def.oneNiceType").CutNameSpace(descriptor.FromTypeUrl(".dhi.def")).String()
	want := "Abc.Def.OneNiceType"

	assert.Equal(t, d, want)
}

func TestJoinJoins(t *testing.T) {
	d1 := descriptor.FromTypeUrl(".abc.def")
	d2 := descriptor.FromTypeUrl(".ghi.abc")
	d3 := d1.Join(d2)

	assert.Equal(t, d3.String(), "Abc.Def.Ghi.Abc")
}

func TestNameGivesName(t *testing.T) {
	d := descriptor.FromTypeUrl(".abc.def")

	assert.Equal(t, d.Name(), "Def")
}

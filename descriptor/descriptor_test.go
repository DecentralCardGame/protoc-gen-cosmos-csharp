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

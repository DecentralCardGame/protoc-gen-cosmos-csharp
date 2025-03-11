package descriptor

import (
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Descriptor struct {
	elems []string
}

func (d Descriptor) String() string {
	return strings.Join(d.elems, ".")
}

func (d Descriptor) Name() string {
	return d.elems[len(d.elems)-1]
}

func (d Descriptor) Join(other Descriptor) Descriptor {
	return Descriptor{elems: append(d.elems, other.elems...)}
}

func (d Descriptor) CutNameSpace(nameSpace Descriptor) Descriptor {
	nsLen := len(nameSpace.elems)
	if nsLen < len(d.elems) {
		if slices.Equal(nameSpace.elems, d.elems[:nsLen]) {
			d.elems = d.elems[nsLen:]
		}
	}

	return d
}

func FromTypeUrl(typeUrl string) Descriptor {
	elems := strings.Split(typeUrl, ".")

	if elems[0] == "" {
		elems = elems[1:]
	}

	var newElems []string
	for _, p := range elems {
		newElems = append(newElems, cases.Title(language.Und, cases.NoLower).String(p))
	}

	return Descriptor{
		elems: newElems,
	}
}

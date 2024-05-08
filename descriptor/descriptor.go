package descriptor

import (
	"slices"
	"strings"
)

type Desriptor struct {
	Elems []string
}

func (d Desriptor) String() string {
	return strings.Join(d.Elems, ".")
}

func FromTypeUrl(typeUrl string) Desriptor {
	elems := strings.Split(typeUrl, ".")

	var newElems []string
	for _, p := range elems[1:] {
		newElems = append(newElems, strings.Title(p))
	}

	return Desriptor{
		Elems: newElems,
	}
}

func (d Desriptor) CutNameSpace(nameSpace Desriptor) Desriptor {
	nsLen := len(nameSpace.Elems)
	if nsLen < len(d.Elems) {
		if slices.Equal(nameSpace.Elems, d.Elems[:nsLen]) {
			d.Elems = d.Elems[nsLen:]
		}
	}

	return d
}

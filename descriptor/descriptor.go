package descriptor

import (
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Desriptor struct {
	elems []string
}

func (d Desriptor) String() string {
	return strings.Join(d.elems, ".")
}

func FromTypeUrl(typeUrl string) Desriptor {
	elems := strings.Split(typeUrl, ".")

	var newElems []string
	for _, p := range elems[1:] {
		newElems = append(newElems, cases.Title(language.English, cases.NoLower).String(p))
	}

	return Desriptor{
		elems: newElems,
	}
}

func (d Desriptor) CutNameSpace(nameSpace Desriptor) Desriptor {
	nsLen := len(nameSpace.elems)
	if nsLen < len(d.elems) {
		if slices.Equal(nameSpace.elems, d.elems[:nsLen]) {
			d.elems = d.elems[nsLen:]
		}
	}

	return d
}

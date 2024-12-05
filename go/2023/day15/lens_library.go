package day15

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

func hash(label string) int {
	h := 0
	for _, c := range label {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}

func partOne(input []string) int {
	res := 0
	for _, line := range input {
		seqs := strings.Split(line, ",")
		for _, seq := range seqs {
			res += hash(seq)
		}
	}
	return res
}

var labelRegex = regexp.MustCompile(`^(\w*)([=-])(\d*)?$`)

type lens struct {
	label string
	len   int
}

type box struct {
	lenses []*lens
}

func (b *box) replace(l *lens) {
	idx := slices.IndexFunc(b.lenses, func(fl *lens) bool {
		return l.label == fl.label
	})
	if idx == -1 {
		b.lenses = append(b.lenses, l)
	} else {
		b.lenses[idx] = l
	}
}

func (b *box) strength() int {
	s := 0
	for i, l := range b.lenses {
		s += (i + 1) * l.len
	}
	return s
}

func (b *box) remove(label string) {
	idx := slices.IndexFunc(b.lenses, func(fl *lens) bool {
		return label == fl.label
	})
	if idx == -1 {
		return
	}
	newLenses := make([]*lens, 0)
	newLenses = append(newLenses, b.lenses[0:idx]...)
	newLenses = append(newLenses, b.lenses[idx+1:]...)
	b.lenses = newLenses
}

func partTwo(input []string) int {
	boxes := make([]*box, 256)

	for _, line := range input {
		seqs := strings.Split(line, ",")
		for _, seq := range seqs {
			labelMatches := labelRegex.FindStringSubmatch(seq)
			if len(labelMatches) < 3 {
				panic(fmt.Sprintf("failed to parse: %s", seq))
			}

			i := hash(labelMatches[1])
			if labelMatches[2] == "=" {
				l := &lens{labelMatches[1], utils.Int(labelMatches[3])}
				if boxes[i] == nil {
					boxes[i] = &box{make([]*lens, 0)}
				}
				boxes[i].replace(l)
			} else if boxes[i] != nil {
				boxes[i].remove(labelMatches[1])
			}
		}
	}

	res := 0
	for i, b := range boxes {
		if b == nil {
			continue
		}
		res += (i + 1) * b.strength()
	}

	return res
}

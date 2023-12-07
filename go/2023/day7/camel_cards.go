package day7

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

type hand struct {
	hand     []int
	strength int
	bet      int
	idx      int
}

func (h *hand) fillStrength(mct int, groups map[uint8]int) {
	if mct >= 4 {
		// five/four of a kind. strength = 6/5
		h.strength = mct + 1
	} else {
		if mct == 3 {
			// full house. strength = 4
			// three of a kind. strength = 3
			h.strength = mct + len(groups) - 1
		} else if mct == 2 {
			// two pair. strength = 2
			// one pair. strength = 1
			h.strength = mct + len(groups) - 2
		}
	}
}

var (
	handRegex       = regexp.MustCompile(`^(.*?)\s+(.*?)\s*$`)
	cardStrengthOne = map[uint8]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
	}
	cardStrengthTwo = map[uint8]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'J': 1,
	}
)

func parseHandPartOne(hs string) hand {
	handMatches := handRegex.FindStringSubmatch(hs)
	if len(handMatches) == 0 {
		panic(fmt.Sprintf("failed to parse: %s", hs))
	}

	h := hand{
		hand:     make([]int, 5),
		strength: 0,
		bet:      utils.Int(strings.TrimSpace(handMatches[2])),
	}
	cards := map[uint8]int{}
	groups := map[uint8]int{} // stores paris and above
	mc := 0
	for i, s := range strings.Split(handMatches[1], "") {
		c := s[0]
		if c >= '2' && c <= '9' {
			h.hand[i] = int(c - '2')
		} else {
			h.hand[i] = cardStrengthOne[c]
		}

		if _, ok := cards[c]; !ok {
			cards[c] = 1
		} else {
			cards[c]++
		}
		if mc < cards[c] {
			mc = cards[c]
		}
		if cards[c] > 1 {
			groups[c] = cards[c]
		}
	}

	h.fillStrength(mc, groups)

	return h
}

func partOne(input []string) int {
	hands := make([]hand, len(input))
	for i, line := range input {
		hands[i] = parseHandPartOne(line)
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.strength != b.strength {
			return a.strength - b.strength
		}
		for i, ac := range a.hand {
			if ac != b.hand[i] {
				return ac - b.hand[i]
			}
		}
		return 0
	})

	res := 0
	for i, h := range hands {
		res += (i + 1) * h.bet
	}
	return res
}

func parseHandPartTwo(hs string) hand {
	handMatches := handRegex.FindStringSubmatch(hs)
	if len(handMatches) == 0 {
		panic(fmt.Sprintf("failed to parse: %s", hs))
	}

	h := hand{
		hand:     make([]int, 5),
		strength: 0,
		bet:      utils.Int(strings.TrimSpace(handMatches[2])),
	}
	cards := map[uint8]int{}
	groups := map[uint8]int{} // stores paris and above
	mct := 0
	mnjt := 0 // max that is no joker
	jokers := 0
	for i, s := range strings.Split(handMatches[1], "") {
		c := s[0]
		if c >= '2' && c <= '9' {
			h.hand[i] = int(c - '0')
		} else {
			h.hand[i] = cardStrengthTwo[c]
		}

		isJoker := c == 'J'
		if isJoker {
			jokers++
		}

		if _, ok := cards[c]; !ok {
			cards[c] = 1
		} else {
			cards[c]++
		}
		if mnjt < cards[c] && !isJoker {
			mnjt = cards[c]
		}
		if mct < cards[c] {
			mct = cards[c]
		}
		if cards[c] > 1 {
			groups[c] = cards[c]
		}
	}

	// correction for jokers
	if jokers > 0 {
		mct = mnjt + jokers
		if len(groups) == 1 && jokers >= mct {
			groups['2'] = mct
		}

		var oi uint8
		for i := range groups {
			if i == 'J' {
				continue
			}
			oi = i
			break
		}
		if oi == 0 {
			groups['J'] = mct
		} else {
			groups[oi] += jokers
			delete(groups, 'J')
		}
	}

	h.fillStrength(mct, groups)

	return h
}

func partTwo(input []string) int {
	hands := make([]hand, len(input))
	for i, line := range input {
		hands[i] = parseHandPartTwo(line)
		hands[i].idx = i
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.strength != b.strength {
			return a.strength - b.strength
		}
		for i, ac := range a.hand {
			if ac != b.hand[i] {
				return ac - b.hand[i]
			}
		}
		return a.idx - b.idx
	})

	res := 0
	for i, h := range hands {
		res += (i + 1) * h.bet
	}
	return res
}

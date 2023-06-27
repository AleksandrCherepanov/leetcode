package main

func predictPartyVictory(senate string) string {
	R := make([]int, 0, len(senate))
	D := make([]int, 0, len(senate))

	for i, r := range senate {
		if r == 'R' {
			R = append(R, i)
		} else {
			D = append(D, i)
		}
	}

	for len(D) > 0 && len(R) > 0 {
		ld := D[0]
		D = D[1:]

		lr := R[0]
		R = R[1:]

		if lr < ld {
			R = append(R, lr+len(senate))
		} else {
			D = append(D, ld+len(senate))
		}
	}

	if len(R) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}

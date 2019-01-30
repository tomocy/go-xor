package xor

func XOR(bs ...bool) bool {
	var xor bool
	for _, b := range bs {
		if !b {
			continue
		}

		if xor {
			return false
		}
		xor = true
	}

	return xor
}

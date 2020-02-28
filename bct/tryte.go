package bct

type Tryte struct {
	Hi, Lo Trybble
}

func TryteFromInt(val int) Tryte {
	res := val / 27
	rem := val % 27
	switch {
	case rem > 13:
		return Tryte{Hi: IntToTrybble(res + 1), Lo: IntToTrybble(rem - 27)}
	case rem < -13:
		return Tryte{Hi: IntToTrybble(res - 1), Lo: IntToTrybble(rem + 27)}
	default:
		return Tryte{Hi: IntToTrybble(res), Lo: IntToTrybble(rem)}
	}
}

func (t Tryte) Value() int {
	low := t.Lo.Value()
	high := t.Hi.Value()
	return (high * 27) + low
}

func (t Tryte) String() string {
	val := t.Value()
	negative := val < 0
	n := abs(val)
	res := make([]byte, 0)

	nonDigs := map[int]byte{
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: 'D',
		6: 'C',
		7: 'B',
		8: 'A',
	}

	for n > 0 {
		rem := n % 9
		n = n / 9
		if rem > 4 {
			n++
		}
		out, ok := nonDigs[rem]
		if !ok {
			out = '0'
		}
		res = append(res, out)
	}

	for diff := 3-len(res); diff > 0; diff-- {
		res = append(res, '0')
	}

	res = reverseBuffer(res)
	if negative {
		res = negateNonaryBuffer(res)
	}
	return string(res)
}

func (t Tryte) Trybbles() [2]Trybble {
	return [2]Trybble{t.Hi, t.Lo}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func reverseBuffer(buf []byte) []byte {
	out := make([]byte, len(buf))
	i := len(buf)
	for _, b := range buf {
		i--
		out[i] = b
	}
	return out
}

func negateNonaryBuffer(buf []byte) []byte {
	out := make([]byte, len(buf))
	negNonDigs := map[byte]byte{
		'D': '4',
		'C': '3',
		'B': '2',
		'A': '1',
		'1': 'A',
		'2': 'B',
		'3': 'C',
		'4': 'D',
	}
	for i, b := range buf {
		neg, ok := negNonDigs[b]
		if !ok {
			neg = '0'
		}
		out[i] = neg
	}
	return out
}

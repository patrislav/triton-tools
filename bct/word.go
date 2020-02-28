package bct

type Word struct {
	Hi, Lo Tryte
}

const tryteMax = 27 * 27
const tryteMid = tryteMax / 2

func WordFromInt(val int) Word {
	res := val / tryteMax
	rem := val % tryteMax
	switch {
	case rem > tryteMid:
		return Word{Hi: TryteFromInt(res + 1), Lo: TryteFromInt(rem - tryteMax)}
	case rem < -tryteMid:
		return Word{Hi: TryteFromInt(res - 1), Lo: TryteFromInt(rem + tryteMax)}
	default:
		return Word{Hi: TryteFromInt(res), Lo: TryteFromInt(rem)}
	}
}

func (w Word) Value() int {
	low := w.Lo.Value()
	high := w.Hi.Value()
	return (high * tryteMax) + low
}

func (w Word) String() string {
	return w.Hi.String() + w.Lo.String()
}

func (w Word) Trytes() [2]Tryte {
	return [2]Tryte{w.Hi, w.Lo}
}

func (w Word) Trybbles() [4]Trybble {
	return [4]Trybble{w.Hi.Hi, w.Hi.Lo, w.Lo.Hi, w.Lo.Lo}
}

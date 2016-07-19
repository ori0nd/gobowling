package gobowling

type scorer struct {
	ball         int
	throws       [21]int
	currentThrow int
}

func (s *scorer) addThrow(pins int) {
	s.throws[s.currentThrow] = pins
	s.currentThrow++
}

func (s *scorer) scoreForFrame(frame int) int {
	score := 0
	s.ball = 0
	for currentFrame := 0; currentFrame < frame; currentFrame++ {
		if s.strike() {
			score += 10 + s.nextTwoBalls()
		} else if s.spare() {
			score += 10 + s.nextBall()
		} else {
			score += s.twoBallsInFrame()
		}
	}
	return score
}

func (s *scorer) strike() bool {
	if s.throws[s.ball] == 10 {
		s.ball++
		return true
	}
	return false
}

func (s *scorer) nextTwoBalls() int {
	return s.throws[s.ball] + s.throws[s.ball+1]
}

func (s *scorer) spare() bool {
	if (s.throws[s.ball] + s.throws[s.ball+1]) == 10 {
		s.ball += 2
		return true
	}
	return false
}

func (s *scorer) nextBall() int {
	return s.throws[s.ball]
}

func (s *scorer) twoBallsInFrame() int {
	value := s.throws[s.ball] + s.throws[s.ball+1]
	s.ball += 2
	return value
}

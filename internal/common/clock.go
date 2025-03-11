package common

import "time"

type Clock interface {
	Now() time.Time
}

type FakeClock struct {
	MockedNow time.Time
}

func (c FakeClock) Now() time.Time {
	return c.MockedNow
}

type RealClock struct {
	MockedNow time.Time
}

func (c RealClock) Now() time.Time {
	return time.Now()
}

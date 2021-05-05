package clock

import "fmt"

type Clock struct {
	minute int
}

const (
	hourInDay     = 24
	minutesInHour = 60
	minuteInDay   = minutesInHour * hourInDay
)

func New(h, m int) Clock {
	m += h * minutesInHour
	m %= minuteInDay

	if m < 0 {
		m += minuteInDay
	}

	return Clock{m}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/minutesInHour, c.minute%minutesInHour)
}

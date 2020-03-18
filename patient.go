package covid19

import "time"

type Patient struct {
	ID   string
	Name string

	Tests []time.Time
}

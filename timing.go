package mssql

import (
	"fmt"
	"sync"
	"time"
)

// Timings is a synchronized map of timing measurments keyed by connectParams
var Timings sync.Map

// Timing is a record of the timing of various driver lifecycle events
type Timing struct {
	Start   time.Time       `json:"start"`
	Dial    time.Duration   `json:"dial"`
	Auth    time.Duration   `json:"auth"`
	Total   time.Duration   `json:"total"`
	Queries []time.Duration `json:"queries"`
	Err     error           `json:"err"`
}

func (t *Timing) start() {
	t.Start = time.Now()
}

func (t *Timing) dial() {
	t.Dial = time.Since(t.Start)
}

func (t *Timing) auth() {
	t.Auth = time.Since(t.Start)
}

func (t *Timing) done() {
	t.Total = time.Since(t.Start)
}

func (t *Timing) TimeQuery(q func() error) error {
	qStart := time.Now()
	err := q()
	t.Queries = append(t.Queries, time.Since(qStart))
	return err
}

func (t Timing) String() string {
	return fmt.Sprintf("Dial:%v\nAuth:%v\nTotal:%v\n", t.Dial, t.Auth, t.Total)
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lock

import (
	"sync"
	"time"

	"github.com/armon/go-metrics"
)

// Timer provides a map of named timers which is safe for concurrent use. Each
// timer is created using time.AfterFunc which will be triggered once the timer
// fires.
type Timer struct {

	// timers is a mapping of timers which represent when a lock TTL will
	// expire. The lock should be used for all access.
	timers map[string]*time.Timer
	lock   sync.RWMutex
}

// NewTimer initializes a new Timer.
func NewTimer() *Timer {
	return &Timer{timers: make(map[string]*time.Timer)}
}

// Get returns the timer with the given id or nil.
func (t *Timer) Get(id string) *time.Timer {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.timers[id]
}

// Set stores the timer under given id. If tm is nil the timer with the given
// id is removed.
func (t *Timer) Set(id string, tm *time.Timer) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if tm == nil {
		delete(t.timers, id)
	} else {
		t.timers[id] = tm
	}
}

// Del removes the timer with the given id.
func (t *Timer) Del(id string) {
	t.Set(id, nil)
}

// Len returns the number of registered timers.
func (t *Timer) len() int {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return len(t.timers)
}

// ResetOrCreate sets the ttl of the timer with the given id or creates a new
// one if it does not exist.
func (t *Timer) ResetOrCreate(id string, ttl time.Duration, afterFunc func()) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if tm := t.timers[id]; tm != nil {
		tm.Reset(ttl)
		return
	}
	t.timers[id] = time.AfterFunc(ttl, afterFunc)
}

// Stop stops the timer with the given id and removes it.
func (t *Timer) Stop(id string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if tm := t.timers[id]; tm != nil {
		tm.Stop()
		delete(t.timers, id)
	}
}

// StopAll stops and removes all registered timers.
func (t *Timer) StopAll() {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, tm := range t.timers {
		tm.Stop()
	}
	t.timers = make(map[string]*time.Timer)
}

// EmitMetrics is a long-running routine used to update a number of server
// periodic metrics.
func (t *Timer) EmitMetrics(shutdownCh chan struct{}) {
	for {
		select {
		case <-time.After(time.Second):
			metrics.SetGauge([]string{"variable", "lock_timer", "num"}, float32(t.len()))
		case <-shutdownCh:
			return
		}
	}
}

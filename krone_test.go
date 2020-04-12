package krone

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	now := time.Now()
	d := 2 * time.Second
	k := New(d)

	k.Do(func() {
		t.Logf("Stop: %v", time.Since(now))
	})
	got := time.Since(now)

	if got < d {
		t.Errorf("expected greater than %v, got %v", d, got)
	}
}

func TestNewCanceled(t *testing.T) {
	now := time.Now()
	d := 5 * time.Second
	k := New(d)

	go func() {
		time.Sleep(2 * time.Second)
		k.Stop()
	}()

	k.Do(func() {
		t.Logf("Stop: %v", time.Since(now))
	})
	got := time.Since(now)

	if got > d {
		t.Errorf("expected leter than %v, got %v", d, got)
	}
}

func TestFromTime(t *testing.T) {
	now := time.Now()
	d := 2 * time.Second
	ti := now.Add(d)
	k := FromTime(ti)
	k.Do(func() {
		t.Logf("Stop: %v", time.Since(now))
	})
	got := time.Since(now)

	if got < d {
		t.Errorf("expected greater than %v, got %v", d, got)
	}
}

func TestEvery(t *testing.T) {
	now := time.Now()
	d := 1 * time.Second

	k := New(d)
	var counter int
	k.Every(func() {
		counter++
		t.Logf("Stop: %v", time.Since(now))
		if counter == 3 {
			k.Stop()
		}
	})

	got := time.Since(now)
	if counter != 3 {
		t.Errorf("expected 3, got %d", counter)
	}

	if got < d {
		t.Errorf("expected greater than %v, got %v", d, got)
	}

}

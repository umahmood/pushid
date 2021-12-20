package pushid

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	// ErrInvalidIDLength invalid ID length was generated
	ErrInvalidIDLength = errors.New("id length is not equal to 20")
	// ErrConvertingTimestamp unable to convert timestamp
	ErrConvertingTimestamp = errors.New("unable to convert entire timestamp")
)

type lockedRandSource struct {
	lock sync.Mutex
	src  rand.Source
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *lockedRandSource) Int63() int64 {
	r.lock.Lock()
	ret := r.src.Int63()
	r.lock.Unlock()
	return ret
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state
func (r *lockedRandSource) Seed(seed int64) {
	r.lock.Lock()
	r.src.Seed(seed)
	r.lock.Unlock()
}

// PushID instance
type PushID struct {
	random        *rand.Rand
	lastPushTime  float64
	lastRandChars []byte
}

// New creates a new PushID instance
func New() *PushID {
	return &PushID{
		random: rand.New(
			&lockedRandSource{
				src: rand.NewSource(time.Now().UnixNano()),
			},
		),
		lastRandChars: make([]byte, 12),
	}
}

// Generate a push id
func (p *PushID) Generate() (string, error) {
	var (
		chars          = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
		id             = ""
		timestampChars = [8]byte{}
		now            = float64(time.Duration(time.Now().UnixNano()).Milliseconds())
		duplicateTime  = (now == p.lastPushTime)
	)
	p.lastPushTime = now
	for i := 7; i >= 0; i-- {
		timestampChars[i] = chars[int64(now)%64]
		now = math.Floor(now / 64)
	}
	if now != 0 {
		return "", ErrConvertingTimestamp
	}
	for _, b := range timestampChars {
		id += string(b)
	}
	if !duplicateTime {
		for i := 0; i < 12; i++ {
			p.lastRandChars[i] = byte(math.Floor(p.random.Float64() * 64))
		}
	} else {
		// if the timestamp hasn't changed since last push, use the same random
		// number, except incremented by 1.
		i := 11
		for ; i >= 0 && p.lastRandChars[i] == 63; i-- {
			p.lastRandChars[i] = 0
		}
		p.lastRandChars[i]++
	}
	for i := 0; i < 12; i++ {
		id += string(chars[p.lastRandChars[i]])
	}
	if len(id) != 20 {
		return "", ErrInvalidIDLength
	}
	return id, nil
}

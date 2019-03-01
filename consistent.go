// Package consistent borrows ideas from "stathat.com/c/consistent"
// simplify it for performance, no concurrence guarantee any more!
package consistent

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
)

type uints []uint32

// Len returns the length of the uints array.
func (x uints) Len() int { return len(x) }

// Less returns true if element i is less than element j.
func (x uints) Less(i, j int) bool { return x[i] < x[j] }

// Swap exchanges elements i and j.
func (x uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

var (
	// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
	ErrEmptyCircle = errors.New("empty circle")
	// ErrNotSorted is the error returned when trying to get an element when the consistent circle is not sorted.
	ErrNotSorted = errors.New("not sorted")
)

// Consistent holds the information about the members of the consistent hash circle.
type Consistent struct {
	circle       map[uint32]string
	sortedHashes uints
	state        state
}

// state : state machine of Consistent
type state int

const (
	// stateBad Before New, or in Sorting
	stateBad state = iota
	// stateAdding After New, Before Sorting
	stateAdding
	// stateSorted After Sorting
	stateSorted
)

// New creates a new Consistent object.
func New() *Consistent {
	c := new(Consistent)
	c.circle = make(map[uint32]string)
	c.state = stateAdding
	return c
}

// eltKey generates a string key for an element with an index.
func (c *Consistent) eltKey(elt string, idx int) string {
	return strconv.Itoa(idx) + elt
}

// Add inserts a string element in the consistent hash replicas times.
func (c *Consistent) Add(elt string, replicas int) {
	if c.state != stateAdding {
		panic("state not stateAdding")
	}
	for i := 0; i < replicas; i++ {
		c.circle[c.hashKey(c.eltKey(elt, i))] = elt
	}
}

// Get returns an element close to where name hashes to in the circle.
func (c *Consistent) Get(name string) (string, error) {
	if c.state != stateSorted {
		return "", ErrNotSorted
	}
	if len(c.circle) == 0 {
		return "", ErrEmptyCircle
	}
	key := c.hashKey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]], nil
}

func (c *Consistent) search(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// SortHashes before get
func (c *Consistent) SortHashes() {
	c.state = stateBad
	defer func() {
		c.state = stateSorted
	}()
	c.sortedHashes = make(uints, 0, len(c.circle))
	for k := range c.circle {
		c.sortedHashes = append(c.sortedHashes, k)
	}
	sort.Sort(c.sortedHashes)
}

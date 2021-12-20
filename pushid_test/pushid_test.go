package pushid_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/umahmood/pushid"
)

func TestGeneratePushIDs(t *testing.T) {
	var (
		output = []string{}
		id     = pushid.New()
	)
	for i := 0; i < 10; i++ {
		thisID, err := id.Generate()
		if err != nil {
			t.Errorf("%v", err)
		}
		output = append(output, thisID)
		time.Sleep(100 * time.Millisecond)
	}
	ids := make([]string, len(output))
	copy(ids, output)
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
	sort.Strings(ids)
	compareSlices := func(a, b []string) bool {
		for i, item := range a {
			if item != b[i] {
				return false
			}
		}
		return true
	}
	if !compareSlices(output, ids) {
		t.Errorf("error generating pushids got %v want %v", output, ids)
	}
}

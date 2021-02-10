package cchannel

import (
	"fmt"
	"testing"
)

type st struct {
	num   int
	value string
}

func TestCchannel(t *testing.T) {
	var cchan Cchannel
	cchan.NewChan(10)

	expected := 0
	for i := 0; i < 15; i++ {
		cchan.Add(st{i, "Next string"})
		expected = i % 9
	}

	for {
		if val, valid := cchan.Get(); valid {
			fmt.Println(val, cchan.Free)
			cval := val.(st)
			if cval.num != expected {
				t.Errorf("Got wrong item %d - expected %d", cval.num, expected)
			}
			expected++
		} else {
			break
		}
	}

	if cchan.Free != cchan.Capacity {
		t.Errorf("Get did not read all items")
	}

}

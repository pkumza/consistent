package consistent

import (
	"strconv"
	"testing"
)

type testCase struct {
	Bucket string
	Weight int
}

var testCases = []testCase{
	testCase{
		Bucket: "10.115.46.1:5663",
		Weight: 100,
	},
	testCase{
		Bucket: "10.115.46.2:5663",
		Weight: 200,
	},
	testCase{
		Bucket: "10.115.46.3:5663",
		Weight: 300,
	},
	testCase{
		Bucket: "10.115.46.4:5663",
		Weight: 400,
	},
}

func Test_Get(t *testing.T) {
	c := New()
	for _, tc := range testCases {
		c.Add(tc.Bucket, tc.Weight)
	}
	c.SortHashes()
	sum := make(map[string]int)
	for i := 0; i < 100000; i++ {
		bucket, err := c.Get("" + strconv.Itoa(i))
		if err != nil {
			t.Fatalf("Err %v", err)
		}
		sum[bucket]++
	}
	for _, tc := range testCases {
		t.Logf("Bucket:%s, Weight:%d, Sum:%d\n", tc.Bucket, tc.Weight, sum[tc.Bucket])
	}
}

func Benchmark_Get(b *testing.B) {
	c := New()
	for _, tc := range testCases {
		c.Add(tc.Bucket, tc.Weight)
	}
	c.SortHashes()
	for i := 0; i < b.N; i++ {
		_, err := c.Get("" + strconv.Itoa(i))
		if err != nil {
			b.Fatalf("Err %v", err)
		}
	}
}

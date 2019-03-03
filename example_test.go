package consistent

import (
	"fmt"
	"testing"
)

func Test_Simplest(t *testing.T) {
	c := New()
	c.Add("Bucket1", 20)
	c.Add("Bucket2", 40)
	c.SortHashes()
	bucket, _ := c.Get("apple")
	fmt.Println(bucket) // Bucket2
	bucket, _ = c.Get("banana")
	fmt.Println(bucket) // Bucket1
	bucket, _ = c.Get("pear")
	fmt.Println(bucket) // Bucket2
}

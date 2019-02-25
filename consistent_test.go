package consistent

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"testing"
)

func Test_Mazinag(t *testing.T) {
	t.Logf("呵呵呵")

	fmt.Printf("%08x\n", crc32.ChecksumIEEE([]byte("2"))%10)
	fmt.Printf("%08x\n", crc32.ChecksumIEEE([]byte("3"))%10)
	fmt.Printf("%08x\n", crc32.ChecksumIEEE([]byte("4"))%10)
	fmt.Printf("%08x\n", crc32.ChecksumIEEE([]byte("5"))%10)

}

func Test_New(t *testing.T) {
	c := New()
	c.Add("10.115.46.2:5668", 100)
	c.Add("10.115.46.2:5666", 200)
	c.Add("10.115.46.2:5667", 300)
	c.Add("10.115.46.2:5665", 400)
	c.SortHashes()
	// t.Logf("Rep %v", c.sortedHashes)
	cnt1 := 0
	cnt2 := 0
	cnt3 := 0
	cnt4 := 0
	for i := 0; i < 100000; i++ {
		keng, err := c.Get(strconv.Itoa(i))
		if err != nil {
			t.Fatalf("Err %v", err)
		}
		if keng == "10.115.46.2:5668" {
			cnt1++
		}
		if keng == "10.115.46.2:5667" {
			cnt2++
		}
		if keng == "10.115.46.2:5666" {
			cnt3++
		}
		if keng == "10.115.46.2:5665" {
			cnt4++
		}
	}
	fmt.Printf("%v, %v, %v, %v", cnt1, cnt2, cnt3, cnt4)
}

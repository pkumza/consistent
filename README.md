# consistent

Consistent hash package for Go.

Performance is better than `stathat.com/c/consistent`, but no concurrence guarantee.

NumOfReplicas of each shard is adjustable.

## Installation

```bash
go get github.com/pkumza/consistent@latest
```

## Docs

https://godoc.org/github.com/pkumza/consistent

## Example

```go
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
```
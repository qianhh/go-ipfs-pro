package gc

import (
	bloom "github.com/ipfs/bbloom"
	"github.com/ipfs/go-cid"
)

type Set interface {
	Visit(c cid.Cid) bool
	Has(c cid.Cid) bool
	Add(c cid.Cid)
}

func NewBloomFilter(bloomSize, hashCount int) (*BloomFilter, error) {
	bl, err := bloom.New(float64(bloomSize), float64(hashCount))
	if err != nil {
		return nil, err
	}
	return &BloomFilter{
		filter: bl,
	}, nil
}

func NewDefaultBloomFilter() (*BloomFilter, error) {
	return NewBloomFilter(1<<28, 7)
}

type BloomFilter struct {
	filter *bloom.Bloom
}

func (bf *BloomFilter) Visit(c cid.Cid) bool {
	bf.Add(c)
	// always return true, because bloom filter has FalsePositive
	return true
}

func (bf *BloomFilter) Has(c cid.Cid) bool {
	return bf.filter.Has(c.Bytes())
}

func (bf *BloomFilter) Add(c cid.Cid) {
	bf.filter.Add(c.Bytes())
}

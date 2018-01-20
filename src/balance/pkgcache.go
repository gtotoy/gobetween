/**
 * pkgcache.go - Packages Cache balance impl
 *
 * @author Gustavo Totoy <gtotoy@gmail.com>
 */

package balance

import (
	"errors"
	"hash/fnv"

	"../core"
)

/**
 * Pkgcache balancer
 */
type PkgcacheBalancer struct {
	/* Packages required by Lambda Function */
	ReqPkgs   []string
	Threshold int
}

func h1(s string) uint32 {
	hf := fnv.New32()
	hf.Write([]byte(s))
	return hf.Sum32()
}

func h2(s string) uint32 {
	hf := fnv.New32a()
	hf.Write([]byte(s))
	return hf.Sum32()
}

/**
 * Elect backend using pkgcache strategy
 */
func (b *PkgcacheBalancer) Elect(context core.Context, backends []*core.Backend) (*core.Backend, error) {

	if len(backends) == 0 {
		return nil, errors.New("Can't elect backend, Backends empty")
	}

	largestPkg := "" // @TODO(Gus): Calculate it using b.ReqPkgs
	targetIndex1 := h1(largestPkg)%uint32(len(backends)) + 1
	targetIndex2 := h2(largestPkg)%uint32(len(backends)) + 1

	targetIndex := targetIndex2
	if backends[targetIndex1].Load < backends[targetIndex2].Load {
		targetIndex = targetIndex1
	}

	if backends[targetIndex].Load > b.Threshold { // Find backend with min Load
		for i, e := range backends {
			if e.Load < backends[targetIndex].Load {
				targetIndex = uint32(i)
			}
		}
	}

	backend := backends[targetIndex]

	return backend, nil
}

/**
 * pkgcache.go - Packages Cache balance impl
 *
 * @author Gustavo Totoy <gtotoy@gmail.com>
 */

package balance

import (
	"errors"

	"../core"
)

/**
 * Pkgcache balancer
 */
type PkgcacheBalancer struct {

	/* Current backend position */
	current int
}

/**
 * Elect backend using pkgcache strategy
 */
func (b *PkgcacheBalancer) Elect(context core.Context, backends []*core.Backend) (*core.Backend, error) {

	if len(backends) == 0 {
		return nil, errors.New("Can't elect backend, Backends empty")
	}

	if b.current >= len(backends) {
		b.current = 0
	}

	backend := backends[b.current]
	b.current += 1

	return backend, nil
}

//
// Copyright (c) 2022 Cloudflare, Inc.
//
// Licensed under Apache 2.0 license found in the LICENSE file
// or at http://www.apache.org/licenses/LICENSE-2.0
//

package main

import "time"

type tokenBucket struct {
	ch      chan struct{}
	rate    time.Duration
	closeCh chan struct{}
}

func (t *tokenBucket) getToken() bool {
	select {
	case <-t.ch:
		return true
	default:
		return false
	}
}

func (t *tokenBucket) creator() {
	ticker := time.NewTicker(t.rate)

	for {
		<-ticker.C

		select {
		case t.ch <- struct{}{}:
			continue
		case <-t.closeCh:
			return

		default:
			continue
		}
	}
}

func (t *tokenBucket) stop() {
	close(t.closeCh)
}

func newTokenBucket(size uint, rate time.Duration) *tokenBucket {
	tb := &tokenBucket{
		rate:    rate,
		ch:      make(chan struct{}, size),
		closeCh: make(chan struct{}),
	}

	go tb.creator()

	return tb
}

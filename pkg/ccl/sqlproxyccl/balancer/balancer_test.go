// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package balancer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach/pkg/ccl/sqlproxyccl/tenant"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/timeutil"
	"github.com/stretchr/testify/require"
)

func TestBalancer(t *testing.T) {
	defer leaktest.AfterTest(t)()

	b := NewBalancer()

	t.Run("no pods", func(t *testing.T) {
		pod, err := b.SelectTenantPod([]*tenant.Pod{})
		require.EqualError(t, err, ErrNoAvailablePods.Error())
		require.Nil(t, pod)
	})

	t.Run("few pods", func(t *testing.T) {
		pod, err := b.SelectTenantPod([]*tenant.Pod{{Addr: "1"}, {Addr: "2"}})
		require.NoError(t, err)
		require.Contains(t, []string{"1", "2"}, pod.Addr)
	})
}

func TestRebalancerQueue(t *testing.T) {
	defer leaktest.AfterTest(t)()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	q, err := newRebalancerQueue(ctx)
	require.NoError(t, err)

	// Use a custom time source for testing.
	t0 := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeSource := timeutil.NewManualTime(t0)

	// Create rebalance requests for the same connection handle.
	conn1 := &testBalancerConnHandle{}
	req1 := &rebalanceRequest{
		createdAt: timeSource.Now(),
		conn:      conn1,
		dst:       "foo1",
	}
	timeSource.Advance(5 * time.Second)
	req2 := &rebalanceRequest{
		createdAt: timeSource.Now(),
		conn:      conn1,
		dst:       "foo2",
	}
	timeSource.Advance(5 * time.Second)
	req3 := &rebalanceRequest{
		createdAt: timeSource.Now(),
		conn:      conn1,
		dst:       "foo3",
	}

	// Enqueue in a specific order. req3 overrides req1; req2 is a no-op.
	q.enqueue(req1)
	q.enqueue(req3)
	q.enqueue(req2)
	require.Len(t, q.elements, 1)
	require.Equal(t, 1, q.queue.Len())

	// Create another request.
	conn2 := &testBalancerConnHandle{}
	req4 := &rebalanceRequest{
		createdAt: timeSource.Now(),
		conn:      conn2,
		dst:       "bar1",
	}
	q.enqueue(req4)
	require.Len(t, q.elements, 2)
	require.Equal(t, 2, q.queue.Len())

	// Dequeue the items.
	item, err := q.dequeue(ctx)
	require.NoError(t, err)
	require.Equal(t, req3, item)
	item, err = q.dequeue(ctx)
	require.NoError(t, err)
	require.Equal(t, req4, item)
	require.Empty(t, q.elements)
	require.Equal(t, 0, q.queue.Len())

	// Cancel the context. Dequeue should return immediately with an error.
	cancel()
	req4, err = q.dequeue(ctx)
	require.EqualError(t, err, context.Canceled.Error())
}

func TestRebalancerQueueBlocking(t *testing.T) {
	defer leaktest.AfterTest(t)()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	q, err := newRebalancerQueue(ctx)
	require.NoError(t, err)

	reqCh := make(chan *rebalanceRequest, 10)
	go func() {
		for {
			req, err := q.dequeue(ctx)
			if err != nil {
				break
			}
			reqCh <- req
		}
	}()

	// Use a custom time source for testing.
	t0 := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeSource := timeutil.NewManualTime(t0)

	const reqCount = 100
	for i := 0; i < reqCount; i++ {
		req := &rebalanceRequest{
			createdAt: timeSource.Now(),
			conn:      &testBalancerConnHandle{},
			dst:       fmt.Sprint(i),
		}
		q.enqueue(req)
		timeSource.Advance(1 * time.Second)
	}

	for i := 0; i < reqCount; i++ {
		req := <-reqCh
		require.Equal(t, fmt.Sprint(i), req.dst)
	}
}

// testBalancerConnHandle is a test connection handle that is used for testing
// the balancer. This currently does not require any methods to be implemented.
type testBalancerConnHandle struct {
	ConnectionHandle
}

var _ ConnectionHandle = &testBalancerConnHandle{}

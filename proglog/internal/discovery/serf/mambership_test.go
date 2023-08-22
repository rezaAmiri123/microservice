package serf_test

import (
	"fmt"
	"github.com/hashicorp/serf/serf"
	"github.com/rezaAmiri123/microservice/proglog/internal/discovery"
	disSerf "github.com/rezaAmiri123/microservice/proglog/internal/discovery/serf"
	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
	"testing"
	"time"
)

func TestMembership(t *testing.T) {
	m, handler := setupMember(t, nil)
	m, _ = setupMember(t, m)
	m, _ = setupMember(t, m)

	require.Eventually(t, func() bool {
		return 2 == len(handler.Joins) &&
			3 == len(m[0].Members()) &&
			0 == len(handler.Leaves)
	}, 3*time.Second, 250*time.Millisecond)

	err := m[2].Leave()
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return 2 == len(handler.Joins) &&
			3 == len(m[0].Members()) &&
			1 == len(handler.Leaves) &&
			serf.StatusLeft == m[0].Members()[2].Status
	}, 3*time.Second, 250*time.Millisecond)

	require.Equal(t, fmt.Sprintf("%d", 2), <-handler.Leaves)
}

func setupMember(t *testing.T, members []*disSerf.MemberShip) (
	[]*disSerf.MemberShip,
	*discovery.FakeHandler,
) {
	id := len(members)
	ports := dynaport.Get(1)
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", ports[0])
	tags := map[string]string{
		"rpc_addr": addr,
	}
	c := disSerf.Config{
		NodeName: fmt.Sprintf("%d", id),
		BindAddr: addr,
		Tags:     tags,
	}

	h := &discovery.FakeHandler{}
	if len(members) == 0 {
		h.Joins = make(chan map[string]string, 3)
		h.Leaves = make(chan string, 1)
	} else {
		c.StartJoinAddrs = []string{
			members[0].BindAddr,
		}
	}
	m, err := disSerf.New(h, c)
	require.NoError(t, err)
	members = append(members, m)
	return members, h
}

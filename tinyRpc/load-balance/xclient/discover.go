package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

type SelectMode int

const (
	RandomSelect     SelectMode = iota // select randomly
	RoundRobinSelect                   // select using Robbin algorithm
)

type Discovery interface {
	Refresh() error                      // refresh from remote registry
	Update(servers []string) error       // update services manually
	Get(mode SelectMode) (string, error) // get a server according to mode
	GetAll() ([]string, error)           // returns all servers in discovery
}

var _ Discovery = (*MultiServicesDiscovery)(nil)

type MultiServicesDiscovery struct {
	r       *rand.Rand   // generate random number
	mu      sync.RWMutex // protect following
	servers []string
	index   int // record the selected position for robin algorithm
}

// NewMultiServerDiscovery created a MultiServerDiscovery instance
func NewMultiServerDiscovery(servers []string) *MultiServicesDiscovery {
	// 初始化时使用时间戳设定随机数种子, 避免每次产生相同的随机数序列
	d := &MultiServicesDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	// index 记录 Round Robin 算法已经轮询到的位置, 为了避免每次从 0 开始, 初始化时随机设定一个值
	d.index = d.r.Intn(math.MaxInt32 - 1)
	return d
}

// Refresh doesn't make sense for MultiServerDiscovery, so ignore it
func (d *MultiServicesDiscovery) Refresh() error {
	return nil
}

// Update the servers of discovery dynamically if needed
func (d *MultiServicesDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

// Get a server according to mode
func (d *MultiServicesDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		s := d.servers[d.index%n] // servers could be updated, so mode n to ensure safety
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

// GetAll returns all servers in discovery
func (d *MultiServicesDiscovery) GetAll() ([]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// return a copy of d.servers
	servers := make([]string, len(d.servers), len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}

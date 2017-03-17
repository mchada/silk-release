// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/vishvananda/netlink"
)

type NetlinkAdapter struct {
	LinkByNameStub        func(string) (netlink.Link, error)
	linkByNameMutex       sync.RWMutex
	linkByNameArgsForCall []struct {
		arg1 string
	}
	linkByNameReturns struct {
		result1 netlink.Link
		result2 error
	}
	linkByNameReturnsOnCall map[int]struct {
		result1 netlink.Link
		result2 error
	}
	ParseAddrStub        func(string) (*netlink.Addr, error)
	parseAddrMutex       sync.RWMutex
	parseAddrArgsForCall []struct {
		arg1 string
	}
	parseAddrReturns struct {
		result1 *netlink.Addr
		result2 error
	}
	parseAddrReturnsOnCall map[int]struct {
		result1 *netlink.Addr
		result2 error
	}
	AddrAddStub        func(netlink.Link, *netlink.Addr) error
	addrAddMutex       sync.RWMutex
	addrAddArgsForCall []struct {
		arg1 netlink.Link
		arg2 *netlink.Addr
	}
	addrAddReturns struct {
		result1 error
	}
	addrAddReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetHardwareAddrStub        func(netlink.Link, net.HardwareAddr) error
	linkSetHardwareAddrMutex       sync.RWMutex
	linkSetHardwareAddrArgsForCall []struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}
	linkSetHardwareAddrReturns struct {
		result1 error
	}
	linkSetHardwareAddrReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetlinkAdapter) LinkByName(arg1 string) (netlink.Link, error) {
	fake.linkByNameMutex.Lock()
	ret, specificReturn := fake.linkByNameReturnsOnCall[len(fake.linkByNameArgsForCall)]
	fake.linkByNameArgsForCall = append(fake.linkByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LinkByName", []interface{}{arg1})
	fake.linkByNameMutex.Unlock()
	if fake.LinkByNameStub != nil {
		return fake.LinkByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.linkByNameReturns.result1, fake.linkByNameReturns.result2
}

func (fake *NetlinkAdapter) LinkByNameCallCount() int {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return len(fake.linkByNameArgsForCall)
}

func (fake *NetlinkAdapter) LinkByNameArgsForCall(i int) string {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return fake.linkByNameArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkByNameReturns(result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	fake.linkByNameReturns = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) LinkByNameReturnsOnCall(i int, result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	if fake.linkByNameReturnsOnCall == nil {
		fake.linkByNameReturnsOnCall = make(map[int]struct {
			result1 netlink.Link
			result2 error
		})
	}
	fake.linkByNameReturnsOnCall[i] = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) ParseAddr(arg1 string) (*netlink.Addr, error) {
	fake.parseAddrMutex.Lock()
	ret, specificReturn := fake.parseAddrReturnsOnCall[len(fake.parseAddrArgsForCall)]
	fake.parseAddrArgsForCall = append(fake.parseAddrArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ParseAddr", []interface{}{arg1})
	fake.parseAddrMutex.Unlock()
	if fake.ParseAddrStub != nil {
		return fake.ParseAddrStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseAddrReturns.result1, fake.parseAddrReturns.result2
}

func (fake *NetlinkAdapter) ParseAddrCallCount() int {
	fake.parseAddrMutex.RLock()
	defer fake.parseAddrMutex.RUnlock()
	return len(fake.parseAddrArgsForCall)
}

func (fake *NetlinkAdapter) ParseAddrArgsForCall(i int) string {
	fake.parseAddrMutex.RLock()
	defer fake.parseAddrMutex.RUnlock()
	return fake.parseAddrArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) ParseAddrReturns(result1 *netlink.Addr, result2 error) {
	fake.ParseAddrStub = nil
	fake.parseAddrReturns = struct {
		result1 *netlink.Addr
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) ParseAddrReturnsOnCall(i int, result1 *netlink.Addr, result2 error) {
	fake.ParseAddrStub = nil
	if fake.parseAddrReturnsOnCall == nil {
		fake.parseAddrReturnsOnCall = make(map[int]struct {
			result1 *netlink.Addr
			result2 error
		})
	}
	fake.parseAddrReturnsOnCall[i] = struct {
		result1 *netlink.Addr
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) AddrAdd(arg1 netlink.Link, arg2 *netlink.Addr) error {
	fake.addrAddMutex.Lock()
	ret, specificReturn := fake.addrAddReturnsOnCall[len(fake.addrAddArgsForCall)]
	fake.addrAddArgsForCall = append(fake.addrAddArgsForCall, struct {
		arg1 netlink.Link
		arg2 *netlink.Addr
	}{arg1, arg2})
	fake.recordInvocation("AddrAdd", []interface{}{arg1, arg2})
	fake.addrAddMutex.Unlock()
	if fake.AddrAddStub != nil {
		return fake.AddrAddStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addrAddReturns.result1
}

func (fake *NetlinkAdapter) AddrAddCallCount() int {
	fake.addrAddMutex.RLock()
	defer fake.addrAddMutex.RUnlock()
	return len(fake.addrAddArgsForCall)
}

func (fake *NetlinkAdapter) AddrAddArgsForCall(i int) (netlink.Link, *netlink.Addr) {
	fake.addrAddMutex.RLock()
	defer fake.addrAddMutex.RUnlock()
	return fake.addrAddArgsForCall[i].arg1, fake.addrAddArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) AddrAddReturns(result1 error) {
	fake.AddrAddStub = nil
	fake.addrAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrAddReturnsOnCall(i int, result1 error) {
	fake.AddrAddStub = nil
	if fake.addrAddReturnsOnCall == nil {
		fake.addrAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addrAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddr(arg1 netlink.Link, arg2 net.HardwareAddr) error {
	fake.linkSetHardwareAddrMutex.Lock()
	ret, specificReturn := fake.linkSetHardwareAddrReturnsOnCall[len(fake.linkSetHardwareAddrArgsForCall)]
	fake.linkSetHardwareAddrArgsForCall = append(fake.linkSetHardwareAddrArgsForCall, struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}{arg1, arg2})
	fake.recordInvocation("LinkSetHardwareAddr", []interface{}{arg1, arg2})
	fake.linkSetHardwareAddrMutex.Unlock()
	if fake.LinkSetHardwareAddrStub != nil {
		return fake.LinkSetHardwareAddrStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetHardwareAddrReturns.result1
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrCallCount() int {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return len(fake.linkSetHardwareAddrArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrArgsForCall(i int) (netlink.Link, net.HardwareAddr) {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return fake.linkSetHardwareAddrArgsForCall[i].arg1, fake.linkSetHardwareAddrArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturns(result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	fake.linkSetHardwareAddrReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturnsOnCall(i int, result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	if fake.linkSetHardwareAddrReturnsOnCall == nil {
		fake.linkSetHardwareAddrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetHardwareAddrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	fake.parseAddrMutex.RLock()
	defer fake.parseAddrMutex.RUnlock()
	fake.addrAddMutex.RLock()
	defer fake.addrAddMutex.RUnlock()
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return fake.invocations
}

func (fake *NetlinkAdapter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

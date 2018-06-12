// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/datianshi/simple-cf-gtm"
)

type FakeHealthCheck struct {
	StartStub          func()
	startMutex         sync.RWMutex
	startArgsForCall   []struct{}
	ReceiveStub        func() []gtm.IP
	receiveMutex       sync.RWMutex
	receiveArgsForCall []struct{}
	receiveReturns     struct {
		result1 []gtm.IP
	}
	receiveReturnsOnCall map[int]struct {
		result1 []gtm.IP
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthCheck) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		fake.StartStub()
	}
}

func (fake *FakeHealthCheck) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeHealthCheck) Receive() []gtm.IP {
	fake.receiveMutex.Lock()
	ret, specificReturn := fake.receiveReturnsOnCall[len(fake.receiveArgsForCall)]
	fake.receiveArgsForCall = append(fake.receiveArgsForCall, struct{}{})
	fake.recordInvocation("Receive", []interface{}{})
	fake.receiveMutex.Unlock()
	if fake.ReceiveStub != nil {
		return fake.ReceiveStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.receiveReturns.result1
}

func (fake *FakeHealthCheck) ReceiveCallCount() int {
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	return len(fake.receiveArgsForCall)
}

func (fake *FakeHealthCheck) ReceiveReturns(result1 []gtm.IP) {
	fake.ReceiveStub = nil
	fake.receiveReturns = struct {
		result1 []gtm.IP
	}{result1}
}

func (fake *FakeHealthCheck) ReceiveReturnsOnCall(i int, result1 []gtm.IP) {
	fake.ReceiveStub = nil
	if fake.receiveReturnsOnCall == nil {
		fake.receiveReturnsOnCall = make(map[int]struct {
			result1 []gtm.IP
		})
	}
	fake.receiveReturnsOnCall[i] = struct {
		result1 []gtm.IP
	}{result1}
}

func (fake *FakeHealthCheck) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHealthCheck) recordInvocation(key string, args []interface{}) {
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

var _ gtm.HealthCheck = new(FakeHealthCheck)

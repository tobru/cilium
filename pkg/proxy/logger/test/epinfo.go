// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build !privileged_tests
// +build !privileged_tests

package test

import (
	"github.com/cilium/cilium/pkg/fqdn/restore"
	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/labels"
	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/proxy/accesslog"
)

type ProxyUpdaterMock struct {
	lock.RWMutex
	Id           uint64
	Ipv4         string
	Ipv6         string
	Labels       []string
	Identity     identity.NumericIdentity
	SidecarProxy bool
}

func (m *ProxyUpdaterMock) GetProxyInfoByFields() (uint64, string, string, []string, string, uint64, error) {
	return m.GetID(), m.GetIPv4Address(), m.GetIPv6Address(), m.GetLabels(), m.GetLabelsSHA(), uint64(m.GetIdentityLocked()), nil
}

func (m *ProxyUpdaterMock) UnconditionalRLock() { m.RWMutex.RLock() }
func (m *ProxyUpdaterMock) RUnlock()            { m.RWMutex.RUnlock() }

func (m *ProxyUpdaterMock) GetID() uint64                                 { return m.Id }
func (m *ProxyUpdaterMock) GetIPv4Address() string                        { return m.Ipv4 }
func (m *ProxyUpdaterMock) GetIPv6Address() string                        { return m.Ipv6 }
func (m *ProxyUpdaterMock) GetLabels() []string                           { return m.Labels }
func (m *ProxyUpdaterMock) GetEgressPolicyEnabledLocked() bool            { return true }
func (m *ProxyUpdaterMock) GetIngressPolicyEnabledLocked() bool           { return true }
func (m *ProxyUpdaterMock) GetIdentityLocked() identity.NumericIdentity   { return m.Identity }
func (m *ProxyUpdaterMock) GetNamedPortLocked(bool, string, uint8) uint16 { return 0 }
func (m *ProxyUpdaterMock) GetLabelsSHA() string {
	return labels.NewLabelsFromModel(m.Labels).SHA256Sum()
}
func (m *ProxyUpdaterMock) HasSidecarProxy() bool       { return m.SidecarProxy }
func (m *ProxyUpdaterMock) ConntrackName() string       { return m.ConntrackNameLocked() }
func (m *ProxyUpdaterMock) ConntrackNameLocked() string { return "global" }

func (m *ProxyUpdaterMock) OnProxyPolicyUpdate(policyRevision uint64) {}
func (m *ProxyUpdaterMock) UpdateProxyStatistics(l4Protocol string, port uint16, ingress, request bool,
	verdict accesslog.FlowVerdict) {
}
func (m *ProxyUpdaterMock) OnDNSPolicyUpdateLocked(rules restore.DNSRules) {}

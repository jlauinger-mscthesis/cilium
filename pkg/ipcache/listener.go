// Copyright 2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipcache

import (
	"net"

	"github.com/cilium/cilium/pkg/identity"
)

// CacheModification represents the type of operation performed upon IPCache.
type CacheModification string

const (
	// Upsert represents Upsertion into IPCache.
	Upsert CacheModification = "Upsert"

	// Delete represents deletion of an entry in IPCache.
	Delete CacheModification = "Delete"
)

// ChangeEvent describes a change to the IP cache. A ChangeEvent is sent to IP cache listeners on
// any modification of the IP cache.
// If an existing CIDR->ID mapping is updated, then
// OldID is not nil; otherwise it is nil.
// HostIP is the IP address of the location of the cidr.
// HostIP is optional and may only be non-nil for an Upsert modification.
// K8sMeta contains the Kubernetes pod namespace and name behind the IP
// and may be nil.

type ChangeEvent struct {
	ModType    CacheModification
	CIDR       net.IPNet
	OldHostIP  net.IP
	NewHostIP  net.IP
	OldID      *identity.NumericIdentity
	NewID      identity.NumericIdentity
	EncryptKey uint8
	K8sMeta    *K8sMetadata
}

// IPIdentityMappingListener represents a component that is interested in
// learning about IP to Identity mapping events.
type IPIdentityMappingListener interface {
	// OnIPIdentityCacheChange will be called whenever there the state of the
	// IPCache has changed.
	OnIPIdentityCacheChange(ce ChangeEvent)

	// OnIPIdentityCacheGC will be called to sync other components which are
	// reliant upon the IPIdentityCache with the IPIdentityCache.
	OnIPIdentityCacheGC()
}

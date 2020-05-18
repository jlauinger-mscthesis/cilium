// Copyright 2017-2020 Authors of Cilium
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

package ciliumio

const (
	// CustomResourceDefinitionGroup is the name of the third party resource group
	CustomResourceDefinitionGroup = "cilium.io"

	// CustomResourceDefinitionVersion is the current version of the resource
	CustomResourceDefinitionVersion = "v2"

	// CustomResourceDefinitionSchemaVersion is semver-conformant version of CRD schema
	// Used to determine if CRD needs to be updated in cluster
	CustomResourceDefinitionSchemaVersion = "1.19"

	// CustomResourceDefinitionSchemaVersionKey is key to label which holds the CRD schema version
	CustomResourceDefinitionSchemaVersionKey = "io.cilium.k8s.crd.schema.version"

	// Cilium Network Policy (CNP)

	// CNPSingularName is the singular name of Cilium Network Policy
	CNPSingularName = "ciliumnetworkpolicy"

	// CNPPluralName is the plural name of Cilium Network Policy
	CNPPluralName = "ciliumnetworkpolicies"

	// CNPKindDefinition is the kind name for Cilium Network Policy
	CNPKindDefinition = "CiliumNetworkPolicy"

	// CNPName is the full name of Cilium Network Policy
	CNPName = CNPPluralName + "." + CustomResourceDefinitionGroup

	// Cilium Cluster wide Network Policy (CCNP)

	// CCNPSingularName is the singular name of Cilium Cluster wide Network Policy
	CCNPSingularName = "ciliumclusterwidenetworkpolicy"

	// CCNPPluralName is the plural name of Cilium Cluster wide Network Policy
	CCNPPluralName = "ciliumclusterwidenetworkpolicies"

	// CCNPKindDefinition is the kind name for Cilium Cluster wide Network Policy
	CCNPKindDefinition = "CiliumClusterwideNetworkPolicy"

	// CCNPName is the full name of Cilium Cluster wide Network Policy
	CCNPName = CCNPPluralName + "." + CustomResourceDefinitionGroup

	// Cilium Endpoint (CEP)

	// CESingularName is the singular name of Cilium Endpoint
	CEPSingularName = "ciliumendpoint"

	// CEPluralName is the plural name of Cilium Endpoint
	CEPPluralName = "ciliumendpoints"

	// CEKindDefinition is the kind name for Cilium Endpoint
	CEPKindDefinition = "CiliumEndpoint"

	// CEPName is the full name of Cilium Endpoint
	CEPName = CEPPluralName + "." + CustomResourceDefinitionGroup

	// Cilium Node (CN)

	// CNSingularName is the singular name of Cilium Node
	CNSingularName = "ciliumnode"

	// CNPluralName is the plural name of Cilium Node
	CNPluralName = "ciliumnodes"

	// CNKindDefinition is the kind name for Cilium Node
	CNKindDefinition = "CiliumNode"

	// CNName is the full name of Cilium Node
	CNName = CNPluralName + "." + CustomResourceDefinitionGroup

	// Cilium Identity

	// CIDSingularName is the singular name of Cilium Identity
	CIDSingularName = "ciliumidentity"

	// CIDPluralName is the plural name of Cilium Identity
	CIDPluralName = "ciliumidentities"

	// CIDKindDefinition is the kind name for Cilium Identity
	CIDKindDefinition = "CiliumIdentity"

	// CIDName is the full name of Cilium Identity
	CIDName = CIDPluralName + "." + CustomResourceDefinitionGroup
)

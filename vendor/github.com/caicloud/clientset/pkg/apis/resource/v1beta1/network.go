package v1beta1

import (
	"net"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network are non-namespaced; the id of the network
type Network struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSpec `json:"spec"`
}

type NetworkSpec struct {
	Subnets          []Subnet `json:"subnets,omitempty"`
	Type             string   `json:"type"`
	CNI              CNI      `json:"cni"`
	IPRecycleTimeout *int32   `json:"ipRecycleTimeout,omitempty"`
	IsFixedIP        bool     `json:"isFixedIP,omitempty"`
}

type CNI struct {
	Calico     *Calico     `json:"calico,omitempty"`
	Canal      *Canal      `json:"canal,omitempty"`
	BridgeVlan *BridgeVlan `json:"bridge-vlan,omitempty"`
}

type Calico struct {
	Mode     string   `json:"mode"`
	BGPMode  string   `json:"bgpMode"`
	RrNodes  []string `json:"rrNodes,omitempty"`
	RrIPs    []string `json:"rrIPs,omitempty"`
	AsNumber int      `json:"asNumber,omitempty"`
}

type Canal struct {
}

type BridgeVlan struct {
	PhyDev string `json:"phyDev"`
	VlanID uint16 `json:"vlanID"`
}

type Subnet struct {
	ID         string `json:"id"`
	CIDR       string `json:"cidr"`
	Gateway    net.IP `json:"gateway,omitempty"`
	RangeStart net.IP `json:"rangeStart,omitempty"`
	RangeEnd   net.IP `json:"rangeEnd,omitempty"`
	// NodeCidrMaskSize define canal network node cidr mask size
	NodeCidrMaskSize int `json:"nodeCidrMaskSize,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList is a collection of networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of Networks
	Items []Network `json:"items"`
}

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualService struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.
	// +optional
	Spec VirtualServiceSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualServiceSpec struct {
	Gateways []Gateway   `protobuf:"bytes,2,rep,name=gateways" json:"gateways,omitempty"`
	Hosts    []Host      `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
	HTTP     []HTTPRoute `protobuf:"bytes,3,rep,name=http" json:"http,omitempty"`
}

type Gateway string
type Host string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HTTPRoute struct {
	Match []MatchRequest      `protobuf:"bytes,1,rep,name=match" json:"match,omitempty"`
	Route []DestinationWeight `protobuf:"bytes,2,rep,name=route" json:"route,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MatchRequest struct {
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DestinationWeight struct {
	Destination Destination `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Destination struct {
	Host string       `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port PortSelector `protobuf:"bytes,3,opt,name=port" json:"port,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PortSelector struct {
	Port PortSelector_Name `protobuf_oneof:"port"`
}

type isPortSelector_Port interface {
	isPortSelector_Port()
	MarshalTo([]byte) (int, error)
	Size() int
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PortSelector_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*PortSelector_Name) isPortSelector_Port() {}

func (m *PortSelector_Name) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovVirtualService(uint64(l))
	return n
}

func sovVirtualService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func (m *PortSelector_Name) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintVirtualService(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}
func encodeVarintVirtualService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

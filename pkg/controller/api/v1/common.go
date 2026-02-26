package v1

type CIDR string

// +kubebuilder:validation:Enum=unspecified;creating;updating;active;deleting;deleted;failed
type ResourceState string

const (
	UNSPECIFIED ResourceState = "unspecified"
	ACTIVE      ResourceState = "active"
	DELETED     ResourceState = "deleted"
	FAILED      ResourceState = "failed"
)

type ResourceStatus struct {
	State   ResourceState `json:"state,omitempty"`
	Message string        `json:"message,omitempty"`
}

const (
	LabelKeyRegionID = "cloudovn.io/regionID"
	LabelKeyAZID     = "cloudovn.io/zoneID"

	LabelKeyHostNetworkEnable = "cloudovn.io/hostNetworkEnable"

	LabelKeyLogicRouterName = "cloudovn.io/logicRouterName"
)

const (
	AnnotationKeyRegionName = "cloudovn.io.regionName"
	AnnotationKeyAZName     = "cloudovn.io.zoneName"
)

package client

const (
	KubeControllerServiceType                       = "kubeControllerService"
	KubeControllerServiceFieldClusterCIDR           = "clusterCidr"
	KubeControllerServiceFieldExtraArgs             = "extraArgs"
	KubeControllerServiceFieldExtraBinds            = "extraBinds"
	KubeControllerServiceFieldImage                 = "image"
	KubeControllerServiceFieldServiceClusterIPRange = "serviceClusterIpRange"
)

type KubeControllerService struct {
	ClusterCIDR           string            `json:"clusterCidr,omitempty" yaml:"clusterCidr,omitempty"`
	ExtraArgs             map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds            []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	Image                 string            `json:"image,omitempty" yaml:"image,omitempty"`
	ServiceClusterIPRange string            `json:"serviceClusterIpRange,omitempty" yaml:"serviceClusterIpRange,omitempty"`
}

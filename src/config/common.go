package config

var Mirrors = []string{}

var TarFiles = []string{
	"ansible.tar",
	"centos7_9_2009.tar",
	"centos8_3_2011.tar",
	"kubespray.tar",
	"kubespray_cache.tar.gz",
}

const (
	Version    = "0.3-2021.07.26"
	El7DirName = "centos7_9_2009"
	El8DirName = "centos8_3_2011"

	KubesprayImage    = "quay.io_kubespray_kubespray_v2.15.1.tar"
	MongoTypeP        = "primary"
	MongoTypeS        = "secondary"
	MongoTypeA        = "arbiter"
	DockerHubRegistry = "registry-1.docker.io"
	ChartVersionKey   = "chart_version"
	InventoryPath     = "inventory/hosts.yaml"
)

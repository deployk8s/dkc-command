package inventory

import (
	"io/ioutil"
	"sort"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

type Inventory struct {
	All            All    `yaml:"all"`
	RemoteUsername string `yaml:"-"`
	RemotePassword string `yaml:"-"`
	RemoteSSHkey   string `yaml:"-"`
	Hosts          Hosts  `yaml:"-"`
}
type Hosts []Host

type ChildrenStruct struct {
	KubeMaster ChildrenBaseStruct `yaml:"kube_control_plane"`
	Om         ChildrenBaseStruct `yaml:"om"`
	Sg1        ChildrenBaseStruct `yaml:"sg1"`
	Rs1        ChildrenBaseStruct `yaml:"rs1"`
	Sg         ChildrenBaseStruct `yaml:"sg"`
	Rs2        ChildrenBaseStruct `yaml:"rs2"`
	Db         ChildrenBaseStruct `yaml:"db"`
	KubeNode   ChildrenBaseStruct `yaml:"kube_node"`
	K8sCluster ChildrenBaseStruct `yaml:"k8s_cluster"`
}

type Variable map[string]interface{}

type ChildrenBaseStruct struct {
	Children map[string]struct{}    `yaml:"children,omitempty"`
	Vars     map[string]interface{} `yaml:"vars,omitempty"`
	Hosts    map[string]Variable    `yaml:"hosts,omitempty"`
}

type HostInfo struct {
	AnsibleConnection     string            `yaml:"ansible_connection,omitempty"`
	RemoteMachineUsername string            `yaml:"remote_machine_username,omitempty"`
	RemoteMachinePassword string            `yaml:"remote_machine_password,omitempty"`
	RemoteSSHkey          string            `yaml:"remote_sshkey,omitempty"`
	AnsibleHost           string            `yaml:"ansible_host,omitempty"`
	AnsiblePort           int               `yaml:"ansible_port,omitempty"`
	NodeLabels            map[string]string `yaml:"node_labels,omitempty"`
}
type All struct {
	Variables Vars                `yaml:"vars"`
	Hosts     map[string]HostInfo `yaml:"hosts"`
	Children  map[string]ChildrenBaseStruct
}

type Vars struct {
	ExternalDomainName      string   `yaml:"external_domain_name"`
	LoginType               string   `yaml:"login_type"`
	MongodbType             string   `yaml:"mongodb_type"`
	MongodbAliyunPassword   string   `yaml:"mongodb_aliyun_password"`
	MongodbAliyunServers    []string `yaml:"mongodb_aliyun_servers"`
	MongodbAliyunReplicaset string   `yaml:"mongodb_aliyun_replicaset"`
}

type HostInstance struct {
	Client *ssh.Client
	Host   Host
}
type Host struct {
	Hostname     string `json:"hostname"`
	Ip           string `json:"ip"`
	Port         int    `json:"-"`
	IsLocal      bool   `json:"-"`
	Username     string `json:"username"`
	LoginType    string
	SSHkey       string `json:"sshkey"`
	Password     string `json:"password"`
	IsMongo      string `json:"mongo"`
	IsMaster     bool   `json:"master"`
	IsNode       bool   `json:"node"`
	IsOm         bool   `json:"om"`
	IsSg         bool   `json:"sg"`
	IsLogging    bool   `json:"is_logging"`
	IsMonitoring bool
	IsTracking   bool
	Arch         string `json:"arch"`
}

func NewHostInstance(h Host) *HostInstance {
	var hi HostInstance
	hi.Host = h
	return &hi
}

func NewInventory() *Inventory {
	d, _ := ioutil.ReadFile(config.Kconfig.InventoryFile)
	var t Inventory
	if err := yaml.Unmarshal(d, &t); err != nil {
		log.Log.Fatal(err.Error())
	}
	for k, v := range t.All.Hosts {
		if k == "localhost" {
			t.RemotePassword = v.RemoteMachinePassword
			t.RemoteUsername = v.RemoteMachineUsername
			t.RemoteSSHkey = v.RemoteSSHkey
			break
		}
	}
	for k, v := range t.All.Hosts {
		if k != "localhost" {
			h := Host{Hostname: k,
				Ip:           v.AnsibleHost,
				Port:         v.AnsiblePort,
				IsLocal:      false,
				Username:     t.RemoteUsername,
				Password:     t.RemotePassword,
				SSHkey:       t.RemoteSSHkey,
				IsMongo:      t.isMongo(k),
				IsMaster:     t.isMaster(k),
				IsNode:       t.isNode(k),
				IsLogging:    t.isLoging(k),
				IsMonitoring: t.isMonitoring(k),
				IsTracking:   t.isTracking(k),
				IsSg:         t.isSg(k),
				Arch:         "connect err.",
			}
			t.Hosts = append(t.Hosts, h)
		}

	}
	sort.Sort(t.Hosts)
	return &t
}

func (i *Inventory) isMaster(hostname string) bool {
	for _, v := range i.GetMaster() {
		if hostname == v {
			return true
		}
	}
	return false
}

func (i *Inventory) isNode(hostname string) bool {
	for _, v := range i.GetNode() {
		if hostname == v {
			return true
		}
	}
	return false
}

// return arbiter/secondary/primary
func (i *Inventory) isMongo(hostname string) string {
	for _, v := range i.GetMongo() {
		if hostname == v {
			return i.All.Children["rs1"].Hosts[hostname]["rs1_role"].(string)
		}
	}
	return ""
}

func (i *Inventory) isOm(hostname string) bool {
	for _, v := range i.GetHostnamesByGroup("om") {
		if hostname == v {
			return true
		}
	}
	return false
}

func (i *Inventory) isLoging(hostname string) bool {
	if v, ok := i.All.Hosts[hostname]; ok {
		if _, exist := v.NodeLabels["nodeType_logging"]; exist {
			return true
		}
	}
	return false
}
func (i *Inventory) isMonitoring(hostname string) bool {
	if v, ok := i.All.Hosts[hostname]; ok {
		if _, exist := v.NodeLabels["nodeType_monitoring"]; exist {
			return true
		}
	}
	return false
}
func (i *Inventory) isTracking(hostname string) bool {
	if v, ok := i.All.Hosts[hostname]; ok {
		if _, exist := v.NodeLabels["nodeType_tracking"]; exist {
			return true
		}
	}
	return false
}

func (i *Inventory) isSg(hostname string) bool {
	for _, v := range i.GetHostnamesByGroup("sg") {
		if hostname == v {
			return true
		}
	}
	return false
}

func (i *Inventory) GetHostnamesByGroup(groupName string) []string {
	var h []string
	hk := make(map[string]struct{})
	for k, v := range i.All.Children {
		if k == groupName {
			for kh, _ := range v.Hosts {
				if _, ok := hk[kh]; ok {
					continue
				}
				hk[kh] = struct{}{}
				h = append(h, kh)
			}
			for kc, _ := range v.Children {
				hc := i.GetHostnamesByGroup(kc)
				for _, hci := range hc {
					if _, ok := hk[hci]; ok {
						continue
					}
					hk[hci] = struct{}{}
					h = append(h, hci)
				}
			}
		}
	}
	return h
}

func (i *Inventory) GetNode() []string {
	return i.GetHostnamesByGroup("kube_node")
}

func (i *Inventory) GetMaster() []string {
	return i.GetHostnamesByGroup("kube_control_plane")
}

func (i *Inventory) GetMongo() []string {
	return i.GetHostnamesByGroup("rs1")
}

func (l Hosts) Len() int {
	return len(l)
}
func (list Hosts) Less(i, j int) bool {
	return list[i].Hostname < list[j].Hostname
}
func (list Hosts) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

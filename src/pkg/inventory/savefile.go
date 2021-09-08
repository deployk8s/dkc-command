package inventory

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Remote struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InventoryData struct {
	Ingress string `json:"ingress"`
	Remote  Remote `json:"remote"`
	Hosts   []Host `json:"hosts"`
}

func SaveFile(id InventoryData) error {
	//d, _ := ioutil.ReadFile("inventory/base")
	var t Inventory
	//if err := yaml.Unmarshal(d, &t); err != nil {
	//	return err
	//}
	t.All.Variables.ExternalDomainName = id.Ingress
	t.All.Hosts = make(map[string]HostInfo)
	t.All.Hosts["localhost"] = HostInfo{RemoteMachineUsername: id.Remote.Username, RemoteMachinePassword: id.Remote.Password}
	t.All.Children = make(map[string]ChildrenBaseStruct)
	for _, v := range id.Hosts {
		t.All.Hosts[v.Hostname] = HostInfo{
			AnsibleHost: v.Ip,
			AnsiblePort: 22,
		}
		if v.IsMaster {
			m := t.All.Children["kube-master"]
			if m.Hosts == nil {
				t.All.Children["kube-master"] = ChildrenBaseStruct{
					Vars:     Variable{},
					Hosts:    make(map[string]Variable),
					Children: make(map[string]struct{}),
				}
			}
			t.All.Children["kube-master"].Hosts[v.Hostname] = Variable{}
		}
		if v.IsSg {
			m := t.All.Children["sg1"]

			if m.Hosts == nil {
				t.All.Children["sg1"] = ChildrenBaseStruct{
					Vars:     Variable{},
					Hosts:    make(map[string]Variable),
					Children: make(map[string]struct{}),
				}
			}
			t.All.Children["sg1"].Hosts[v.Hostname] = Variable{}
		}
		if v.IsOm {
			m := t.All.Children["om"]
			if m.Hosts == nil {
				t.All.Children["om"] = ChildrenBaseStruct{
					Vars:     Variable{},
					Hosts:    make(map[string]Variable),
					Children: make(map[string]struct{}),
				}

			}
			t.All.Children["om"].Hosts[v.Hostname] = Variable{}
		}
		if v.IsMongo != "" {
			m := t.All.Children["rs1"]
			if m.Hosts == nil {
				t.All.Children["rs1"] = ChildrenBaseStruct{
					Vars:     Variable{},
					Hosts:    make(map[string]Variable),
					Children: make(map[string]struct{}),
				}
			}
			t.All.Children["rs1"].Hosts[v.Hostname] = map[string]interface{}{
				"rs1_role": v.IsMongo,
			}
		}
		if !v.IsMaster {
			m := t.All.Children["kube-node"]
			if m.Hosts == nil {
				//t.All.Children["kube-node"].Hosts = make(map[string]Variable)
				t.All.Children["kube-node"] = ChildrenBaseStruct{
					Vars:  Variable{},
					Hosts: make(map[string]Variable),
					Children: map[string]struct{}{
						"rs": {},
						"db": {},
						"om": {},
					},
				}
			}
			t.All.Children["kube-node"].Hosts[v.Hostname] = Variable{}
		}
	}
	t.All.Children["k8s-cluster"] = ChildrenBaseStruct{
		Children: map[string]struct{}{
			"kube-master": {},
			"kube-node":   {},
		},
	}
	t.All.Children["db"] = ChildrenBaseStruct{
		Children: map[string]struct{}{
			"rs1": {},
			"rs2": {},
		},
	}
	t.All.Children["sg"] = ChildrenBaseStruct{
		Children: map[string]struct{}{
			"sg1": {},
		},
	}

	//fmt.Printf("########## %+v", t)
	out, _ := yaml.Marshal(t)
	//out,_ :=json.Marshal(t)
	newfile, _ := os.Create("inventory/hosts.yaml")
	defer newfile.Close()
	io.Copy(newfile, bytes.NewReader(out))
	//fmt.Printf("%+v", t)

	return nil
}

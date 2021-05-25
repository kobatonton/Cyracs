package sub

import (
	"cyracs/data"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func CreateMetadata(Vm data.Vm) {
	Dir, _ := os.Getwd()
	funcMap := template.FuncMap{
		"inc": func(i int) int { return i + 1 },
	}
	tpl, err := template.New("metadata.tpl").Funcs(funcMap).ParseFiles("template/metadata.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	if len(Vm) != 0 {
		for j, vm := range Vm {
			vmdir := Dir + "/initfile/" + vm.Hostname
			if _, err := os.Stat(vmdir); os.IsNotExist(err) {
				os.Mkdir(vmdir, 0777)
			}
			cf, err := os.Create(Dir + "/initfile/" + vm.Hostname + "/meta-data")
			if err != nil {
				log.Fatalln("error creating vm xml file", err)
			}
			defer cf.Close()
			err = tpl.Execute(cf, Vm[j])
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func CreateUserdata(Vm data.Vm) {
	var tplfile string
	Dir, _ := os.Getwd()
	if len(Vm) != 0 {
		for j, vm := range Vm {
			funcMap := template.FuncMap{
				"mac": func(i int) string {
					net := vm.Interfaces[i].Network
					cmd := "virsh dumpxml " + vm.Hostname + " |grep -B1 vbr_" + net + " |grep mac | grep -i -o '[0-9A-F]\\{2\\}\\(:[0-9A-F]\\{2\\}\\)\\{5\\}'"
					out, _ := exec.Command("sh", "-c", cmd).Output()
					return string(out)
				},
				"defmac": func() string {
					cmd := "virsh dumpxml " + vm.Hostname + " |grep -B1 virbr0 |grep mac | grep -i -o '[0-9A-F]\\{2\\}\\(:[0-9A-F]\\{2\\}\\)\\{5\\}'"
					out, _ := exec.Command("sh", "-c", cmd).Output()
					return string(out)
				},
			}
			if vm.Nettype == "network-scripts" {
				tplfile = "userdata-ns.tpl"
			} else if vm.Nettype == "netplan" {
				tplfile = "userdata-np.tpl"
			}
			tpl, err := template.New(tplfile).Funcs(funcMap).ParseFiles("template/" + tplfile)
			if err != nil {
				log.Fatalln(err)
			}
			cf, err := os.Create(Dir + "/initfile/" + vm.Hostname + "/user-data")
			if err != nil {
				log.Fatalln("error creating vm xml file", err)
			}
			defer cf.Close()
			err = tpl.Execute(cf, Vm[j])
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func Genisoimage(Vm data.Vm) {
	if len(Vm) != 0 {
		for _, vm := range Vm {
			//Create cloud-init files
			cmd := "genisoimage -o initfile/" + vm.Hostname + "/" + vm.Hostname + ".iso -V cidata -r -J initfile/" + vm.Hostname + "/user-data initfile/" + vm.Hostname + "/meta-data"
			err := exec.Command("sh", "-c", cmd).Run()
			if err != nil {
				log.Fatalln("error creating file", err)
			}
			cmd = "virsh attach-device --config " + vm.Hostname + " xml/" + vm.Hostname + "-iso.xml"
			err = exec.Command("sh", "-c", cmd).Run()
			if err != nil {
				log.Fatalln("error attach:", err)
			}
		}
	}
}

package sub

import (
	"cyracs/data"
	"log"
	"os/exec"
)

func Nicadd(Vm data.Vm) {
	for _, vm := range Vm {
		cmd := "virsh attach-interface --source virbr0 --type bridge --model virtio --persistent " + vm.Hostname
		err := exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Fatalln(err)
		}
		for _, iface := range vm.Interfaces {
			cmd = "virsh attach-interface --source vbr_" + iface.Network + " --type bridge --model virtio --persistent " + vm.Hostname
			err = exec.Command("sh", "-c", cmd).Run()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

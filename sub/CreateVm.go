package sub

import (
	"cyracs/data"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func CreateVm(Vm data.Vm) {
	fmt.Printf("####Create VM#####\n")
	for _, vm := range Vm {
		//Create XML
		cmd := "qemu-img create -f qcow2 -F qcow2 -b " + vm.Baseimg + " images/" + vm.Hostname + ".qcow2 " + strconv.Itoa(vm.Disk) + "G"
		err := exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Fatalln("error creating file", err)
		}
		//define VM
		cmd = "virsh define xml/" + vm.Hostname + ".xml"
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Fatalln("error define:", err)
		}
		fmt.Printf("create vm success: %+v\n", vm.Hostname)
	}
}

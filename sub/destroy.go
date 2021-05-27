package sub

import (
	"cyracs/data"
	"fmt"
	"log"
	"os/exec"
)

func Destroy(Vm data.Vm) {
	fmt.Printf("####Destroy VM#####\n")
	for _, vm := range Vm {
		cmd := "virsh destroy " + vm.Hostname
		err := exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error destroy:", err)
		}
		cmd = "virsh undefine " + vm.Hostname
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error undefine:", err)
		}
		fmt.Printf("destroyed vm: %+v\n", vm.Hostname)
	}
}

func DestroyVnet(Vnet data.Vnet) {
	fmt.Printf("####Destroy Vnet#####\n")
	for _, vnet := range Vnet {
		cmd := "virsh net-destroy " + vnet.Name
		err := exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error destroy:", err)
		}
		cmd = "virsh net-undefine " + vnet.Name
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error undefine:", err)
		}
		fmt.Printf("destroyed vnet: %+v\n", vnet.Name)
	}
}

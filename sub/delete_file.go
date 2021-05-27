package sub

import (
	"cyracs/data"
	"log"
	"os"
)

//delete some files
func Deletefile(Vm data.Vm, Vnet data.Vnet) {
	for _, vm := range Vm {
		fname := "xml/" + vm.Hostname + ".xml"
		if err := os.Remove(fname); err != nil {
			log.Fatalln("error destroy:", err)
		}
		fname = "initfile/" + vm.Hostname
		if err := os.RemoveAll(fname); err != nil {
			log.Fatalln("error destroy:", err)
		}
		fname = "images/" + vm.Hostname + ".qcow2"
		if err := os.Remove(fname); err != nil {
			log.Fatalln("error destroy:", err)
		}
	}
	for _, vnet := range Vnet {
		fname := "xml/" + vnet.Name + ".xml"
		if err := os.Remove(fname); err != nil {
			log.Fatalln("error destroy:", err)
		}
	}
}

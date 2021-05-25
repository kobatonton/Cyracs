package sub

import (
	"cyracs/data"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func CreateVnetXml(Vnet data.Vnet) {
	fmt.Printf("####Create Vnet#####\n")
	Dir, _ := os.Getwd()
	tpl, err := template.New("vnettpl.xml").ParseFiles("template/vnettpl.xml")
	if err != nil {
		log.Println(err)
	}
	if len(Vnet) != 0 {
		//createfile
		for j, vnet := range Vnet {
			cf, err := os.Create(Dir + "/xml/" + vnet.Name + ".xml")
			if err != nil {
				log.Println("error creating vnet xml", err)
			}
			defer cf.Close()
			err = tpl.Execute(cf, Vnet[j])
			if err != nil {
				log.Println(err)
			}
		}
	}
	//define
	for _, vnet := range Vnet {
		cmd := "virsh net-define xml/" + vnet.Name + ".xml"
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error net-define:", err)
		}
	}
	//autstart and start
	for _, vnet := range Vnet {
		cmd := "virsh net-autostart " + vnet.Name
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error net-autostart:", err)
		}
		cmd = "virsh net-start " + vnet.Name
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			log.Println("error net-start:", err)
		}
		fmt.Printf("create vnet success: %+v\n", vnet.Name)
	}
}

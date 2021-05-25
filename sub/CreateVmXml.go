package sub

import (
	"cyracs/data"
	"log"
	"os"
	"text/template"
)

func CreateVmXml(Vm data.Vm) {
	Dir, _ := os.Getwd()
	funcMap := template.FuncMap{
		"mem": func(a int) int { return a * 1024 * 1024 },
		"dir": func() string { dir, _ := os.Getwd(); return dir },
	}
	tpl, err := template.New("vmtpl.xml").Funcs(funcMap).ParseFiles("template/vmtpl.xml")
	if err != nil {
		log.Fatalln(err)
	}
	if len(Vm) != 0 {
		//create VM xml file
		for j, vm := range Vm {
			cf, err := os.Create(Dir + "/xml/" + vm.Hostname + ".xml")
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

func CreateIsoXml(Vm data.Vm) {
	Dir, _ := os.Getwd()
	funcMap := template.FuncMap{
		"dir": func() string { dir, _ := os.Getwd(); return dir },
	}
	tpl, err := template.New("isotpl.xml").Funcs(funcMap).ParseFiles("template/isotpl.xml")
	if err != nil {
		log.Fatalln(err)
	}
	if len(Vm) != 0 {
		for j, vm := range Vm {
			//create disk attaching xml file
			cf, err := os.Create(Dir + "/xml/" + vm.Hostname + "-iso.xml")
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

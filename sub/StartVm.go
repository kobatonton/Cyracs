package sub

import (
	"cyracs/data"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

func StartVm(Vm data.Vm) {
	fmt.Printf("####Waiting start vms####\n")
	var p float64 = data.Core * data.Threads * 0.7
	parallel := int(p) //parallel execute
	sig := make(chan string, parallel)
	defer close(sig)
	Worker(sig, Vm)
}

func Worker(sig chan string, Vm data.Vm) {
	var wg sync.WaitGroup
	for _, vm := range Vm {
		sig <- fmt.Sprintf("sig %s", vm.Hostname)
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			defer func() {
				<-sig
			}()
			state(host)
		}(vm.Hostname)
	}
	wg.Wait()
}

func state(hostname string) {
	//Start vm
	fmt.Printf("waiting start vm: %+v\n", hostname)
	cmd := "virsh start " + hostname
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		log.Fatalln("error start vm:", hostname, err)
	}
	cmd = "virsh domstate " + hostname + "|head -n 1 |sed 's/ //g'"
	timer := time.NewTimer(time.Duration(data.Timeout) * time.Second) //Timeout
	defer timer.Stop()
L:
	for range time.Tick(5 * time.Second) {
		select {
		case <-timer.C:
			fmt.Printf("Timeout start vm: %+v \n", hostname)
			runtime.Goexit()
		default:
			out, _ := exec.Command("sh", "-c", cmd).Output()
			if strings.Contains(string(out), "shutoff") == true {
				fmt.Printf("start %+v successfully\n", hostname)
				break L
			}
		}
	}
	//detach iso
	cmd = "virsh detach-disk " + hostname + " sda --config"
	err = exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		log.Fatalln("error detach-disk:", hostname, err)
	}
	//Restart vm
	cmd = "virsh start " + hostname
	err = exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		log.Fatalln("error start vm:", hostname, err)
	}
}

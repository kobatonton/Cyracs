package main

import (
	"cyracs/data"
	"cyracs/sub"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if 1 <= len(args) && len(args) <= 2 {
		if args[0] == "create" { //create option
			file := args[1]
			sub.Check(file)
			result := sub.Road(file)
			sub.CreateVmXml(result.Vm)
			sub.CreateVnetXml(result.Vnet)
			sub.CreateIsoXml(result.Vm)
			sub.CreateVm(result.Vm)
			sub.Nicadd(result.Vm)
			sub.CreateMetadata(result.Vm)
			sub.CreateUserdata(result.Vm)
			sub.Genisoimage(result.Vm)
			sub.StartVm(result.Vm)
		} else if args[0] == "destroy" { //cdestroy option
			file := args[1]
			sub.Check(file)
			result := sub.Road(file)
			sub.Destroy(result.Vm)
			sub.DestroyVnet(result.Vnet)
			sub.Deletefile(result.Vm, result.Vnet)
		} else if args[0] == "help" { //cdestroy option
			fmt.Println(data.Help)
		} else {
			fmt.Printf("Invalid option '%s' is specified.\n", args[0])
		}
	} else {
		fmt.Printf("Try 'go run main.go help' for more information.\n")
		os.Exit(0)
	}
}

//This script is Definition of struct and strings.
package data

var rangeinf Range

var Timeout = 600 //Timeout in Create VM [Seconds].Can be changed.

var Help = `Usage:

main.go <options> [include_file]

The options are:

	create	crate range
	delete	derete range

Include_file:
choose a using configuration file.`

type Range struct {
	Id   int  `json:"id"`
	Vnet Vnet `json:"network"`
	Vm   Vm   `json:"vm"`
}

type Vnet []struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Vm []struct {
	Uuid       string       `json:"uuid"`
	Hostname   string       `json:"hostname"`
	Baseimg    string       `json:"baseimage"`
	Cpu        int          `json:"cpu"`
	Memory     int          `json:"memory"`
	Disk       int          `json:"disk"`
	Interfaces []Interfaces `json:"interfaces"`
	Nettype    string       `json:"nettype"`
	Packages   []string     `json:"packages"`
}

type Interfaces struct {
	Network string `json:"network"`
	Ip      string `json:"ip"`
	Prefix  string `json:"prefix"`
	Gateway string `json:"gw"`
	Dns     string `json:"dns"`
}

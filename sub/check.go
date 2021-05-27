package sub

import (
	"log"
	"os"
)

func Check(file string) {
	cdir, _ := os.Getwd()
	slice := []string{"config", "xml", "initfile", "images"}
	for _, dir := range slice {
		Dir := cdir + "/" + dir
		if _, err := os.Stat(Dir); os.IsNotExist(err) {
			os.Mkdir(Dir, 0777)
		}
	}
	fl := cdir + "/config/" + file
	if _, err := os.Stat(fl); os.IsNotExist(err) {
		log.Fatal(err)
	}
}

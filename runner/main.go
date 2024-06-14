package main

import (
	"fmt"
	"gopyload"
)

func main() {
	user := "pyload"
	pass := "pyload"
	p, err := gopyload.NewPyloadAndLogin("http://pyload.lxc.1jz.cc:8000", user, pass)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	p.Kill()
	b, err := p.StatusDownloads()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("b: %v\n", b)
	// fmt.Printf("f: %v\n", f)
	// b, err := p.GetPluginConfig()
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	return
	// }
	// fmt.Printf("b: %v\n", b)
	gopyload.SaveAnyToJson(b, "tmp\\new", "StatusDownloads")

}

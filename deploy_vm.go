package main

import (
	"fmt"

	"github.com/alexander-uljev/kuber-go/args"
	"github.com/alexander-uljev/kuber-go/virtm"
)

func main() {
	vmName, imagePath := args.Parse() // "go-vm-1" "ubuntu24"
	conn := virtm.KVMConnect()
	def := virtm.GenerateDefinition(vmName, imagePath)
	domain := virtm.DefineVM(conn, def)
	virtm.CreateVM(domain)
	fmt.Println("All good! The %name VM spawned", vmName)
}

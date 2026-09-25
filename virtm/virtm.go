package virtm

import (
	"log"
	"os"

	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

const KVM_DRIVER = "qemu:///system"
const IMAGES_PATH = "/var/lib/libvirt/images/"

func KVMConnect() *libvirt.Connect {
	conn, err := libvirt.NewConnect(KVM_DRIVER)
	if err != nil {
		log.Fatalf("Failed to connect to default hypervisor. Is it running?")
	}
	return conn
}

func GenerateDefinition(name, imagePath string) string {
	checkPath(imagePath)
	domainCfg := &libvirtxml.Domain{
		Type: "kvm",
		Name: name,
		Memory: &libvirtxml.DomainMemory{
			Value: 2048,
			Unit:  "MiB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: 2,
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch:    "x86_64",
				Machine: "pc",
				Type:    "hvm",
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{Name: "qemu", Type: "qcow2"},
					Source: &libvirtxml.DomainDiskSource{File: &libvirtxml.DomainDiskSourceFile{File: imagePath}},
					Target: &libvirtxml.DomainDiskTarget{Dev: "vda", Bus: "virtio"},
				},
			},
			Serials: []libvirtxml.DomainSerial{
				{
					Target: &libvirtxml.DomainSerialTarget{Port: new(uint)}, // Порт 0
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				{
					Target: &libvirtxml.DomainConsoleTarget{Type: "serial", Port: new(uint)},
				},
			},
		},
	}
	domainXML, err := domainCfg.Marshal()
	if err != nil {
		log.Fatal("Failed to generate vm's xml definition")
	}
	return domainXML
}

func DefineVM(conn *libvirt.Connect, definition string) *libvirt.Domain {
	domain, err := conn.DomainDefineXML(definition)
	if err != nil {
		log.Fatal("Failed to define the VM in libvirt\n", err)
	}
	defer domain.Free()
	return domain
}

func CreateVM(domain *libvirt.Domain) {
	err := domain.Create()
	if err != nil {
		log.Fatal("Failed to spawn a VM\n", err)
	}
}

func checkPath(path string) error {
	_, err := os.Open(path)
	switch err {
	case os.ErrNotExist:
		log.Fatal("Image path does not exist")
	case os.ErrPermission:
		log.Fatal("No permissions to access image path")
	}
	return nil
}

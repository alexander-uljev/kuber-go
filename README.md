# Расширения для kubernetes на go
На текущий момент пакет умеет создавать виртуальную машину посредством libvirt и KVM/QEMU. Все необходимые функции вынесены в отдельный пакет, вам лишь только нужно вызывать их со своими аргументами и своим образом диска. Пример кода для запуска ВМ:
```go
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
```

## Установка
1. Скачайте пакет `git clone https://github.com/alexander-uljev/kuber-go`
2. Перейдите в скачанную папку и установите пакет `go build` или `go install`

## Образы
Указывайте полный путь до вашего образа диска и следите за его правами

## Развитие
1. Научить пакет делать тонкую копию образа и перемещать его в дефолтную локацию
2. Доработать модель запускаемой ВМ
3. Добавить возможности редактировать, удалять ВМ и запуск пачками
4. Добавить управление именами ВМ: генерация последовательных, говорящих имён
5. Добавить управление доменными именами: создавать сетевое имя ВМ через DNS-запись
6. Добавить поддержку cloud-init

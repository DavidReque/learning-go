package interfaces

import "fmt"

type Mobile struct {
	Brand string
}

func (m Mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d", m, m.Brand, power)
}

type Laptop struct {
	CPU string
}

func (l Laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d", l, l.CPU, power)
}

type Toaster struct {
	Amount int
}

func (t Toaster) Draw(power int) {
	fmt.Printf("%T -> amount: %d, power: %d", t, t.Amount, power)
}

type Kettle struct {
	Quantity string
}

func (k Kettle) Draw(power int) {
	fmt.Printf("%T -> quantity: %s, power: %d", k, k.Quantity, power)
}

type PowerDrawer interface {
	Draw(power int)
}

type Socket struct{}


func (Socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}
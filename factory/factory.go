package factory

import "github.com/bimbimprasetyoafif/simple-go-design-pattern/keyboard"

type Vendor string

const (
	Ducky Vendor = "ducky"
	Razor Vendor = "razor"
)

func NewKeyboardFactory(vendor Vendor) keyboard.Keyboard {
	switch vendor {
	case Ducky:
		return keyboard.NewDucky()
	case Razor:
		return keyboard.NewRazor()
	}

	return nil
}

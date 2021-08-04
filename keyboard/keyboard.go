package keyboard

type Keyboard interface {
	Type() string
	OnBacklit() string
}

type Ducky struct {
}

func (d Ducky) Type() string {
	return "hello from ducky"
}
func (d Ducky) OnBacklit() string {
	return "ducky LIT!!"
}

func NewDucky() Keyboard {
	return &Ducky{}
}

type Razor struct {
}

func (r Razor) Type() string {
	return "hello world from razor"
}
func (r Razor) OnBacklit() string {
	return "razor going LIT!!"
}

func NewRazor() Keyboard {
	return &Razor{}
}

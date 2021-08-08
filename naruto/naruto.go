package naruto

type KagebunshinJutsu interface {
	Clone() KagebunshinJutsu
	GetCloneID() int
}

type naruto struct {
	CloneNumber int
}

func (n naruto) GetCloneID() int {
	return n.CloneNumber
}

func (n naruto) Clone() KagebunshinJutsu {
	return &naruto{
		CloneNumber: n.CloneNumber + 1,
	}
}

func CreateNaruto() KagebunshinJutsu {
	return &naruto{}
}

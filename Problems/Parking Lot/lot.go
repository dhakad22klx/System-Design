package main

type LotSize int

const (
	SMALL_LOT  LotSize = 1
	MEDIUM_LOT LotSize = 2
	LARGE_LOT  LotSize = 3
)

type Lot struct {
	id     string
	size   LotSize
	locked bool
}

type ILot interface {
	// GetId returns the ID of the lot.
	GetId() string
	// Lock locks the lot.
	Lock()

	// Unlock unlocks the lot.
	Unlock()

	// IsLocked checks if the lot is locked.
	IsLocked() bool

	// GetSize returns the size of the lot.
	GetSize() LotSize
}

func NewLot(id string, size LotSize) *Lot {
	return &Lot{
		id:     id,
		size:   size,
		locked: false,
	}
}

func (l *Lot) GetId() string {
	return l.id
}

func (l *Lot) Lock() {
	l.locked = true
}

func (l *Lot) Unlock() {
	l.locked = false
}

func (l *Lot) IsLocked() bool {
	return l.locked
}

func (l *Lot) GetSize() LotSize {
	return l.size
}

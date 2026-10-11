package main

type LotSize string

const (
	SMALL_LOT  LotSize = "SMALL"
	MEDIUM_LOT LotSize = "MEDIUM"
	LARGE_LOT  LotSize = "LARGE"
)

type Lot struct {
	id string
	size LotSize
}

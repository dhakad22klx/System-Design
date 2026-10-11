package main

type Level struct {
	LevelId    string
	SmallLots  []*Lot
	MediumLots []*Lot
	LargeLots  []*Lot
}

type ILevel interface {
	// GetLevelId returns the ID of the level.
	GetLevelId() string

	// GetSmallLots returns the small lots in the level.
	GetSmallLots() []*Lot

	// GetMediumLots returns the medium lots in the level.
	GetMediumLots() []*Lot

	// GetLargeLots returns the large lots in the level.
	GetLargeLots() []*Lot
}

func NewLevel(levelId string, smallLots, mediumLots, largeLots []*Lot) *Level {
	return &Level{
		LevelId:    levelId,
		SmallLots:  smallLots,
		MediumLots: mediumLots,
		LargeLots:  largeLots,
	}
}

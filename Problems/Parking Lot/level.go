package main

type Level struct {
	LevelId    string
	SmallLots  []*Lot
	MediumLots []*Lot
	LargeLots  []*Lot
}

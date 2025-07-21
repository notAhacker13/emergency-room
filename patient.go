package main

type SeverityStatus int

const (
	MildInjury SeverityStatus = iota
	Moderate
	Critical
)

type Patient struct {
	name     string
	severity SeverityStatus

	//required by the heap.Interface
	index int
}

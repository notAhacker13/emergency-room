package main

type PatientsQueue []*Patient

// Push is required by the heap Interface in heap.go implemenaition
func (pq *PatientsQueue) Push(patientData interface{}) {
	patient := patientData.(*Patient)
	patient.index = len(*pq) //required by the heap interface

	*pq = append(*pq, patient)
}

// Pop is required by the heap Interface in heap.go implemenaition
func (pq *PatientsQueue) Pop() interface{} {
	currentQueue := *pq

	n := len(currentQueue)

	patient := currentQueue[n-1]
	patient.index = -1 //reqd by the heap

	*pq = currentQueue[0 : n-1]

	return patient
}

// Len function is required by the sort.Interface ..in heap's implementation
func (pq *PatientsQueue) Len() int {
	return len(*pq)
}

func (pq PatientsQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]

	pq[a].index = a
	pq[b].index = b
}

// Less reqd by sort.Interface
// we flip the comparer to greater than, because we need the comparer to sort by the highest prio, not lowest
func (pq PatientsQueue) Less(a, b int) bool {
	return pq[a].severity > pq[b].severity //condition for max heap
}

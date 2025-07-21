//implementing a priority queue using heaps

package main

import "container/heap"

type EmergencyRoom struct {
	patientsQueue PatientsQueue
}

func NewEmergencyRoom() *EmergencyRoom {
	er := &EmergencyRoom{patientsQueue: make(PatientsQueue, 0)}

	heap.Init(&er.patientsQueue)

	return er
}

func (er *EmergencyRoom) AdmitPatient(patient interface{}) {
	potentialPatient := patient.(*Patient) //type assertion

	//Type Assertion -> is the process of retrieving the concrete value of a specific type from an interface value

	heap.Push(&er.patientsQueue, potentialPatient)

}

func (er *EmergencyRoom) HandleNextPatient() *Patient {

	if er.patientsQueue.Len() == 0 {
		return nil
	}

	nextPatient := heap.Pop(&er.patientsQueue)

	return nextPatient.(*Patient)
}

func (er *EmergencyRoom) UpdatePatientStatus(patient *Patient, newStatus SeverityStatus) {
	patient.severity = newStatus

	//the heap is a heap because it satisfies the heap properties, which means that the parent node contains max value (in max heap) and min value (in case of min heap)
	heap.Fix(&er.patientsQueue, patient.index)

}

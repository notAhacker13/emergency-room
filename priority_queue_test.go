package main

import "testing"

func TestPriorityQueuePopWhenQueueEmptyShouldNotCrash(t *testing.T) {
	er := NewEmergencyRoom()

	next := er.HandleNextPatient()

	if next != nil {
		t.Error("expected empty queue to return nil")
	}
}

func TestPriorityQueueWhenTwoPatientsReturnHighestPriority(t *testing.T) {
	er := NewEmergencyRoom()

	er.AdmitPatient(&Patient{name: "John", severity: Moderate})
	er.AdmitPatient(&Patient{name: "Sam", severity: Critical})

	next := er.HandleNextPatient()

	if next.name != "Sam" {
		t.Error("expected Sam to be the next patient as per priotity")
	}
}

func TestPriorityQueueAfterUpdatePatientStatusReturnHighestPrio(t *testing.T) {
	er := NewEmergencyRoom()

	diana := &Patient{name: "Diana", severity: Moderate}
	er.AdmitPatient(diana)

	liam := &Patient{name: "Liam", severity: Moderate}
	er.AdmitPatient(liam)

	er.AdmitPatient(&Patient{name: "Jamal", severity: Moderate})

	er.UpdatePatientStatus(liam, Critical)

	next := er.HandleNextPatient()

	if next.name != "Liam" {
		t.Error("Liam should be returned")
	}

	er.UpdatePatientStatus(diana, Critical)

	next2 := er.HandleNextPatient()

	if next2.name != "Diana" {
		t.Error("Diana should be returned")
	}
}

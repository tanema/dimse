package main

import (
	"log"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse"
	"github.com/tanema/dimse/src/query"
)

func checkErr(scope string, err error) {
	if err != nil {
		log.Fatalf("🛑 [%s]: %v", scope, err)
	}
	log.Printf("✅ [%s]\n", scope)
}

func main() {
	client, err := dimse.NewClient("www.dicomserver.co.uk:104")
	if err != nil {
		log.Fatalf("connection err: %v", err)
	}
	checkErr("echo", client.Echo())

	q, err := client.Query(
		query.Patient,
		[]*dicom.Element{newElem(tag.PatientID, []string{"3af4bf39-601f-4917-a577-9bbbc8b99366"})},
	)
	checkErr("query", err)
	checkErr("find", q.Find())
	client.Close()
}

func newElem(t tag.Tag, val any) *dicom.Element {
	elem, err := dicom.NewElement(t, val)
	if err != nil {
		log.Fatalf("Err while creating element %v %v %T: %v", t, val, val, err)
	}
	return elem
}

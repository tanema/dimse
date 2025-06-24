package main

import (
	"context"
	"log"
	"sync"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse"
	"github.com/tanema/dimse/src/query"
)

type action func(context.Context, *sync.WaitGroup, *dimse.Client)

func checkErr(scope string, err error) {
	if err != nil {
		log.Fatalf("🛑 [%s]: %v", scope, err)
	}
	log.Printf("✅ [%s]\n", scope)
}

func main() {
	client := dimse.NewClient("www.dicomserver.co.uk:104", nil)
	ctx := context.Background()
	var wg sync.WaitGroup
	acts := []action{echo, find, get}
	for _, act := range acts {
		wg.Add(1)
		act(ctx, &wg, client)
	}
	wg.Wait()
}

func echo(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	checkErr("echo", client.Echo(ctx))
}

func find(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	q, err := client.Query(
		query.Study,
		[]*dicom.Element{
			newElem(tag.UID, []string{"*"}),
			newElem(tag.StudyInstanceUID, []string{""}),
			newElem(tag.SeriesInstanceUID, []string{""}),
			newElem(tag.PatientID, []string{"N8V08W1E6N7C"}),
			newElem(tag.StudyDescription, []string{""}),
		},
	)
	checkErr("find query", err)
	data, err := q.Find(ctx)
	checkErr("find", err)
	printResp("C-FIND", data)
}

func get(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	q, err := client.Query(
		query.Patient,
		[]*dicom.Element{
			newElem(tag.PatientID, []string{"N8V08W1E6N7C"}),
		},
	)
	data, err := q.Get(ctx)
	checkErr("get", err)
	printResp("C-GET", data)
}

func printResp(label string, d []dicom.Dataset) {
	if len(d) == 0 {
		return
	}
	log.Printf("%s response\n", label)
	for i, doc := range d {
		log.Printf("-> doc %v\n", i)
		for _, e := range doc.Elements {
			info, _ := tag.Find(e.Tag)
			log.Printf("\t-> %v = %v\n", info.Name, e.Value)
		}
	}
}

func newElem(t tag.Tag, val any) *dicom.Element {
	elem, err := dicom.NewElement(t, val)
	if err != nil {
		log.Fatalf("Err while creating element %v %v %T: %v", t, val, val, err)
	}
	return elem
}

package main

import (
	"context"
	"log"
	"sync"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse"
	"github.com/tanema/dimse/src/defn/query"
)

type action func(context.Context, *sync.WaitGroup, *dimse.Client)

const (
	AETitle   = "golang-dimse"
	PatientID = "PAT001"
)

var TestAE = dimse.Entity{Title: "test-serv", Host: "www.dicomserver.co.uk", Port: 104}

func checkErr(scope string, err error) {
	if err != nil {
		log.Fatalf("🛑 [%s]: %v", scope, err)
	}
	log.Printf("✅ [%s]\n", scope)
}

func main() {
	oneCmd()
}

func oneCmd() {
	client, err := dimse.NewClient(dimse.Config{AETitle: AETitle})
	checkErr("new client", err)
	defer client.Close()
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	echo(ctx, &wg, client)
}

func allCmds() {
	client, err := dimse.NewClient(dimse.Config{AETitle: AETitle})
	checkErr("new client", err)
	defer client.Close()
	ctx := context.Background()
	query, err := client.Query(
		TestAE,
		query.Patient,
		[]*dicom.Element{
			newElem(tag.UID, []string{"*"}),
			newElem(tag.StudyInstanceUID, []string{""}),
			newElem(tag.SeriesInstanceUID, []string{""}),
			newElem(tag.PatientID, []string{PatientID}),
			newElem(tag.StudyDescription, []string{""}),
		},
	)
	checkErr("query", err)

	var wg sync.WaitGroup
	wg.Add(3)
	go echo(ctx, &wg, client)
	go find(ctx, &wg, query)
	go get(ctx, &wg, query)
	wg.Wait()
}

func echo(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	checkErr("echo", client.Echo(ctx, TestAE))
}

func find(ctx context.Context, wg *sync.WaitGroup, q *dimse.Query) {
	defer wg.Done()
	data, err := q.Find(ctx)
	checkErr("find", err)
	printResp("C-FIND", data)
}

func get(ctx context.Context, wg *sync.WaitGroup, q *dimse.Query) {
	defer wg.Done()
	data, err := q.Get(ctx)
	checkErr("get", err)
	printResp("C-GET", data)
}

func move(ctx context.Context, wg *sync.WaitGroup, q *dimse.Query) {
	defer wg.Done()
	data, err := q.Move(ctx, AETitle)
	checkErr("move", err)
	printResp("C-GET", data)
}

func store(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	ds, err := dicom.ParseFile("./data/4.dcm", nil)
	checkErr("parsing dicom", err)
	checkErr("store", client.Store(ctx, TestAE, ds))
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

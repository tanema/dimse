package main

import (
	"context"
	"flag"
	"log"
	"os"
	"slices"
	"sync"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse"
	"github.com/tanema/dimse/src/defn/query"
	"github.com/tanema/dimse/src/defn/transfersyntax"
)

type action func(context.Context, *sync.WaitGroup, *dimse.Client)

var (
	AEHost     string
	AEPort     int64
	AETitle    string
	PatientID  string
	TestAE     dimse.Entity
	builtQuery *dimse.Query
)

func init() {
	flag.StringVar(&AEHost, "host", "www.dicomserver.co.uk", "the pacs host you are connecting to")
	flag.Int64Var(&AEPort, "port", 104, "the pacs port you are connecting to")
	flag.StringVar(&AETitle, "title", "golang-dimse", "the aetitle of your client")
	flag.StringVar(&PatientID, "patient", "PAT001", "the patient query")
}

func checkErr(scope string, err error) {
	if err != nil {
		log.Fatalf("🛑 [%s]: %v", scope, err)
	}
	log.Printf("✅ [%s]\n", scope)
}

func main() {
	flag.Parse()
	TestAE = dimse.Entity{Title: "test-serv", Host: AEHost, Port: int(AEPort)}

	client, err := dimse.NewClient(dimse.Config{AETitle: AETitle, Port: 103})
	checkErr("new client", err)
	defer client.Close()
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(4)
	go echo(ctx, &wg, client)
	go find(ctx, &wg, client)
	go get(ctx, &wg, client)
	go store(ctx, &wg, client)
	wg.Wait()
}

func getQuery(client *dimse.Client) *dimse.Query {
	if builtQuery != nil {
		return builtQuery
	}
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
	builtQuery = query
	return query
}

func echo(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	checkErr("echo", client.Echo(ctx, TestAE))
}

func find(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	q := getQuery(client)
	data, err := q.Find(ctx)
	checkErr("find", err)
	printResp("C-FIND", data)
}

func get(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	q := getQuery(client)
	data, err := q.Get(ctx)
	checkErr("get", err)
	printResp("C-GET", data)
	for _, ds := range data {
		f, err := os.Create("./tmp/file1.dcm")
		checkErr("open file", err)
		checkErr("file write", dicom.Write(f, ds, dicom.SkipVRVerification(), dicom.OverrideMissingTransferSyntax(string(transfersyntax.ImplicitVRLittleEndian))))
		checkErr("file close", f.Close())
	}
}

func move(ctx context.Context, wg *sync.WaitGroup, q *dimse.Query) {
	defer wg.Done()
	data, err := q.Move(ctx, AETitle)
	checkErr("move", err)
	printResp("C-GET", data)
}

func store(ctx context.Context, wg *sync.WaitGroup, client *dimse.Client) {
	defer wg.Done()
	ds, err := dicom.ParseFile("./test/data/testfile.dcm", nil)
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
			if slices.Contains(info.VRs, "OB") {
				log.Printf("\t-> %v = %T \n", info.Name, e.Value)
			} else {
				log.Printf("\t-> %v = %v\n", info.Name, e.Value)
			}
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

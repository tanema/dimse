package main

import (
	"context"
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
	ctx := context.Background()
	client, err := dimse.NewClient("www.dicomserver.co.uk:104", nil)
	checkErr("connection", err)
	checkErr("echo", echo(ctx, client))

	q, err := client.Query(
		query.Patient,
		[]*dicom.Element{
			newElem(tag.PatientID, []string{"3af4bf39-601f-4917-a577-9bbbc8b99366"}),
		},
	)
	checkErr("query", err)
	checkErr("find", find(ctx, q))
	checkErr("get", get(ctx, q))
}

func echo(ctx context.Context, client *dimse.Client) error {
	return client.Echo(ctx)
}

func find(ctx context.Context, q *dimse.Query) error {
	data, err := q.Find(ctx)
	if err != nil {
		return err
	}
	log.Printf("Got find response, found %v docs\n", len(data))
	for i, doc := range data {
		log.Printf("-> doc %v\n", i)
		for _, e := range doc.Elements {
			info, _ := tag.Find(e.Tag)
			log.Printf("\t-> %v = %v\n", info.Name, e.Value)
		}
	}
	return nil
}

func get(ctx context.Context, q *dimse.Query) error {
	data, err := q.Get(ctx)
	if err != nil {
		return err
	}
	log.Printf("Got find response, found %v docs\n", len(data))
	for i, doc := range data {
		log.Printf("-> doc %v\n", i)
		for _, e := range doc.Elements {
			info, _ := tag.Find(e.Tag)
			log.Printf("\t-> %v = %v\n", info.Name, e.Value)
		}
	}
	return nil
}

func newElem(t tag.Tag, val any) *dicom.Element {
	elem, err := dicom.NewElement(t, val)
	if err != nil {
		log.Fatalf("Err while creating element %v %v %T: %v", t, val, val, err)
	}
	return elem
}

func mustGet(doc dicom.Dataset, t tag.Tag) any {
	elem, _ := doc.FindElementByTag(t)
	if elem == nil {
		return nil
	}
	return elem.Value
}

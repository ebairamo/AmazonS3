package models

import (
	"encoding/xml"
	"time"
)

type Bucket struct {
	Name         string    `xml:"Name"`
	CreationTime time.Time `xml:"CreationDate"`
}

type BucketsAll struct {
	Buckets []Bucket `xml:"Bucket"`
}

type ListAllBucketsResult struct {
	XMLName xml.Name   `xml:"ListAllBucketsResult"`
	Buckets BucketsAll `xml:"Buckets"`
}

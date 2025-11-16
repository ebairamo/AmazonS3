package models

import "time"

type Bucket struct {
	Name         string
	CreationTime time.Time
}

type BucketXml struct {
	Name         string `xml:"name"`
	CreationDate string `xml:"creationDate"`
}

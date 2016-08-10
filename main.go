package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/HugoSTorres/quasimodo/amazon"
	"github.com/HugoSTorres/quasimodo/hugo"
)

func init() {
	flag.Usage = func() {
		fmt.Println("usage: quasimodo --bucket <s3 bucket> [--region <s3 region>]")
		flag.PrintDefaults()
	}
}

func main() {
	var bucket, region string

	flag.StringVar(&region, "region", "us-east-1", "the region containing the s3 bucket")
	flag.StringVar(&bucket, "bucket", "", "the bucket hosting the site")

	flag.Parse()

	if len(bucket) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("building site")

	err := hugo.Build()
	if err != nil {
		fmt.Printf("error building site: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("getting aws credentials and opening session")

	svc, err := amazon.NewS3Session(region, bucket)
	if err != nil {
		fmt.Printf("error creating s3 session: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("cleaning bucket")

	err = svc.Clean()
	if err != nil {
		fmt.Printf("error cleaning bucket %v: %v\n", svc.Bucket, err)
		os.Exit(1)
	}

	fmt.Println("uploading public folder")

	err = svc.Publish()
	if err != nil {
		fmt.Printf("error publishing site: %v\n", err)
		os.Exit(1)
	}
}

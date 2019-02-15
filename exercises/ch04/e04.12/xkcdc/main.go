package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Printf("%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		dataFilename = "store.data"
		root         = "https://xkcd.com"
		sufx         = "/info.0.json"
		expiry       = time.Hour * 24 * 60
		pause        = time.Millisecond * 50
	)

	fmt.Println("hello")

	rs := newRecords()
	if err := unmarshalStored(dataFilename, rs); err != nil {
		return err
	}

	a := newAccess(root, sufx)

	fmt.Println("starting warmup (may take a few minutes)")
	if err := updateRecords(expiry, pause, a, rs); err != nil {
		return err
	}
	fmt.Println("completed warmup")

	if err := marshalStored(dataFilename, rs); err != nil {
		return err
	}

	if err := runUserInteraction(rs); err != nil {
		return err
	}

	fmt.Println("goodbye")

	return nil
}

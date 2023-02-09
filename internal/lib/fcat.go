package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gaurav-gogia/libexfat"
)

const datadir = "./data"

func extractentries(exfatdata libexfat.ExFAT, entries []libexfat.Entry) error {
	if _, err := os.Stat(datadir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(datadir, os.ModeDir)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	for _, entry := range entries {
		err := extractentry(exfatdata, entry)
		if err != nil {
			return err
		}
	}
	return nil
}

func extractentry(exfatdata libexfat.ExFAT, entry libexfat.Entry) error {
	if entry.IsDir() || entry.IsDeleted() || entry.IsInvalid() {
		return nil
	}
	path := filepath.Join(datadir, entry.GetName())
	err := exfatdata.ExtractFileContent(entry, path)
	if err != nil {
		return err
	}
	fmt.Println("Extracted: ", entry.GetName())
	return nil
}

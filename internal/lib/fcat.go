package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aoiflux/libxfat"
)

const datadir = "./data"

func recursiveExtract(exfatdata libxfat.ExFAT, rootentries []libxfat.Entry) error {
	_, err := os.Stat(datadir)
	if os.IsNotExist(err) {
		err = os.Mkdir(datadir, os.ModeDir)
		if err != nil {
			return err
		}
	}

	abspath, err := filepath.Abs(datadir)
	if err != nil {
		return err
	}

	return exfatdata.ExtractAllFiles(rootentries, abspath)
}

func collectedExtract(exfatdata libxfat.ExFAT, entries []libxfat.Entry) error {
	_, err := os.Stat(datadir)
	if os.IsNotExist(err) {
		err = os.Mkdir(datadir, os.ModeDir)
		if err != nil {
			return err
		}
	}

	abspath, err := filepath.Abs(datadir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsInvalid() || entry.IsDeleted() {
			continue
		}

		fmt.Println("Extracting: ", entry.GetName())
		fpath := filepath.Join(abspath, entry.GetName())
		err = exfatdata.ExtractEntryContent(entry, fpath)
		if err != nil {
			return err
		}
	}

	return nil
}

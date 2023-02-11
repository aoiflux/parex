package lib

import (
	"os"
	"path/filepath"

	"github.com/aoiflux/libxfat"
)

const datadir = "./data"

func extractentries(exfatdata libxfat.ExFAT, rootentries []libxfat.Entry) error {
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

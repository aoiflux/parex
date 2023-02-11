package lib

import (
	"fmt"
	"os"

	"github.com/aoiflux/libxfat"
)

func Explore(imagefile *os.File, offset uint64, level int) error {
	exfatdata, err := libxfat.New(imagefile, true, offset)
	if err != nil {
		return err
	}

	rootentries, err := exfatdata.ReadRootDir()
	if err != nil {
		return err
	}

	if level > 0 {
		return recursivefls(exfatdata, rootentries, level)
	}

	printentries(rootentries)
	return nil
}

func recursivefls(exfatdata libxfat.ExFAT, rootentries []libxfat.Entry, level int) error {
	if level < 2 {
		return exfatdata.ShowAllEntriesInfo(rootentries, "/", false)
	}

	if level < 3 {
		entries, err := exfatdata.GetFiles(rootentries)
		if err != nil {
			return err
		}
		printentries(entries)
		return nil
	}

	return extractentries(exfatdata, rootentries)
}

func printentries(entries []libxfat.Entry) {
	for index, entry := range entries {
		fmt.Println(index+1, "-->", entry.GetName())
	}
}

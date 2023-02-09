package lib

import (
	"fmt"
	"os"

	"github.com/gaurav-gogia/libexfat"
)

func Explore(imagefile *os.File, offset uint64, level int) error {
	exfatdata, err := libexfat.New(imagefile, true, offset)
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

func recursivefls(exfatdata libexfat.ExFAT, rootentries []libexfat.Entry, level int) error {
	if level < 2 {
		return exfatdata.ShowAllFilesInfo(rootentries, "/", false)
	}

	entries, err := exfatdata.GetFiles(rootentries)
	if err != nil {
		return err
	}

	if level < 3 {
		printentries(entries)
		return nil
	}

	return extractentries(exfatdata, entries)
}

func printentries(entries []libexfat.Entry) {
	for index, entry := range entries {
		fmt.Println(index+1, "-->", entry.GetName())
	}
}

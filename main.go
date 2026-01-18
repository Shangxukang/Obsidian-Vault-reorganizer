package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ExtractPngsFromFile(content string) []string {
	re := regexp.MustCompile(`!\[\[(.*?)\]\]`)
	matches := re.FindAllStringSubmatch(content, -1)

	out := make([]string, 0, len(matches))
	for _, m := range matches {
		if strings.Contains(m[1], "|") {
			m[1] = strings.Split(m[1], "|")[0]
		}
		out = append(out, m[1])
	}

	return out
}

func main() {
	args := os.Args
	newPath := args[1]
	oldPath := args[2]

	// Collect all markdown files in new vault
	var mdFiles []string
	err := filepath.WalkDir(newPath, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			mdFiles = append(mdFiles, path)
		}
		return nil
	})

	var oldPics []string
	e := filepath.WalkDir(oldPath, func(path string, d os.DirEntry, e error) error {
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), ".jpg") || strings.HasSuffix(strings.ToLower(d.Name()), ".png") {
			oldPics = append(oldPics, path)
		}
		return nil
	})

	if err != nil || e != nil {
		panic(err)
	}

	// Create folder that contains the pics
	picsFolder := filepath.Join(newPath, "/pics")
	ee := os.MkdirAll(picsFolder, 0755)
	if ee != nil {
		return
	}

	for _, f := range mdFiles {
		fmt.Println("")
		fmt.Println("processing file: ", f)

		content, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		pics := ExtractPngsFromFile(string(content))
		if len(pics) == 0 {
			fmt.Println("No pics in this file")
		}

		// for each pic, find it in the old vault, and move it to new one
		for _, p := range pics {
			fmt.Print("processing pic:", p, ".")
			var original string
			for _, full := range oldPics {
				if strings.Contains(full, p) {
					original = full
					break
				}
			}
			err := os.Rename(original, filepath.Join(picsFolder, p))
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Print("---Succeed.\n")
		}

	}

}

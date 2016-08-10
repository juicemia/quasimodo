package fs

import (
	"mime"
	"os"
	"regexp"
	"strings"
)

// File encapsulates the file data with some information we need
// to pass over to the put object request.
type File struct {
	Data *os.File
	Path string
	Mime string
}

// GetSite returns the file data for the site.
func GetSite() ([]File, error) {
	siterepo, err := os.Open("./public")
	if err != nil {
		return nil, err
	}

	return getSite(siterepo, nil)
}

func getSite(baseDir *os.File, acc []File) ([]File, error) {
	names, err := baseDir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	if len(names) < 1 {
		return acc, nil
	}

	for _, name := range names {
		if name[0] == '.' {
			continue
		}

		f, err := os.Open(baseDir.Name() + "/" + name)
		if err != nil {
			return nil, err
		}

		info, err := f.Stat()
		if err != nil {
			return nil, err
		}

		if info.IsDir() {
			acc, err = getSite(f, acc)
			if err != nil {
				return nil, err
			}

			continue
		}

		// get everything after public here
		re := regexp.MustCompile("public(\\/(\\w(\\w+|-)*(\\.*\\w+)*))+")
		publicDir := re.FindString(baseDir.Name() + "/" + name)

		key := strings.Join(strings.Split(publicDir, "/")[1:], "/")

		rext := regexp.MustCompile("\\.\\w+$")
		ext := rext.FindString(key)

		mimeType := mime.TypeByExtension(ext)

		acc = append(acc, File{
			Path: key,
			Data: f,
			Mime: mimeType,
		})
	}

	return acc, err
}

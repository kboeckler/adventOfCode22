package solution

import (
	"math/big"
	"sort"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(7, day7{})
}

type day7 struct {
}

func (d day7) SolvePart1(input []string) interface{} {
	filesByPath := make(map[string][]file, 0)
	currentDir := make([]string, 0)
	visitDirectory(partsToDirectoryPath(currentDir), filesByPath)
	for _, row := range input {
		cmd := strings.TrimLeft(row, "$ ")
		if len(cmd) < len(row) {
			commandParam := strings.Split(cmd, " ")
			if strings.EqualFold(commandParam[0], "cd") {
				if strings.EqualFold(commandParam[1], "/") {
					currentDir = make([]string, 0)
					visitDirectory(partsToDirectoryPath(currentDir), filesByPath)
				} else if strings.EqualFold(commandParam[1], "..") {
					currentDir = currentDir[0 : len(currentDir)-1]
					visitDirectory(partsToDirectoryPath(currentDir), filesByPath)
				} else {
					currentDir = append(currentDir, commandParam[1])
					visitDirectory(partsToDirectoryPath(currentDir), filesByPath)
				}
			}
		} else if len(strings.TrimLeft(row, "dir")) == len(row) {
			sizeFile := strings.Split(row, " ")
			size, _ := strconv.Atoi(sizeFile[0])
			files := filesByPath[partsToDirectoryPath(currentDir)]
			files = append(files, file{sizeFile[1], big.NewInt(int64(size))})
			filesByPath[partsToDirectoryPath(currentDir)] = files
		}
	}

	directories := directories{make([]directory, 0)}
	for path, files := range filesByPath {
		directories.elements = append(directories.elements, directory{path, files, big.NewInt(0)})
	}
	sort.Sort(directories)

	for i := range directories.elements {
		directory := &directories.elements[i]
		myFilesSize := big.NewInt(0)
		for _, file := range directory.files {
			myFilesSize = myFilesSize.Add(myFilesSize, file.size)
		}
		subdirSize := big.NewInt(0)
		for _, otherDir := range directories.elements {
			if len(strings.TrimLeft(otherDir.path, directory.path)) < len(otherDir.path) {
				subdirSize = subdirSize.Add(subdirSize, otherDir.totalSize)
			}
		}
		directory.totalSize = directory.totalSize.Add(myFilesSize, subdirSize)
	}

	result := big.NewInt(0)
	for _, directory := range directories.elements {
		if directory.totalSize.Cmp(big.NewInt(100000)) <= 0 {
			result = result.Add(result, directory.totalSize)
		}
	}

	return result
}

func visitDirectory(path string, directories map[string][]file) {
	_, exists := directories[path]
	if !exists {
		directories[path] = make([]file, 0)
	}
}
func partsToDirectoryPath(dir []string) string {
	sb := strings.Builder{}
	sb.WriteString("/")
	for _, d := range dir {
		sb.WriteString(d + "/")
	}
	return strings.ReplaceAll(sb.String(), "//", "/")
}

type directories struct {
	elements []directory
}

func (d directories) Len() int {
	return len(d.elements)
}

func (d directories) Less(i, j int) bool {
	return strings.Count(d.elements[i].path, "/")-strings.Count(d.elements[j].path, "/") > 0
}

func (d directories) Swap(i, j int) {
	temp := d.elements[i]
	d.elements[i] = d.elements[j]
	d.elements[j] = temp
}

type directory struct {
	path      string
	files     []file
	totalSize *big.Int
}

type file struct {
	name string
	size *big.Int
}

func (d day7) SolvePart2(input []string) interface{} {
	return ""
}

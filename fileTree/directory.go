package fileTree

import "sort"

type FileList []string
func (fl FileList) Len() int {return len(fl)}
func (fl FileList) Less(a,b int) bool {return fl[a]<fl[b]}
func (fl FileList) Swap(a,b int) {fl[a], fl[b] = fl[b], fl[a]}
func (fl *FileList) Add(a string) {
	*fl = append(*fl, a)
}
type Directory struct {
	Path string
	Files FileList
}

// declare the DirList type
type DirList []Directory
func (dl DirList) Len() int {return len(dl)}
func (dl DirList) Less(a,b int) bool {return dl[a].Path<dl[b].Path}
func (dl DirList) Swap(a,b int) {dl[a], dl[b] = dl[b], dl[a]}
func (dl *DirList) Add(newDir Directory) { *dl = append(*dl, newDir)}
func (dl *DirList) Sort() {
	sort.Sort(*dl)
	for _,v := range *dl{
		sort.Sort(v.Files)
	}
}

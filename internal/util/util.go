package util
import (
	scan "github.com/shield61/prepare-photo-import/internal/scan"
)

type ByCountAndExtension []scan.ExtensionCount

func (a ByCountAndExtension) Len() int      { return len(a) }
func (a ByCountAndExtension) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCountAndExtension) Less(i, j int) bool {
	if a[i].Count != a[j].Count {
		return a[i].Count > a[j].Count // Sort by count in descending order
	}
	return a[i].Extension < a[j].Extension // Sort alphabetically for extensions with the same count
}

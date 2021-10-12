package fmt

import (
	"fmt"
)

func (f *factory) register() {
	f.Set(`Errorf`, fmt.Errorf)
	f.Set(`Fprint`, fmt.Fprint)
	f.Set(`Fprintf`, fmt.Fprintf)
	f.Set(`Fprintln`, fmt.Fprintln)
	f.Set(`Fscan`, fmt.Fscan)
	f.Set(`Fscanf`, fmt.Fscanf)
	f.Set(`Fscanln`, fmt.Fscanln)
	f.Set(`Print`, fmt.Print)
	f.Set(`Printf`, fmt.Printf)
	f.Set(`Println`, fmt.Println)
	f.Set(`Scan`, fmt.Scan)
	f.Set(`Scanf`, fmt.Scanf)
	f.Set(`Scanln`, fmt.Scanln)
	f.Set(`Sprint`, fmt.Sprint)
	f.Set(`Sprintf`, fmt.Sprintf)
	f.Set(`Sprintln`, fmt.Sprintln)
	f.Set(`Sscan`, fmt.Sscan)
	f.Set(`Sscanf`, fmt.Sscanf)
	f.Set(`Sscanln`, fmt.Sscanln)

	f.Set(`isFormatter`, isFormatter)
	f.Set(`isGoStringer`, isGoStringer)
	f.Set(`isScanState`, isScanState)
	f.Set(`isScanner`, isScanner)
	f.Set(`isState`, isState)
	f.Set(`isStringer`, isStringer)
}
func isFormatter(i interface{}) bool {
	_, result := i.(fmt.Formatter)
	return result
}
func isGoStringer(i interface{}) bool {
	_, result := i.(fmt.GoStringer)
	return result
}
func isScanState(i interface{}) bool {
	_, result := i.(fmt.ScanState)
	return result
}
func isScanner(i interface{}) bool {
	_, result := i.(fmt.Scanner)
	return result
}
func isState(i interface{}) bool {
	_, result := i.(fmt.State)
	return result
}
func isStringer(i interface{}) bool {
	_, result := i.(fmt.Stringer)
	return result
}

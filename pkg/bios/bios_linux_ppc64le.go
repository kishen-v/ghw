// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package bios

import "fmt"

func (i *Info) load() error {
	fmt.Println("BIOS related information is not available for ppc64le")
	return nil
}

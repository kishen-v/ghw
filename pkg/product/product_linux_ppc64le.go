// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package product

import "fmt"

func (i *Info) load() error {
	fmt.Println("Product related information is not available on ppc64le")
	return nil
}

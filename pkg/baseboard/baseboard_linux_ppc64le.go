// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package baseboard

import (
	"fmt"
)

func (i *Info) load() error {
	fmt.Println("baseboardFillInfo not implemented on ppc64le")
	return nil
}

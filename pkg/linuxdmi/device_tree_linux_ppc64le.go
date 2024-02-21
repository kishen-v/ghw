//go:build ppc64le
// +build ppc64le

package linuxdmi

import (
	"fmt"

	"github.com/jaypipes/ghw/pkg/context"
	"github.com/jaypipes/ghw/pkg/util"
)

func Item(ctx *context.Context, value string) string {
	fmt.Println("Device tree not implemented yet")
	return util.UNKNOWN
}

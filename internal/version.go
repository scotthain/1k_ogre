/*
Copyright © 2022 Scott Hain elejia@gmail.com
*/
package version

import (
	"os"

	"github.com/pkg/errors"
)

func Version() string {
	dat, err := os.ReadFile("VERSION")
	errors.Wrap(err, "Failed to read from VERSION file")
	str := string(dat[:])
	return str
}

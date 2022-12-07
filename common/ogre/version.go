/*
Copyright © 2022 Scott Hain elejia@gmail.com
*/
package ogre

import (
	"os"
)

func Version() string {
	dat, err := os.ReadFile("VERSION")
	CatchError(err)
	str := string(dat[:])
	return str
}

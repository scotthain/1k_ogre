/*
Copyright © 2022 Scott Hain elejia@gmail.com
*/
package ogre

func CatchError(e error) {
	if e != nil {
		panic(e)
	}
}

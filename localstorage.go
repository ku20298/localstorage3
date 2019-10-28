package localstorage

import (
	"github.com/gopherjs/gopherjs/js"
)

// SetItem calls localstorage.setItem()
func SetItem(tag string, a interface{}) {
	js.Global.Get("localStorage").Call("setItem", tag, a)
}

// GetItem calls localstorage.getItem()
func GetItem(tag string) string {
	return js.Global.Get("localStorage").Call("getItem", tag).String()
}

// RemoveItem calls localstorage.removeItem()
func RemoveItem(tag string) {
	js.Global.Get("localStorage").Call("removeItem", tag)
}

// Clear calls localstorage.clear()
func Clear() {
	js.Global.Get("localStorage").Call("clear")
}

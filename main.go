package main
import "plugin"
func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	HelloFn, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	HelloFn.(func())()
}

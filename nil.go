package main

func NewMap(name string) map[string]string {

	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}

func main39() {

	var person = NewMap("adi")

	if person != nil {
		println(person["name"])
	} else {
		println("nil")
	}

}

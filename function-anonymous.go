package main

type BlackList func(string) bool

func main28() {
	blackList := func(name string) bool {
		return name == "Anjing" || name == "Jancok"
	}

	registerUser("Anjing", blackList)
	registerUser("Jancok", blackList)
	registerUser("Bambang", blackList)
}

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		println("User", name, "is blacklisted")
	} else {
		println("User", name, "is not blacklisted")
	}
}

// func nameBlackList(name string) bool {
// 	return name == "Anjing" || name == "Jancok"
// }

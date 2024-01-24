package main

import (
	"os"

	em "gopkg.in/electricitymaps.v0/electricitymaps"
)

const Usage = `Usage: go run package [-]

"package" is the main package of this module.

"-" processes standard input.`

func main() {
	var session em.Session
	switch len(os.Args) {
	case 1:
		session = em.GetSession()
	case 2:
		if os.Args[1] == "-" {
			sessionp, err := em.GetSessionFromStdin()
			if err != nil {
				panic(err.Error())
			}
			session = *sessionp
		} else {
			println(Usage)
			return
		}
	default:
		println(Usage)
		return
	}
	println("apiRoot", session.ApiRoot)
	println(len(session.AuthToken) > 0)
}

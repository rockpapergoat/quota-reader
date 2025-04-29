package main

import (
	"fmt"
	"os/user"
	"strconv"
)

func getgroup() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	gidInt, _ := strconv.Atoi(currentUser.Gid)
	group, _ := user.LookupGroupId(strconv.Itoa(gidInt))
	if group == nil {
		return "err: unknown group"
	}
	return group.Name
}

func main() {
	fmt.Println(getgroup())
}

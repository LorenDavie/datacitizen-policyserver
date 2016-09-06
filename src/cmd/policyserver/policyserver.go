package main

import (
	"datacitizen.me/identity"
	"fmt"
)

// Mainline
func main() {
	policy := identity.Identity{Signature: "Main Identity"}
	fmt.Println(policy)
}

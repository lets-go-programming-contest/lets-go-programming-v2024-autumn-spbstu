//go:build add
// +build add

package main

func init() {
	oper = append(oper, "+", "-")
}

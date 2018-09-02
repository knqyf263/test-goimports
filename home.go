package main

import "github.com/mitchellh/go-homedir"

func Home() (string, error) {
	return homedir.Dir()
}

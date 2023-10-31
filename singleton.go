package main

import "fmt"

func main() {
	singleton := GetInstance()

	// Use the singleton instance.
	fmt.Println(singleton)
}

type Singleton struct{}

var instance *Singleton

func init() {
	instance = &Singleton{}
}

func GetInstance() *Singleton {
	return instance
}


package main

import "github.com/hoisie/redis"
import "fmt"
import "log"

func main() {
	var client redis.Client
	keys, err := client.Keys("*")
	if err != nil {
		log.Fatal("Could not retrieve keys")
	}
	fmt.Printf("There are %d keys\n", len(keys))
	for _, k := range keys {
		fmt.Printf("key: %s\n", k)
	}
}

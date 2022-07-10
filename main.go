package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"encoding/json"
)

type Employee struct{
	Name string `json: "name"`
	Id int `json: "id"`	
}


func main(){
	fmt.Println("Go Redis Tutorial")

	redis_client := redis.NewClient(&redis.Options{
              Addr: "127.0.0.1:6379",
              Password: "",
              DB: 0,
	})

	pong, err := redis_client.Ping().Result()
	fmt.Println(pong, err)

	// As we know that redis supports Set, Sorted Set, List, Hash, String
    json, err := json.Marshal(Employee{Name: "Shree", Id: 728375})
    if err != nil{
    	fmt.Println(err)
    }
    
    // Set
    err = redis_client.Set("id007",json,0).Err()
    if err != nil{
    	fmt.Println(err)
    }

   // Get
  val, err := redis_client.Get("id007").Result()
   if err != nil{
   	fmt.Println(err)
   }
   fmt.Println(val)
}
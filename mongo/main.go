/**
 * File: main
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/9/28 3:40 下午
 **/

package main

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main()  {
	p1 := newPlayer("baba", 1)
	coll := mgm.Coll(p1)
	mgm.
	coll.Create(p1)
	fmt.Println(p1)
	p1.Name = "didi"
	coll.Update(p1)
	fmt.Println(p1)

	p2 := &Player{}
	coll.First(bson.D{{"name","didi"}}, p2)
	fmt.Println(p2)
}

func init() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "test_war", options.Client().ApplyURI("mongodb://admin:admin@localhost:10001"))
	if err !=  nil{
		fmt.Println(err)
	}
}



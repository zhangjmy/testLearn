/**
 * File: model
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/9/28 5:01 下午
 **/

package main

import "github.com/kamva/mgm/v3"

type Player struct {
	mgm.DefaultModel `bson:",inline"`
	Name string `json:"name" bson:"name""`
	Age int `json:"age" bson:"age"`
}

func newPlayer(name string, age int) *Player {
	return &Player{
		Name: name,
		Age:  age,
	}
}

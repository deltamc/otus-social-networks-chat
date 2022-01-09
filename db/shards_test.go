package db

import (
	"fmt"
	//"strconv"
	"testing"

)

func TestGetShardNodeByKey(t *testing.T) {
	sn := ShardNodes{}

	sn.Add(1,"node-1")
	sn.Add(2,"node-2")

	fmt.Println(sn.GetShardNodeByKey("Москва"))
	fmt.Println(sn.GetShardNodeByKey("Анапа"))



}
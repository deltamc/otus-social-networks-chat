package db

import (
	"hash/crc32"
	"sort"
)

type ShardNodes struct {
	Nodes []ShardNode
}


type ShardNode struct {
	Id     uint8
	Name   string
	HashId uint8
}

func (n *ShardNodes) Add(id uint8, name string) {
	n.Nodes = append(n.Nodes, ShardNode{
		Id:id,
		Name:name,
		HashId: uint8(crc32.Checksum([]byte(name), crc32.MakeTable(2)) % 128),
	})
}

func (n ShardNodes) GetShardNodeByKey(key string) ShardNode {
	searchfn := func(i int) bool {
		return n.Nodes[i].HashId >= uint8(crc32.Checksum([]byte(key), crc32.MakeTable(2))%128)
	}

	i := sort.Search(len(n.Nodes), searchfn)
	if i >= len(n.Nodes) {
		i = 0
	}
	return n.Nodes[i]
}


package mymap

import (
	"fmt"
	"hashTableStructureWithChaining/linked-list"
)
const sizeOfMap = 100000

type Map interface{
	Get(K interface{})
	Set(K,V interface{})
}

type MyMap struct{
	Data [sizeOfMap]*linked_list.List
	HashFunc func(K interface{}) int
}
func(mm * MyMap) Get(K interface{}) interface{}{
	hashValue := mm.HashFunc(K)
	return mm.Data[hashValue].SearchNode(K).Value
}

func(mm * MyMap) Set(K interface{},V interface{}) {

	hashValue := mm.HashFunc(K)
	if mm.Data[hashValue] == nil {
		mm.Data[hashValue] = linked_list.NewList(K,V)
	} else {
		mm.Data[hashValue].Add(K,V)
	}
}
func HashWithAsciiModSize(K interface{}) int {
	valStr := fmt.Sprint(K)
	hashValue := 0
	for _, element := range valStr {
		ascii := int(element)
		hashValue += ascii
	}
	return hashValue % sizeOfMap
}

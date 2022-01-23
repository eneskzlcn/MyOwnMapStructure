package linked_list

import (
	"errors"
	"log"
)

type Node struct{
	Key interface{}
	Next *Node
	Value interface{}
}

func NewNode(K,V interface{}) *Node{
	node := Node{
		Key:   K,
		Next:  nil,
		Value: V,
	}
	return &node
}
type List struct{
	Head *Node
	Size int
}
func (list *List) Add(K,V interface{}){
	list.addNode(NewNode(K,V))
}
// NewList initializes the list with given head node
func NewList(K,V interface{}) *List{
	head := NewNode(K,V)
	var list = List{
		Head: head,
		Size: 1,
	}
	return &list
}
func (list *List) TraverseAndPrint(){
	tempNode := list.Head
	for tempNode != nil {
		log.Printf("\nBook:%s\nAuthor:%s\n",tempNode.Key,tempNode.Value)
		tempNode = tempNode.Next
	}
}
func (list *List) Delete(K interface{}) *Node{
	tempNode := list.Head
	previousNode := list.Head
	for tempNode != nil {
		if tempNode.Key == K{
			previousNode.Next= tempNode.Next
			return tempNode
		}
		previousNode = tempNode
		tempNode = tempNode.Next
	}
	return nil
}
func (list *List) UpdateNode(K interface{}, V interface{})error{
	updateNode := list.SearchNode(K)
	if updateNode != nil {
		return errors.New("node with key not found")
	}
	updateNode.Value = V
	return nil
}
func (list *List) addNode(node *Node){
	if list.Head == nil {
		list.Head = node
	}else{
		lastNode:= list.FindLastNode()
		lastNode.Next = node
	}
	
	list.Size= list.Size+1
}
func (list * List) SearchNode(K interface{}) *Node{
	tempNode := list.Head
	for tempNode != nil {
		if tempNode.Key == K{
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}
func (list *List) FindLastNode() *Node{
	tempNode := list.Head
	for tempNode.Next != nil {
		tempNode = tempNode.Next
	}
	return tempNode
}
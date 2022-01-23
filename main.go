package main

import (
	"fmt"
	linked_list "hashTableStructureWithChaining/linked-list"
	"hashTableStructureWithChaining/mymap"
	"log"

	"time"
)

func main(){
	//initialize map with 100 000 maximum available size of array
	myNewMap := mymap.MyMap{
		Data:     [100000]*linked_list.List{},
		HashFunc: mymap.HashWithAsciiModSize,
	}
	//fill the map with random data
	for i := 0 ; i < 100000; i++ {
		myNewMap.Set(fmt.Sprint(i),fmt.Sprint(i*2))
	}
	//initialize efficiency test struct for each case
	type TimeEfficiencies struct {
		gettingFirstPushedElementToMapEfficiency int
		gettingLastPushedElementToMapEfficiency int
		gettingFirstPushedElementToArrayEfficiency int
		gettingLastPushedElementToArrayEfficiency int
	}
	timeEfficiencies := TimeEfficiencies{}
	//now we calculate accessing the first element '0' of map as nanosecond
	now := time.Now()
	myNewMap.Get("0")
	after := time.Now()
	timeEfficiencies.gettingFirstPushedElementToMapEfficiency = after.Nanosecond() -now.Nanosecond()
	//noe we calculate accessing the last pushed element '99999' of map as nanoseconds
	now = time.Now()
	myNewMap.Get("99999")
	after = time.Now()
	timeEfficiencies.gettingLastPushedElementToMapEfficiency = after.Nanosecond() -now.Nanosecond()

	//trying this operation with array for best and worst case
	type Book struct{
		K string
		V string
	}
	myArray := [100000]Book{}
	//Suppose that the index and K are different in this case. Because otherwise we can reach the element like mySlice[key]. And it is just 0(1)
	for i := 0 ; i < 100000; i++ {
		myArray[i].K = fmt.Sprint(i)
		myArray[i].V = fmt.Sprint(i*2)
	}
	//Best case trying for accessing element that has key '0' which is the first pushed element to array
	now = time.Now()
	for _,item := range myArray {
		if item.K == "0"{
			break
		}
	}
	after = time.Now()
	timeEfficiencies.gettingFirstPushedElementToArrayEfficiency = after.Nanosecond() -now.Nanosecond()

	// Worst case trying for accessing element that has key '99999' which is the last pushed element to array
	now = time.Now()
	for _,item := range myArray {
		if item.K == "99999"{
			break
		}
	}
	after = time.Now()
	timeEfficiencies.gettingLastPushedElementToArrayEfficiency = after.Nanosecond() - now.Nanosecond()

	//Printing results, for ms we need to convert
	log.Printf("\n Time efficiency for reaching key 0 from array: %d nanoseconds",timeEfficiencies.gettingFirstPushedElementToArrayEfficiency)
	log.Printf("\n Time efficiency for reaching key 99999 from array: %d nanoseconds",timeEfficiencies.gettingLastPushedElementToArrayEfficiency)
	log.Printf("\n Time efficiency for reaching key 0 from mymap: %d nanoseconds",timeEfficiencies.gettingFirstPushedElementToMapEfficiency)
	log.Printf("\n Time efficiency for reaching key 99999 from mymap: %d nanoseconds",timeEfficiencies.gettingLastPushedElementToMapEfficiency)


}

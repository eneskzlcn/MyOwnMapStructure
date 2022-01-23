### Linked List
Linked lists are dynamic data structures consist of nodes. Each node keeps
the address of node that came after it. The linked list structure implemented
in my project is works in this way.

### Hash Function

Hash functions are used in different purposes. In hash table structures, hash function
uses for the purpose of generating index for given key. What the indexes do will mention
in Map.

### Map Structures

Map structures contains data in format of key value pair and provides a constant
efficiency for accessing an element (generally). For the purpose of using maps is,
provide a constant efficiency independent of data that map contains.

As example of that, suppose that you want to reach value of element with key 'IMPORTANT'
in a map that contains almost 10 key value pairs like that and, you access the value
in 0.1 seconds. Suppose that the size of map increased and map contains 1000000 key
value pairs now. And you still access the value in 0.1 seconds.

### My Map Structure

MyMap structure that I implemented in this project is consist of a hash function,
an LinkedList pointer array with constant size. Hash function takes the ascii value
of given keys string then takes the mod size of it to get an index for the key. Then
it is get pushed to the LinkedList that stands on the generated index of the 
LinkedList array. This is how we almost access all the element in O(1) efficiency.

### Result

The image below shows the efficiencies of normal array and my map on access elements
which are pushed first and pushed last to the structures.

![Not Loaded Image](https://eneskzlcn.github.io/MyOwnMapStructure/result.png)
## Problem Statement
Millions of people and cargo are transferred around the world through over 100,000 flights per day, which are operated by various carriers/agencies. The objective is to track a particular person's flight path by sorting through their flight records. The service must be created as a simple microservice API that can receive a request containing a list of flights, each defined by a source and destination airport code. These flights might not be listed in order and must be sorted to find the total flight paths' starting and ending airports.

## Solution
The solution is to create an API that receives a request containing a list of flights in the format [[source1, destination1],[source2, destination2],...]. The API should respond with a list containing the sorted source and destination airports in the order that forms the complete flight path from the starting airport to the final destination airport.
### Other failed attempted methods
#### psudocode
```
Double LinkedList


LinkedNode {
Data string
Next *LinkedNode
Prev *LinkedNode
}

LinkedList {
nodesLocation mapOfNodes
Head *Node
Tail *Node
}


## populate the map and nodes
For each datapoint
	check if nodes not already defined
	if defined
		# bind them together
		
	if not
		create newer ones and bind them together
		set the head to the source
		tail to destination


```
#### issues
Can't decide the head and tail that much even tho the wiring will make total sense, meaning every time I encounter node will find it in the map and chain it's next node and so on so force but in order to traverse the list I need the head and with this solution it's hard to get it
### Required JSON structure
Input: [[source1, destination1],[source2, destination2],...]<br>
Output: [starting airport, final destination airport]

## API Endpoint
The microservice listens on port 8080 and exposes the flight path tracker under the /calculate endpoint.

## Implementation
To sort the flight paths, the solution involves creating two maps of nodes, one for the destination and an inverted one. The head of the list would be the node that doesn't occur in the inverted map, and the tail would be the node that doesn't occur in the destination when looking through the inverted map.


## API Specification
The API specification can be found in  [api.yml](api%2Fapi.yml). You can use the [swagger online editor](https://editor.swagger.io/) to view it.

## Getting Started
The application can be run manually by installing Golang and executing the following commands:

```
$ make generate
$ make build
$ ./bin/app
```
Alternatively, the application can be run using Docker by executing:
```
$ docker build -t mycoolapp .
$ docker run  -dp 8080:8080 mycoolapp
```

## Time and Space Complexity
Assuming that there are N number of destinations or sources, the **time complexity** of the algorithm is **O(N)**, as the successful logic will need to traverse all N destinations and repeat the process for the source. **The space complexity** of the algorithm is **O(N^2)** because it stores all the destination mapping as well as the inverted version of the map.

## Scalability
To handle the scale of the service, we could use an auto-scaling group to handle a large volume of requests. With one server capable of handling 10,000 requests, we may need 10 servers to handle 100,000 flights per day. Furthermore, caching could be used to handle repeated queries, with LRU eviction policies or other policies based on cache miss rates.

## Concerns
To prevent errors, we could serialize the data before storing it. We could also use something like map-reduce if the route data cannot be stored on a single machine.

## Future Improvements
-[x] Add CI
-[ ] Add CD
-[ ] Add alerts over the API and logs
-[ ] Autoscaling group based on the cpu/memory utilization
-[x] Dockerize the app
-[x] Generate the api from a swagger file instead of keeping 2 different copies between the code and design.
-[ ] Add auth, rate limiters
-[x] Add versioning
-[x] Add testcases for the api, logic
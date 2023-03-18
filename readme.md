## Problem definition
Story: There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

Goal: To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

Required JSON structure:

[["SFO", "EWR"]]                                                                           => ["SFO", "EWR"]

[["ATL", "EWR"], ["SFO", "ATL"]]                                                   => ["SFO", "EWR"]

[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

Specifications:
* Your miscroservice must listen on port 8080 and expose the flight path tracker under the /calculate endpoint.

* Create a private GitHub repo and add https://github.com/taariq as a collaborator to the project. Please only add the collaborators when you are sure you are finished.

* Define and document the format of the API endpoint in the README.



## Thought process
seems like a graph problem, not relaying on any search tried to figure another way to address this problem

### Example 1 (failure)
#### issues
Can't decide the head and tail that much even tho the wiring will make total sense, meaning every time I encounter node will find it in the map and chain it's next node and so on so force but in order to traverse the list I need the head and with this solution it's hard to get it
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
——
HEAD IND
Tail  EWR

MAP{
IND => #1 
EWR => #2
SFO => #3
ATL => #4
GSO => #5
}
Head = GSO
Tail = IND

```



### Example 2 (successful-picked)
Instead of looking through all the nodes, why not creat 2 maps of nodes, one for the destination and the inverted one, 
and to get the head, it will be the node that doesn't occur in the inverted and to get the tail it will be the node that doesn't occur in the destination when looking through the inverted.

the reasoning behind this is that the head would not have any nodes pointing to it, as well as the tail in the inverted map.

## Time complexity
Assume that we have 10 destinations as well as the sources, then the successful logic (example #2) will need to traverse all 10 and since we are doing this process twice for the source and destination, this means 2*N,
  - ### Time complexity: O(N) 
    - where is N is the number of destinations or sources 
  - ### Space complexity: O(N^2)
    - Because we are storing all the destination mapping as well as the inverted version of the map as well
## Scale
- Can we handle 100K flights?
  - Assuming that a single server could handle 10K requests, most likely we will need 10 servers to handle that load, and maybe we could have some sort of auto-scaling group to handle the scale automatically for us
- Would it all fit within the same machine?
  - Assuming that the max point name is 1000 character long around 1KB then 10K requests will need 1GB of memory which is applicable in this case 
- We could also caching since some flights might be repeatly queried
  - we can use LRU eviction policy or others based on some trial with the cache miss rate.
## Concerns 
- Could the route has no end?, Could be there some naming issues? some points have upper and lower?
  - We could serialize the data to prevent such thing 
- What if we can not fit the route data into a single machine?
  - We could use something like map reduce in order to get the result.
## Future improvements
- Add CI/CD
- Add alerts over the API and logs
- Autoscaling group based on the cpu/memory utilization
- Dockerize the app
- Generate the api from a swagger file instead of keeping 2 different copies between the code and design.
- Add auth, rate limiters
- Add versioning

//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* After creating the map, perform the following actions:
//  - call display server info function
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServers(servers map[string]int) {
	fmt.Println("There are", len(servers), "servers")
	// number of servers for each status
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are in maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := make(map[string]int)
	//* Set all of the server statuses to `Online` when creating the map
	for _, name := range servers {
		serverMap[name] = Online
	}

	//  - call display server info function
	printServers(serverMap)

	//  - change server status of `darkstar` to `Retired`
	serverMap["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverMap["aiur"] = Offline
	printServers(serverMap)

	//  - change server status of all servers to `Maintenance`
	for key := range serverMap {
		serverMap[key] = Maintenance
	}

	// bonus
	delete(serverMap, "omicron")
	printServers(serverMap)
}

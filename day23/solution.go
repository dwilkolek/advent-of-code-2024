package day23

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day X, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day X, part 2: %d", solve())
}

var network = map[string][]string{}

func solve() int {
	file, _ := os.Open("day23/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		conn := scanner.Text()
		computers := strings.Split(conn, "-")
		network[computers[0]] = append(network[computers[0]], computers[1])
		network[computers[1]] = append(network[computers[1]], computers[0])
	}
	counter := 0
	mNetworks := [][]string{}
	for c, _ := range network {
		mNetworks = append(mNetworks, buildNetwork(c)...)
	}
	allCliques := map[string][]string{}
	for c, _ := range network {
		for _, clickque := range buildNetwork(c) {
			gid := groupId(clickque)
			if _, ok := allCliques[gid]; !ok {
				allCliques[gid] = clickque
			}
		}
	}

	for _, cliques := range allCliques {
		if hasT(cliques) {
			counter++
		}
	}

	return counter
}
func groupId(computers []string) string {
	copied := make([]string, len(computers))
	copy(copied, computers)
	slices.Sort(copied)
	//logger.Println(strings.Join(copied, "-"))
	return strings.Join(copied, "-")
}

func hasT(computers []string) bool {
	for _, computer := range computers {
		if strings.HasPrefix(computer, "t") {
			return true
		}
	}
	return false
}

func isInNetwork(n []string, c string) bool {
	for _, ni := range n {
		if ni == c {
			return true
		}
	}
	return false
}

//func buildNetwork(networkDepth int, mNetworks [][]string) [][]string {
//	newMNetworks := [][]string{}
//	for _, mNetwork := range mNetworks {
//		for _, c := range network[mNetwork[len(mNetwork)-1]] {
//			if !isInNetwork(mNetwork, c) {
//				newMNetworks = append(newMNetworks, append(mNetwork, c))
//			}
//		}
//	}
//	if networkDepth == 0 {
//
//		return newMNetworks
//	}
//	return buildNetwork(networkDepth-1, newMNetworks)
//}

func find3rd(a, b string) [][]string {
	networks := [][]string{}
	for k, v := range network {
		if slices.Contains(v, b) && slices.Contains(v, a) && k != a && k != b {
			networks = append(networks, []string{
				a, b, k,
			})
		}
	}
	return networks
}
func buildNetwork(computer string) [][]string {
	connected := network[computer]
	groups := [][]string{}
	for _, n1 := range connected {
		groups = append(groups, find3rd(computer, n1)...)
	}
	return groups
}

func expand(computer string, pairs [][]string) [][]string {
	ngGroup := [][]string{}
	for _, c := range network[computer] {
		for _, pair := range pairs {
			ngGroup = append(ngGroup, append(pair, c))
		}
	}
	return ngGroup
}

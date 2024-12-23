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
	parse()
	cliques := findCliques()
	counter := 0
	for _, clique := range cliques {
		if hasT(clique) {
			counter++
		}
	}
	logger.Printf("Day 23, part 1: %d", counter)
}

func Part2() {
	parse()
	cliques := findCliques()

	logger.Printf("Day 23, part 2: %s", extendCliqueAndReturnPassword(cliques))
}

type pair struct {
	a, b string
}

var network = map[string][]string{}
var adjNetwork = map[pair]bool{}

func parse() {
	file, _ := os.Open("day23/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	network = map[string][]string{}
	adjNetwork = map[pair]bool{}
	for scanner.Scan() {
		conn := scanner.Text()
		computers := strings.Split(conn, "-")
		network[computers[0]] = append(network[computers[0]], computers[1])
		network[computers[1]] = append(network[computers[1]], computers[0])
		adjNetwork[pair{computers[0], computers[1]}] = true
		adjNetwork[pair{computers[1], computers[0]}] = true
	}
}

func findCliques() map[string][]string {

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
	return allCliques
}

func extendCliqueAndReturnPassword(cliques map[string][]string) string {
	maxLen := 0
	var biggest []string
	for _, c := range cliques {
		m := findNext(c)
		if maxLen < len(m) {
			maxLen = len(m)
			biggest = m
		}
	}
	return groupId(biggest)
}

func findNext(g []string) []string {
	net := map[string]bool{}
	for _, gi := range g {
		net[gi] = true
	}

	newAddition := true
	for newAddition {
		newAddition = false
		for k, _ := range network {
			if net[k] {
				continue
			}
			connectedToRest := true
			for c2, _ := range net {
				if !adjNetwork[pair{a: k, b: c2}] {
					connectedToRest = false
					break
				}
			}

			if connectedToRest {
				newAddition = true
				net[k] = true
			}
		}
	}

	group := []string{}
	for k, _ := range net {
		group = append(group, k)
	}

	return group
}

func groupId(computers []string) string {
	copied := make([]string, len(computers))
	copy(copied, computers)
	slices.Sort(copied)
	return strings.Join(copied, ",")
}

func hasT(computers []string) bool {
	for _, computer := range computers {
		if strings.HasPrefix(computer, "t") {
			return true
		}
	}
	return false
}

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

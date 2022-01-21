// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

// This package checks that two DFA has intersection or not by DFS algorithm.
package intersection

import (
	"fmt"
	"log"
	"regexinter/dfa"
	"regexinter/nfa"
	"regexinter/runerange"
	"regexp"
)

var nodeMap map[string]*CombineNode

type CombineNode struct {
	Name        string // state1_state2
	Final       bool
	Node1       *dfa.Node
	Node2       *dfa.Node
	Transitions []T
}

type T struct {
	RuneRanges []rune       // rune ranges
	Node       *CombineNode // node
}

func HasIntersection(expr1, expr2 string) bool {
	node1, node2 := convert2Dfa(expr1), convert2Dfa(expr2)
	if node1.Final && node2.Final {
		return true
	}

	nodeMap = make(map[string]*CombineNode)
	firstNode := createNode(node1, node2)
	return dfs(firstNode)
}

func convert2Dfa(expr string) *dfa.Node {
	_, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(fmt.Sprintf("invalid regexp: %q", expr))
	}

	nfaNode, err := nfa.New(expr)
	if err != nil {
		log.Fatal(err)
	}

	return dfa.NewFromNFA(nfaNode)
}

func createNode(node1, node2 *dfa.Node) *CombineNode {
	return &CombineNode{
		Name:  nodeName(node1, node2),
		Node1: node1,
		Node2: node2,
		Final: node1.Final && node2.Final,
	}
}

func dfs(node *CombineNode) bool {
	ranges := findOverlapRanges(node.Node1.Transitions, node.Node2.Transitions)

	for _, r := range ranges {
		nextNode1 := node.Node1.NextState(r)
		nextNode2 := node.Node2.NextState(r)
		node, ok := nodeMap[nodeName(nextNode1, nextNode2)]
		if !ok {
			node = createNode(nextNode1, nextNode2)
			nodeMap[nodeName(nextNode1, nextNode2)] = node
		} else {
			// point to the same node, should ignore
			continue
		}

		// expend new node
		node.Transitions = append(node.Transitions, T{r, node})

		if node.Final {
			return true
		}

		// recursive
		return dfs(node)
	}

	return false
}

func findOverlapRanges(trans1, trans2 []dfa.T) [][]rune {
	result := make([][]rune, 0)

	for _, t1 := range trans1 {
		for _, t2 := range trans2 {
			if runerange.Contains(t1.RuneRanges, t2.RuneRanges) {
				result = append(result, t2.RuneRanges)
			}

			if runerange.Contains(t2.RuneRanges, t1.RuneRanges) {
				result = append(result, t1.RuneRanges)
			}
		}
	}

	return result
}

func nodeName(node1, node2 *dfa.Node) string {
	return fmt.Sprintf("%d_%d", node1.State, node2.State)
}

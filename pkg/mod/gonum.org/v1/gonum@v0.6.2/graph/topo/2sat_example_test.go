// Copyright Β©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package topo_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

var systems = []string{
	// Unsatisfiable system.
	`π₯_a β¨ Β¬π₯_b
Β¬π₯_b β¨ π₯_f
π₯_h β¨ π₯_i
π₯_a β¨ π₯_b
π₯_k β¨ π₯_c
Β¬π₯_f β¨ π₯_h
π₯_c β¨ π₯_g
π₯_f β¨ π₯_g
π₯_h β¨ Β¬π₯_l
Β¬π₯_h β¨ π₯_i
π₯_i β¨ π₯_b
Β¬π₯_i β¨ Β¬π₯_h
π₯_i β¨ Β¬π₯_c
π₯_l β¨ π₯_d
Β¬π₯_j β¨ Β¬π₯_i
Β¬π₯_a β¨ Β¬π₯_j
Β¬π₯_a β¨ π₯_b
Β¬π₯_d β¨ π₯_e
Β¬π₯_k β¨ π₯_h
π₯_l β¨ Β¬π₯_d
π₯_l β¨ π₯_d
π₯_l β¨ Β¬π₯_f
π₯_b β¨ π₯_d
π₯_b β¨ Β¬π₯_g
π₯_d β¨ Β¬π₯_l
Β¬π₯_l β¨ Β¬π₯_k
`,
	// Satisfiable system.
	`π₯_a β¨ Β¬π₯_b
Β¬π₯_b β¨ π₯_f
π₯_h β¨ π₯_i
π₯_a β¨ π₯_b
π₯_k β¨ π₯_c
Β¬π₯_f β¨ π₯_h
π₯_c β¨ π₯_g
π₯_f β¨ π₯_g
π₯_h β¨ Β¬π₯_l
Β¬π₯_h β¨ π₯_i
π₯_i β¨ π₯_b
Β¬π₯_i β¨ π₯_e
π₯_i β¨ Β¬π₯_c
Β¬π₯_g β¨ Β¬π₯_a
π₯_l β¨ π₯_f
Β¬π₯_j β¨ Β¬π₯_i
Β¬π₯_a β¨ Β¬π₯_j
Β¬π₯_a β¨ π₯_b
Β¬π₯_d β¨ π₯_e
π₯_k β¨ Β¬π₯_a
π₯_k β¨ π₯_h
π₯_l β¨ Β¬π₯_d
π₯_l β¨ π₯_e
π₯_l β¨ Β¬π₯_f
π₯_b β¨ π₯_d
π₯_b β¨ Β¬π₯_g
π₯_d β¨ Β¬π₯_l
π₯_l β¨ π₯_e
`,

	`fun β¨ Β¬fun
fun β¨ Β¬Gonum
Gonum β¨ Gonum
`,
}

// twoSat returns whether the system described in the data read from r is
// satisfiable and a set of states that satisfies the system.
// The syntax used by twoSat is "π₯ β¨ π¦" where π₯ and π¦ may be negated by
// leading "Β¬" characters. twoSat uses the implication graph approach to
// system analysis.
func twoSat(r io.Reader) (state map[string]bool, ok bool) {
	g := simple.NewDirectedGraph()

	sc := bufio.NewScanner(r)
	nodes := make(map[string]node)
	for count := 1; sc.Scan(); count++ {
		line := sc.Text()
		fields := strings.Split(line, "β¨")
		if len(fields) != 2 {
			log.Fatalf("failed to parse on line %d %q: invalid syntax", count, line)
		}
		var variables [2]node
		for i, f := range fields {
			f = strings.TrimSpace(f)
			var negate bool
			for strings.Index(f, "Β¬") == 0 {
				f = strings.TrimPrefix(f, "Β¬")
				negate = !negate
			}
			n, ok := nodes[f]
			if !ok {
				n = node{
					id:   int64(len(nodes) + 1), // id must not be zero.
					name: f,
				}
				nodes[f] = n
			}
			if negate {
				n = n.negated()
			}
			variables[i] = n
		}

		// Check for tautology.
		if variables[0].negated().ID() == variables[1].ID() {
			for _, v := range variables {
				if g.Node(v.ID()) == nil {
					g.AddNode(v)
				}
			}
			continue
		}

		// Add implications to the graph.
		g.SetEdge(simple.Edge{F: variables[0].negated(), T: variables[1]})
		g.SetEdge(simple.Edge{F: variables[1].negated(), T: variables[0]})
	}

	// Find implication inconsistencies.
	sccs := topo.TarjanSCC(g)
	for _, c := range sccs {
		set := make(map[int64]struct{})
		for _, n := range c {
			id := n.ID()
			if _, ok := set[-id]; ok {
				return nil, false
			}
			set[id] = struct{}{}
		}
	}

	// Assign states.
	state = make(map[string]bool)
unknown:
	for _, c := range sccs {
		for _, n := range c {
			if _, known := state[n.(node).name]; known {
				continue unknown
			}
		}
		for _, n := range c {
			n := n.(node)
			state[n.name] = n.id > 0
		}
	}

	return state, true
}

type node struct {
	id   int64
	name string
}

func (n node) ID() int64     { return n.id }
func (n node) negated() node { return node{-n.id, n.name} }

func ExampleTarjanSCC_2sat() {
	for i, s := range systems {
		state, ok := twoSat(strings.NewReader(s))
		if !ok {
			fmt.Printf("system %d is not satisfiable\n", i)
			continue
		}
		var ps []string
		for v, t := range state {
			ps = append(ps, fmt.Sprintf("%s:%t", v, t))
		}
		sort.Strings(ps)
		fmt.Printf("system %d is satisfiable: %s\n", i, strings.Join(ps, " "))
	}

	// Output:
	// system 0 is not satisfiable
	// system 1 is satisfiable: π₯_a:true π₯_b:true π₯_c:true π₯_d:true π₯_e:true π₯_f:true π₯_g:false π₯_h:true π₯_i:true π₯_j:false π₯_k:true π₯_l:true
	// system 2 is satisfiable: Gonum:true fun:true
}

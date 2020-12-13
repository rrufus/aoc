package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	units      = regexp.MustCompile(`(\d+) units each`)
	hp         = regexp.MustCompile(`(?P<HP>\d+) hit points`)
	damage     = regexp.MustCompile(`(?P<Damage>\d+) (?P<DamageType>\S+) damage`)
	initiative = regexp.MustCompile(`initiative (?P<Initiative>\d+)`)
	weakTo     = regexp.MustCompile(`weak to (?P<Types>[\S\s]+)`)
	immuneTo   = regexp.MustCompile(`immune to (?P<Types>[\S\s]+)`)
)

const (
	ImmuneSystem = Army(iota)
	Infection
)

type Army int

type Group struct {
	Allegiance Army
	Number     int
	Units      int
	Health     int
	Damage     int
	DamageType string
	Initiative int
	WeakTo     []string
	ImmuneTo   []string
}

func (g *Group) EffectivePower() int {
	return g.Units * g.Damage
}

func (g *Group) SelectTarget(groups []*Group, selected map[*Group]bool) *Group {
	highestDamage := 0
	var target *Group
	for _, newDefendingGroup := range groups {
		if newDefendingGroup.Allegiance == g.Allegiance || selected[newDefendingGroup] || newDefendingGroup.Units <= 0 {
			continue
		}
		damageEstimate := g.EstimateDamage(newDefendingGroup, true)
		if damageEstimate > highestDamage {
			target = newDefendingGroup
			highestDamage = damageEstimate
		} else if damageEstimate == highestDamage && target != nil {
			newEffectivePower := newDefendingGroup.EffectivePower()
			targetEffectivePower := target.EffectivePower()
			if newEffectivePower > targetEffectivePower {
				target = newDefendingGroup
			} else if targetEffectivePower == newEffectivePower && newDefendingGroup.Initiative > target.Initiative {
				target = newDefendingGroup
			}
		}
	}
	if target != nil {
		selected[target] = true
	}
	return target
}

func (g *Group) EstimateDamage(enemy *Group, print bool) int {
	damageMultiplier := 1
	if strings.Contains(strings.Join(enemy.ImmuneTo, " "), g.DamageType) {
		return 0
	}
	if strings.Contains(strings.Join(enemy.WeakTo, " "), g.DamageType) {
		damageMultiplier = 2
	}
	damage := g.EffectivePower() * damageMultiplier

	// if print {
	// 	if g.Allegiance == ImmuneSystem {
	// 		fmt.Printf("Immune System group %v would deal defending group %v %v damage\n", g.Number, enemy.Number, damage)
	// 	} else {
	// 		fmt.Printf("Infection group %v would deal defending group %v %v damage\n", g.Number, enemy.Number, damage)
	// 	}
	// }
	return damage
}

type ByMoveOrder []*Group

func (a ByMoveOrder) Len() int { return len(a) }
func (a ByMoveOrder) Less(i, j int) bool {
	groupIEffectivePower := a[i].EffectivePower()
	groupJEffectivePower := a[j].EffectivePower()
	if groupIEffectivePower == groupJEffectivePower {
		return a[i].Initiative > a[j].Initiative
	}
	return groupIEffectivePower > groupJEffectivePower
}
func (a ByMoveOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByInitiative []*Group

func (a ByInitiative) Len() int           { return len(a) }
func (a ByInitiative) Less(i, j int) bool { return a[i].Initiative > a[j].Initiative }
func (a ByInitiative) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	groups := Parse(ReadFromInput())

	fmt.Println("Part 1")
	for CountGroupsInArmy(groups, ImmuneSystem) > 0 && CountGroupsInArmy(groups, Infection) > 0 {
		sort.Sort(ByMoveOrder(groups))
		// select targets
		// fmt.Println("")
		targets := map[*Group]*Group{}
		selected := map[*Group]bool{}
		for _, group := range groups {
			if group.Units > 0 {
				target := group.SelectTarget(groups, selected)
				if target != nil {
					targets[group] = target
				}
			}
		}
		// fmt.Println("")
		// attack!
		sort.Sort(ByInitiative(groups))
		for _, attackingGroup := range groups {
			if defendingGroup, exists := targets[attackingGroup]; exists && attackingGroup.Units > 0 {
				damage := attackingGroup.EstimateDamage(defendingGroup, false)
				unitsKilled := damage / defendingGroup.Health
				defendingGroup.Units -= unitsKilled
				// if attackingGroup.Allegiance == ImmuneSystem {
				// 	fmt.Printf("Immune System group %v attacks defending group %v, killing %v units\n", attackingGroup.Number, defendingGroup.Number, unitsKilled)
				// } else {
				// 	fmt.Printf("Infection group %v attacks defending group %v, killing %v units\n", attackingGroup.Number, defendingGroup.Number, unitsKilled)
				// }
			}
		}

		// fmt.Printf("\n\n")
	}
	fmt.Println(CountUnits(groups))

	fmt.Println("Part 2")

}

func CountUnits(groups []*Group) (total int) {
	for _, group := range groups {
		if group.Units > 0 {
			total += group.Units
		}
	}
	return
}

func CountGroupsInArmy(groups []*Group, army Army) (count int) {
	// if army == ImmuneSystem {
	// 	fmt.Println("Immune System:")
	// } else {
	// 	fmt.Println("Infection:")
	// }
	for _, group := range groups {
		if group.Allegiance == army && group.Units > 0 {
			count++
			// fmt.Printf("Group %v contains %v units\n", group.Number, group.Units)
		}
	}

	return
}

func Parse(lines []string) []*Group {
	groups := []*Group{}
	allegiance := ImmuneSystem
	n := 1

	for _, line := range lines {
		if line == "" || line == "Immune System:" {
			continue
		}
		if line == "Infection:" {
			allegiance = Infection
			n = 1
			continue
		}
		groups = append(groups, ParseGroup(line, allegiance, n))
		n++
	}

	return groups
}

func ParseGroup(in string, allegiance Army, number int) *Group {
	group := &Group{Allegiance: allegiance, Number: number}

	group.Units, _ = strconv.Atoi(GetSubMatch(units, in)[0])
	group.Health, _ = strconv.Atoi(GetSubMatch(hp, in)[0])
	group.Initiative, _ = strconv.Atoi(GetSubMatch(initiative, in)[0])
	DamageSubMatches := GetSubMatch(damage, in)
	group.Damage, _ = strconv.Atoi(DamageSubMatches[0])
	group.DamageType = DamageSubMatches[1]

	if strings.Contains(in, "(") {
		bracketsGroup := strings.Split(in, "(")[1]
		bracketsGroup = strings.Split(bracketsGroup, ")")[0]

		parts := strings.Split(bracketsGroup, ";")
		for _, part := range parts {
			if weakTo.MatchString(part) {
				group.WeakTo = strings.Split(GetSubMatch(weakTo, part)[0], ", ")
			}
			if immuneTo.MatchString(part) {
				group.ImmuneTo = strings.Split(GetSubMatch(immuneTo, part)[0], ", ")

			}
		}
	}

	return group
}

func GetSubMatch(r *regexp.Regexp, input string) []string {
	res := r.FindStringSubmatch(input)
	return res[1:]
}

func ReadFromStdIn() []string {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)

read_loop:
	for {
		text, _ := reader.ReadString('\n')
		if text == "go\n" {
			break read_loop
		}
		lines = append(lines, strings.TrimSpace(text))
	}

	return lines
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}
	return ints
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

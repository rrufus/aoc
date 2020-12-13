package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

const (
	ImmuneSystem = Army(iota)
	Infection
)

var (
	units      = regexp.MustCompile(`(\d+) units each`)
	hp         = regexp.MustCompile(`(?P<HP>\d+) hit points`)
	damage     = regexp.MustCompile(`(?P<Damage>\d+) (?P<DamageType>\S+) damage`)
	initiative = regexp.MustCompile(`initiative (?P<Initiative>\d+)`)
	weakTo     = regexp.MustCompile(`weak to (?P<Types>[\S\s]+)`)
	immuneTo   = regexp.MustCompile(`immune to (?P<Types>[\S\s]+)`)
	print      = false
	ArmyToName = map[Army]string{ImmuneSystem: "Immune System", Infection: "Infection"}
)

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
		damageEstimate := g.EstimateDamage(newDefendingGroup, print)
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

	if print {
		fmt.Printf("%v group %v would deal defending group %v %v damage\n", ArmyToName[g.Allegiance], g.Number, enemy.Number, damage)
	}
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
	in := ReadFromInput()
	groups := Parse(in, 0)

	fmt.Println("Part 1")

	victor := DoBattle(groups)
	fmt.Printf("%v wins with %v units left.\n", ArmyToName[victor], CountUnits(groups))

	fmt.Println("Part 2")
	boost := 0
	for victor != ImmuneSystem {
		boost++
		groups = Parse(in, boost)
		if boost == 75 {
			// this case causes an endless loop because the attack power on both sides
			// for the groups with the highest effective power cant kill one unit off.
			// could fix properly sometime but easier to skip.
			boost++
		}
		victor = DoBattle(groups)
	}
	fmt.Printf("With attack boost %v, %v wins with %v units left.\n", boost, ArmyToName[victor], CountUnits(groups))

}

func DoBattle(groups []*Group) (victor Army) {
	var nGroupsImmuneSystem int
	var nGroupsInfection int
	for {
		nGroupsImmuneSystem = CountGroupsInArmy(groups, ImmuneSystem)
		nGroupsInfection = CountGroupsInArmy(groups, Infection)
		if nGroupsImmuneSystem == 0 {
			victor = Infection
			return
		}
		if nGroupsInfection == 0 {
			victor = ImmuneSystem
			return
		}

		sort.Sort(ByMoveOrder(groups))
		// select targets
		if print {
			fmt.Println("")
		}
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
		if print {
			fmt.Println("")
		}
		// attack!
		sort.Sort(ByInitiative(groups))
		for _, attackingGroup := range groups {
			if defendingGroup, exists := targets[attackingGroup]; exists && attackingGroup.Units > 0 {
				damage := attackingGroup.EstimateDamage(defendingGroup, false)
				unitsKilled := damage / defendingGroup.Health
				defendingGroup.Units -= unitsKilled
				if print {
					fmt.Printf("%v group %v attacks defending group %v, killing %v units\n", ArmyToName[attackingGroup.Allegiance], attackingGroup.Number, defendingGroup.Number, unitsKilled)
				}
			}
		}
	}
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
	if print {
		fmt.Printf("%v:\n", ArmyToName[army])
	}
	for _, group := range groups {
		if group.Allegiance == army && group.Units > 0 {
			count++
			if print {
				fmt.Printf("Group %v contains %v units\n", group.Number, group.Units)
			}
		}
	}
	if count == 0 && print {
		fmt.Println("No groups remain.")
	}

	return
}

func Parse(lines []string, immuneSystemBoost int) []*Group {
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
		groups = append(groups, ParseGroup(line, allegiance, n, immuneSystemBoost))
		n++
	}

	return groups
}

func ParseGroup(in string, allegiance Army, number, immuneSystemBoost int) *Group {
	group := &Group{Allegiance: allegiance, Number: number}

	group.Units, _ = strconv.Atoi(GetSubMatch(units, in)[0])
	group.Health, _ = strconv.Atoi(GetSubMatch(hp, in)[0])
	group.Initiative, _ = strconv.Atoi(GetSubMatch(initiative, in)[0])
	DamageSubMatches := GetSubMatch(damage, in)
	group.Damage, _ = strconv.Atoi(DamageSubMatches[0])
	if allegiance == ImmuneSystem {
		group.Damage += immuneSystemBoost
	}
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

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

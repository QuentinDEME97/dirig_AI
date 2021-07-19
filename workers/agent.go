package worker

import (
	"agents/world"
	"math/rand"
	"time"
	"strconv"
)

var names = [5]string{"Robert", "Henry", "Louis", "Philippe", "François"}
var qualify = [5]string{"Le Noble", "Le Sanglant", "Le Juste", "Le Bon", "Le Riche"}
var names_attributed = make(map[string]int)

type Ruler struct {
	Name		string
	City		world.City
}

func (r Ruler) GetCity() world.City {
	return r.City
}

func GenerateRuler(cities []world.City) []Ruler {
	n := len(cities)
	name := ""
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := make([]Ruler, n)
	for i := 0; i < n; i++ {
		r_name := names[gen.Intn(len(names))]
		r_qualify := qualify[gen.Intn(len(qualify))]
		number_of := incrementWithCounter(names_attributed, r_name)
		if number_of > 1 {
			name = r_name + " " + strconv.Itoa(number_of) + " " + r_qualify
		} else {
			name = r_name + " " + r_qualify
		}
		r := Ruler{
			Name: name,
			City: cities[i],
		}
		t[i] = r
	}
	return t 
}

func (r Ruler) SetCity(c world.City) {
	r.City = c
}

func find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func incrementWithCounter(m map[string]int, name string) int {
	if _, ok := m[name]; ok {
		//do something here
		m[name] += 1
		return m[name]
	}
	m[name] = 1
	return 1
}
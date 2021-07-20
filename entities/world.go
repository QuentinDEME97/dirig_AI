package world

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

var RessourcesCreated = 0
var AllRessources []Ressource

var CityCreated = 0
var AllCities []City

var names = [5]string{"Château", "Champs", "Fort"}
var qualify = [5]string{"Sanglant", "Brillant", "Fort", "Fertiles", "De Bataille"}
var names_attributed = make(map[string]int)

type Entity interface {
	Display()
}

type node struct{
	X		float32
	Y		float32
}


type Ressource struct {
	node
	Name	string
	Id		int
}

type City struct {
	node
	Name	string
	Id		int
}

type Map struct {
	X		int
	Y 		int
	Grid	[][] Entity
}


func (c City) GetNode() node {
	return c.node
}

func (c City) String() string {
	return fmt.Sprintf("%v (%v, %v)", c.Name, c.node.X, c.node.Y)
}

func NewRessource(x float32, y float32, name string) *Ressource {
    RessourcesCreated += 1
	r := Ressource{
		node: node{
			X: x,
			Y: y,
		},
		Name: name,
		Id: RessourcesCreated,
	}
	AllRessources = append(AllRessources, r)
	return &r
}

func NewCity(x float32, y float32, name string) *City {
	name = ""
	// gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	CityCreated += 1
	r_name := names[rand.Intn(len(names))]
	r_qualify := qualify[rand.Intn(len(qualify))]
	name = r_name + " " + r_qualify
	number_of := incrementWithCounter(names_attributed, name)
	if number_of > 1 {
		name += " " + strconv.Itoa(number_of)
	}
	r := City{
		node: node{
			X: x,
			Y: y,
		},
		Name: name,
		Id: CityCreated,
	}
	AllCities = append(AllCities, r)
	return &r
}

// func CreateRessource(x float32, y float32, name string) Ressource {
// 	RessourcesCreated += 1
// 	r := Ressource{
// 		node: node{
// 			X: x,
// 			Y: y,
// 		},
// 		Name: name,
// 		Id: RessourcesCreated,
// 	}
// 	AllRessources = append(AllRessources, r)
// 	return r
// }

// func CreateCity(x float32, y float32, name string) City {
// 	CityCreated += 1
// 	r := City{
// 		node: node{
// 			X: x,
// 			Y: y,
// 		},
// 		Name: name,
// 		Id: CityCreated,
// 	}
// 	AllCities = append(AllCities, r)
// 	return r
// }

func (n node) Display() {
	fmt.Printf("(%#v, %#v)\n", n.X, n.Y)
}

func (c City) Display() {
	fmt.Printf("City (%#v, %#v)\n", c.X, c.Y)
}

func (r Ressource) Display() {
	fmt.Printf("%#v is at (%#v, %#v)\n", r.Name, r.X, r.Y)
}

// func (n Node) Display() {
// 	fmt.Println("Coucou")
// }
// Following relative to Map

// func CreateMap(x int, y int, nbCity int, nbRessources int) Map {
// 	grid := make([]Entity, x*y)
// 	index := 0
// 	for i := 0; i < y; i++ {
// 		for j := 0; j < x; j++ {
// 			City_place := rand.Intn(100)
// 			nb := rand.Intn(100)
// 			if City_place <= 1 && nbCity != 0 {
// 				grid[index] = City{
// 					node: node{
// 						X: float32(j),
// 						Y: float32(i),
// 					},
// 					Name: "City",
// 					Id: nbCity,
// 				}
// 				nbCity -= 1
// 			} else {
// 				if nb <= 5 {
// 					grid[index] = CreateRessource(float32(j), float32(i), "Wood")
// 					nbRessources -= 1
// 				} else {
// 					grid[index] = node {
// 						X: float32(j),
// 						Y: float32(i),
// 					}
// 				}
// 			}
// 			index += 1
// 		}
// 	}
// 	return Map {
// 		X: x,
// 		Y: y, 
// 		Grid: grid,
// 	}
// }

func getPoints(n int, x_dim int, y_dim int) []node {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	points := make([]node, n)
	for i := 0; i < n; i++ {
		x := r.Intn(x_dim)
		y := r.Intn(y_dim)
		n := node {
			X: float32(x),
			Y: float32(y),
		}
		_, found := find(points, n)
		if !found {
			points[i] = n
		} else {
			i--
		}
	}
	return points
}

func CreateMapV2(x int, y int, nbCity int, nbRessources int) Map {
	cities := getPoints(nbCity, x, y)
	Ressources := getPoints(nbRessources, x, y)
	grid := make([][]Entity, y)
	index := 0
	for i := 0; i < y; i++ {
		grid[i] = make([]Entity, x)
		for j := 0; j < x; j++ {
			var created = false

			for res_i, res := range(Ressources) {
				if int(res.X) == j && int(res.Y) == i {
					grid[i][j] = NewRessource(float32(j), float32(i), "Wood")
					// fmt.Printf("Ressources - We remove %v \nfrom %v\n", res, Ressources)
					Ressources = remove(Ressources, res_i)
					// fmt.Printf("%v\n", Ressources)
					created = true
				}
			}

			for City_i, City := range(cities) {
				if int(City.X) == j && int(City.Y) == i {
					grid[i][j] = NewCity(float32(j), float32(i), "City")
					// fmt.Printf("City - We remove %v \nfrom %v\n", City, cities)
					cities = remove(cities, City_i)
					// fmt.Printf("%v\n", cities)
					created = true
				}
			}

			if created == false {
				grid[i][j] = node {
					X: float32(j),
					Y: float32(i),
				}
			}
			created = false
			index += 1
		}
	}

	return Map {
		X: x,
		Y: y, 
		Grid: grid,
	}
}

func remove(s []node, i int) []node {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func find(slice []node, val node) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func (m Map) Display() {
	index := 0
	for i := 0; i < m.X; i++ {
		for j := 0; j < m.Y; j++ {
			eType := fmt.Sprintf("%T", m.Grid[j][i])
			switch eType {
				case "world.node":
				fmt.Printf("_ ")
				case "*world.Ressource":
					fmt.Printf("x ")
				case "*world.City":
					fmt.Printf("¤ ")
			}
			index += 1
		}
		fmt.Printf("\n")
	}
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

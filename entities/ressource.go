package world

var RessourcesCreated = 0
var AllRessources []Ressource

type Ressource struct {
	node
	Name	string
	value	int
	Id		int
}

func Ressource(name string, value int) *Ressource {
	RessourcesCreated += 1
	r := Ressource{
		node: node{
			X: x,
			Y: y,
		},
		Name: name,
		Value: value
		Id: RessourcesCreated,
	}
	AllRessources = append(AllRessources, r)
	return &r
}
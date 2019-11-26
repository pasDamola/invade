package invade

import(
	
)

type Aliens struct{
	Name, City []string
	Moves int
}

type Alien map[string]*Aliens

func (data *Aliens) AddAlien(name, cityName string)  {
	data.Name = append(data.Name, name)
	data.City = append(data.City, cityName)
	data.Moves = 0
}

func Get(name string){
	
}

func GetAll() []*Aliens{
	aliens := make([]*Aliens, 0)
	return aliens
}




func(a Alien) Generate()  {
	
}
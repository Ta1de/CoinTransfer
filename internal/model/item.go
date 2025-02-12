package model

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var AvailableItems = map[string]Item{
	"t-shirt":    {ID: 1, Name: "t-shirt", Price: 80},
	"cup":        {ID: 2, Name: "cup", Price: 20},
	"book":       {ID: 3, Name: "book", Price: 50},
	"pen":        {ID: 4, Name: "pen", Price: 10},
	"powerbank":  {ID: 5, Name: "powerbank", Price: 200},
	"hoody":      {ID: 6, Name: "hoody", Price: 300},
	"umbrella":   {ID: 7, Name: "umbrella", Price: 200},
	"socks":      {ID: 8, Name: "socks", Price: 10},
	"wallet":     {ID: 9, Name: "wallet", Price: 50},
	"pink-hoody": {ID: 10, Name: "pink-hoody", Price: 500},
}

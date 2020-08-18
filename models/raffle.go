package models

type Draw struct {
	ID                  string   `json:"id" bson:"id"`
	Concourse           int      `json:"concourse" bson:"concourse"`
	Numbers             []int    `json:"numbers" bson:"numbers"`
	Collection          float64  `json:"collection" bson:"collection"`
	Winners             []Winner `json:"winners" bson:"winners"`
	Accumulated         bool     `json:"accumulated" bson:"accumulated"`
	AccumulatedValue    float64  `json:"accumulatedValue" bson:"accumulatedValue"`
	EstimatedPrize      float64  `json:"estimatedPrize" bson:"estimatedPrize"`
	AccumulatedLastDraw float64  `json:"accumulatedLastDraw" bson:"accumulatedLastDraw"`
}

type Winner struct {
	Number int     `json:"number" bson:"number"`
	Count  int     `json:"count" bson:"count"`
	Amount float64 `json:"amount" bson:"amount"`
	Cities []City  `json:"cities" bson:"cities"`
}

type City struct {
	Name string `json:"name" bson:"name"`
	UF   string `json:"uf" bson:"uf"`
}

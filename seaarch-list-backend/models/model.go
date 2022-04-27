package models

type Movie struct { //create a model for database
	ID   uint   `json:"id" gorm:"primary_key"` // with tags we can specify json name
	Name string `json:"name"`
}

type CreateMovie struct {
	Name string `json:"name"`
}

type UpdateMovie struct { // we make this not compulsary to give
	Name string `json:"name"`
}

type SendMovie struct { // we make this not compulsary to give
	Name string `json:"name"`
}

var Records []SendMovie

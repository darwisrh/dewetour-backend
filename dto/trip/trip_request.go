package tripdto

type TripRequest struct {
	Title        string `json:"title" form:"title" validate:"required"`
	CountryId    int    `json:"countr_id" form:"country_id"`
	Accomodation string `json:"accomodation" form:"accomodation" `
	Transport    string `json:"transport" form:"transport" `
	Eat          string `json:"eat" form:"eat" `
	Day          int    `json:"day" form:"day" `
	Night        int    `json:"night" form:"night" `
	Date         string `json:"date" form:"date" `
	Price        int    `json:"price" form:"price" `
	Kuota        int    `json:"kuota" form:"kuota" `
	QtyCounter   int    `json:"qtycounter" form:"qtycounter"`
	Description  string `json:"description" form:"description" `
	Image        string `json:"image" form:"image"`
}

type TripUpdateRequest struct {
	Title        string `json:"title" form:"title"`
	CountryId    int    `json:"country_id" form:"country_id"`
	Accomodation string `json:"accomodation" form:"accomodation"`
	Transport    string `json:"transport" form:"transport"`
	Eat          string `json:"eat" form:"eat"`
	Day          int    `json:"day" form:"day"`
	Night        int    `json:"night" form:"night"`
	Date         string `json:"date" form:"date"`
	Price        int    `json:"price" form:"price"`
	Kuota        int    `json:"kuota" form:"kuota"`
	QtyCounter   int    `json:"qtycounter" form:"qtycounter"`
	Description  string `json:"description" form:"description"`
	Image        string `json:"image" form:"image"`
}

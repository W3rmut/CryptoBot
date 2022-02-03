package data

type ResponseRUB struct {
	Record Record `xml:"Record"`
}

type Record struct {
	Nominal int    `xml:"Nominal"`
	Value   string `xml:"Value"`
}

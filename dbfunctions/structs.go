package functions

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}
type Objects struct {
	Name   string
	Age    int
	Degree string
	Skill  Skills
}
type Skills struct {
	Primary   []string
	Secondary []string
	Add       Address
}
type Address struct {
	Street   string
	Door_num int
	District string
}
type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}

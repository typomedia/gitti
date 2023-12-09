package hdlr

type Refs struct {
	Branches []Branch
}

type Branch struct {
	Name     string
	Revision string
}

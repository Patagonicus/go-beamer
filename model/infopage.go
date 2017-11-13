package model

type Infopage struct {
	id          uint
	Description string
}

func (p Infopage) ID() uint {
	return p.id
}

type infopage struct {
	id          uint
	description string
}

func (p *infopage) GetPage() Page {
	return Infopage{p.id, p.description}
}

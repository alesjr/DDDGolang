package core_domain_model

type FederationUnit int
type Color int
type Driver string

const (
	RO FederationUnit = 11
	AC FederationUnit = 12
	AM FederationUnit = 13
	RR FederationUnit = 14
	PA FederationUnit = 15
	AP FederationUnit = 16
	TO FederationUnit = 17
	MA FederationUnit = 21
	PI FederationUnit = 22
	CE FederationUnit = 23
	RN FederationUnit = 24
	PB FederationUnit = 25
	PE FederationUnit = 26
	AL FederationUnit = 27
	SE FederationUnit = 28
	BA FederationUnit = 29
	MG FederationUnit = 31
	ES FederationUnit = 32
	RJ FederationUnit = 33
	SP FederationUnit = 35
	PR FederationUnit = 41
	SC FederationUnit = 42
	RS FederationUnit = 43
	MS FederationUnit = 50
	MT FederationUnit = 51
	GO FederationUnit = 52
	DF FederationUnit = 53
)

const (
	Branca Color = iota + 1
	Preta
	Parda
	Amarela
	Indígena
)

const (
	Mongo = "mongo"
	Postgres = "pgx"
)
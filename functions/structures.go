package functions

type Farm struct {
	Ants int
	Start string
	End string
	Rooms []string
	Links []string
	AdjList map[string][]string
}

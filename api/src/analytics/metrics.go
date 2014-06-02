package analytics

type Metric struct {
	Id string
	Name string
	ClientId string
	query QueryDefinition
}
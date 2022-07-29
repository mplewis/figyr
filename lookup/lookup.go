package lookup

type getter interface {
	Get(string) (string, bool)
}

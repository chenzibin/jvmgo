package classpath

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
}

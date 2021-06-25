package resources

type IResource interface {
	GetStatus() int
	GetData() map[string]interface{}
}

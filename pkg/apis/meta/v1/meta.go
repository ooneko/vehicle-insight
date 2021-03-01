package v1

type ObjectMetaAccessor interface {
	GetObjectMeta() Object
}

type Object interface {
	GetName() string
	SetName(name string)
	GetResourceVersion() string
	SetResourceVersion(version string)
}

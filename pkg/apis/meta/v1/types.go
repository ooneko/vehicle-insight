package v1

type ObjectMeta struct {
	Name            string
	ResourceVersion string
}

func (meta *ObjectMeta) GetName() string                   { return meta.Name }
func (meta *ObjectMeta) SetName(name string)               { meta.Name = name }
func (meta *ObjectMeta) GetResourceVersion() string        { return meta.ResourceVersion }
func (meta *ObjectMeta) SetResourceVersion(version string) { meta.ResourceVersion = version }

func (meta *ObjectMeta) GetObjectMeta() Object { return meta }

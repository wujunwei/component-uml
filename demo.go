package component_uml

type Component struct {
	meta Meta
	spec ComponentSpec
}

//Core 表示k8s内置资源
type Core struct {
	group   string
	version string
	kind    string
}

const (
	StatefulSet = "statefulset"
	Deployment  = "deployment"
	CloneSet    = "cloneset"
)

//Workload 表示具体承载组件的工作负载的类型，可能的值： Deployment StatefulSet CloneSet 等等只要实现对应的接口即可
type Workload interface {
	Name() string
	Replica() *int32
	Healthy() bool
	Core() Core
}

type Ability struct {
}

//ComponentDefine 表示一种组件类型的定义用以支持多种部署模式的组件 todo:后期需要修改接口名
type ComponentDefine interface {
	Name() string
	Description() string
	WorkLoad() Workload
	Install() error
	Template() string
	Parameters() map[string]interface{}
}

type ComponentSpec struct {
	ComponentDefine
	trait []Trait
}

type Meta struct {
	Name        string
	namespace   string
	Annotations map[string]string
}
type Trait interface {
}

type Application struct {
	meta  Meta
	scope string //todo
	c     []Component
}

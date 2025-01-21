package environment

type Environment struct {
    Variables map[string]interface{}
}

func NewEnvironment() *Environment {
    return &Environment{
        Variables: make(map[string]interface{}),
    }
}

func (env *Environment) Set(name string, value interface{}) {
    env.Variables[name] = value
}

func (env *Environment) Get(name string) interface{} {
    return env.Variables[name]
}
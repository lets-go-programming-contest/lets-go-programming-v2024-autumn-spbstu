package condition

import ( 
    "github.com/zafod42/task-2-1/errors/operation"
    "github.com/zafod42/task-2-1/errors/temperature"
)

type Condition struct {
    MinT int
    MaxT int
}

func (c *Condition) Init() {
    c.MinT = 15
    c.MaxT = 30 
}

func (c *Condition) GetOptimal() int {
    var optimal int
    if c.MaxT >= c.MinT {
        optimal = c.MinT
    } else {
        optimal = -1
    }
    return optimal
}

func (c *Condition) Set(sign string, temp int) error {
    if temp < 15 || temp > 30 {
        return temperatureError.TemperatureError{}
    }
    switch sign {
    case ">=":
        c.setMin(temp) 
    case "<=":
        c.setMax(temp)
    default:
        return  operationError.OperationError{}
    }
    return nil
}

func (c *Condition) setMin(temp int) {
    if c.MinT < temp  {
        c.MinT = temp 
    }
}

func (c *Condition) setMax(temp int) {
    if c.MaxT > temp {
        c.MaxT = temp
    }
}

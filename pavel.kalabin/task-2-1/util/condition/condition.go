package condition

import ( 
    "fmt"
    "os"
)

type Condition struct {
    MinT int
    MaxT int
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

func (c *Condition) Set(sign string, temp int) {
    switch sign {
    case ">=":
        c.setMin(temp) 
    case "<=":
        c.setMax(temp)
    default:
        fmt.Fprintf(os.Stderr, "Undefined operation\n");
    }
}

func (c *Condition) setMin(temp int) {
    c.MinT = temp 
}

func (c *Condition) setMax(temp int) {
    c.MaxT = temp
}

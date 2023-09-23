package developer

type Developer struct {
    position string
    salary   float64
    address  string
}

func NewDeveloper(position string, salary float64, address string) *Developer {
    return &Developer{
        position: position,
        salary:   salary,
        address:  address,
    }
}

func (d *Developer) GetPosition() string {
    return d.position
}

func (d *Developer) SetPosition(position string) {
    d.position = position
}

func (d *Developer) GetSalary() float64 {
    return d.salary
}

func (d *Developer) SetSalary(salary float64) {
    d.salary = salary
}

func (d *Developer) GetAddress() string {
    return d.address
}

func (d *Developer) SetAddress(address string) {
    d.address = address
}
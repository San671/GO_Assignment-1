package marketing

type Marketing struct {
    position string
    salary   float64
    address  string
}

func NewMarketing(position string, salary float64, address string) *Marketing {
    return &Marketing{
        position: position,
        salary:   salary,
        address:  address,
    }
}

func (m *Marketing) GetPosition() string {
    return m.position
}

func (m *Marketing) SetPosition(position string) {
    m.position = position
}

func (m *Marketing) GetSalary() float64 {
    return m.salary
}

func (m *Marketing) SetSalary(salary float64) {
    m.salary = salary
}

func (m *Marketing) GetAddress() string {
    return m.address
}

func (m *Marketing) SetAddress(address string) {
    m.address = address
}
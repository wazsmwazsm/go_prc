package person

type Person struct {
	firstName, lastName string // 小写开头字段，无法外部导出
}

// getter
func (p *Person) FirstName() string {
	return p.firstName
}

// setter
func (p *Person) SetFirstName(name string) {
	p.firstName = name
}

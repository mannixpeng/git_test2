package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name        string
	Role        string
	Age         int32
	EmployeCode int64 `copier:"EmployeNum"` // specify field name

	// Explicitly ignored in the destination struct.
	Salary int

	AgeTwo []int
	D Data
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

// Tags in the destination Struct provide instructions to copier.Copy to ignore
// or enforce copying and to panic or return an error if a field was not copied.
type Employee struct {
	// Tell copier.Copy to panic if this field is not copied.
	Name string `copier:"must"`

	// Tell copier.Copy to return an error if this field is not copied.
	Age int32 `copier:"must,nopanic"`

	// Tell copier.Copy to explicitly ignore copying this field.
	Salary int `copier:"-"`

	DoubleAge int32
	EmployeId int64 `copier:"EmployeNum"` // specify field name
	SuperRole string
	AgeTwo    []int
	D Data
}

type Test struct {
	A int
	B int
}

type Data struct {
	A int
	B int
	C string
	D *Test
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}

func main() {
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000, AgeTwo: []int{1, 2, 3}, D: Data{
			A: 1,
			B: 2,
			C: "3",
			D: &Test{
				A: 4,
				B: 5,
			},
		}}
		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 100000}, {Name: "jinzhu 2", Age: 30, Role: "Dev", Salary: 60000}}
		employee  = Employee{Salary: 150000}
		employees = []Employee{}
	)

	copier.CopyWithOption(&employee, &user, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    true,
	})

	fmt.Printf("%#v \n", employee)
	employee.D.D.B = 7
	fmt.Println(*employee.D.D)
	fmt.Println(*user.D.D)

	return
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    Salary:150000,            // Copying explicitly ignored
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRole: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	copier.Copy(&employees, &user)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("employees: %#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, Salary:0, DoubleAge: 60, EmployeId: 0, SuperRole: "Super Dev"},
	// }

	// Copy map to map
	map1 := map[int]int{3: 6, 4: 8}
	map2 := map[int32]int8{}
	copier.Copy(&map2, map1)

	fmt.Printf("%#v \n", map2)
	// map[int32]int8{3:6, 4:8}
}
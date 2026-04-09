package main

import "fmt"

// Entity defines the contract for activatable domain entities.
// Any type that implements activate, deactivate and info satisfies this interface.
type Entity interface {
	// activate enables the entity.
	activate()
	// deactivate disables the entity.
	deactivate()
	// info returns a human-readable summary of the entity's current state.
	info() string
}

// Address holds the physical location details of a client.
type Address struct {
	street   string
	zip_code string
	state    string
	city     string
}

// Client represents a registered user in the system.
type Client struct {
	name      string
	age       int16
	is_active bool
	address   Address
}

// activate sets the client as active.
// Implements [Entity].
func (c *Client) activate() {
	c.is_active = true
	fmt.Printf("Client %s activated\n", c.name)
}

// deactivate sets the client as inactive.
// Implements [Entity].
func (c *Client) deactivate() {
	c.is_active = false
	fmt.Printf("Client %s deactivated\n", c.name)
}

// info returns a formatted summary of the client's current state.
// Implements [Entity].
func (c *Client) info() string {
	str := fmt.Sprintf("Client{name: %s, age: %d, is_active: %v}", c.name, c.age, c.is_active)
	fmt.Printf("%s\n", str)
	return str
}

// Admin represents a privileged system user with an access level.
type Admin struct {
	name      string
	is_active bool
	// level defines the admin's permission tier.
	level int
}

// activate sets the admin as active.
// Implements [Entity].
func (a *Admin) activate() {
	a.is_active = true
	fmt.Printf("Admin %s activated\n", a.name)
}

// deactivate sets the admin as inactive.
// Implements [Entity].
func (a *Admin) deactivate() {
	a.is_active = false
	fmt.Printf("Admin %s deactivated\n", a.name)
}

// info returns a formatted summary of the admin's current state.
// Implements [Entity].
func (a *Admin) info() string {
	return fmt.Sprintf("Admin{name: %s, level: %d, is_active: %v}", a.name, a.level, a.is_active)
}

// toggle_entity activates or deactivates any [Entity] and prints its current state.
// If activate is true, calls e.activate(); otherwise calls e.deactivate().
func toggle_entity(e Entity, activate bool) {
	if activate {
		e.activate()
	} else {
		e.deactivate()
	}

	fmt.Println(e.info())
}

func main() {
	gus := &Client{
		name:      "Gus",
		age:       42,
		is_active: false,
		address: Address{
			street:   "Eugenio",
			zip_code: "02060000",
			state:    "SP",
			city:     "SP",
		},
	}

	root := &Admin{
		name:  "Root",
		level: 9,
	}

	toggle_entity(gus, true)
	toggle_entity(root, true)
	toggle_entity(gus, false)

	fmt.Printf("%v", gus)
	fmt.Println()
}

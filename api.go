package ssdir

// Student services directory
// supports multiple lookup methods

// Person structure
type Person struct {
  Name string
  Email string
  Location string
  Phone string
  Mobile string
}

// Notes structure for people and businesses
type Note struct{
  Author string
  Timestamp string
  Title string
  Text string
}

// Business listing structure
type Business struct {
  Name string
  Location string
  Phone string
  OpenHours []string
}

// Database base object
type Database struct {
  socket string
}

// Connect to application database
func (db *Database) Connect() {
  db.socket = "connected"
}

// Search for people in the directory
func (db *Database) SearchPeople(query string) ([]*Person, bool) {
  results := make([]*Person, 1)
  results[0] = new(Person)
  results[0].Name = "Val Lyashov"
  return results, true
}

// Get an individual by email address
func (db *Database) GetPerson(email string) (*Person, bool){
  person := new(Person)
  person.Email = email
  person.Name = "Val Lyashov"

  return person, true
}

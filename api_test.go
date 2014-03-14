package ssdir

// student services directory

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

// potential search phrases to be supported
/*allTests := map[string]string {
  "science student centre": ""
  "commerce student centre CAP": ""
  "engineering CAP": ""
  "venue hire": ""
  "arts timetabling": ""
  "timetables law": ""
  "westpac": ""
  "gym": ""
  "it helpdesk": ""
  "student account problem": ""
  "student housing": ""
  "science scholarships": ""
}*/

// Test people search functionality
func TestSearchPeople(t *testing.T) {
  // init database
  db := new(Database)
  db.Connect()

  // define input/output test list
  tests := map[string]string{
    "val l": "Val Lyashov",
  }

  // run through the tests
  for a,b := range tests {
    results, ok := db.SearchPeople(a)
    assert.NotNil(t, ok, "Error retrieving person")
    assert.Equal(t, b, results[0].Name, "Unexpected results returned.")
  }
}

// Test getperson functionality
func TestGetPerson(t *testing.T) {
  // init database
  db := new(Database)
  db.Connect()

  // define input/output test list
  tests := map[string]string{
    "val@plstr.com": "Val Lyashov",
  }

  // run through the tests
  for a,b := range tests {
    results, ok := db.GetPerson(a)
    assert.NotNil(t, ok, "Error retrieving person")
    assert.Equal(t, b, results.Name, "Unexpected results returned.")
  }
}

func TestGetBusiness(t *testing.T) {
// something
}

func TestGetFunctionArea(t *testing.T) {
// something
}

func TestSearchResponsibility(t *testing.T) {
// something
}



func TestSearchFunctionalArea(t *testing.T) {
// something
}

func TestGetNotes(t *testing.T) {
}

func TestCreateNote(t *testing.T) {}

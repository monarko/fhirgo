package main

import (
	"fmt"

	d "github.com/monarko/fhirgo/datatypes"
	r "github.com/monarko/fhirgo/resources"
)

func main() {
	var identifier []d.Identifier

	var active *d.Boolean

	var name []d.HumanName
	use := d.Code("official")
	family := d.String("Doe")
	given := []d.String{"John", "Michael"}
	nm := d.HumanName{Use: &use, Family: &family, Given: given}
	name = make([]d.HumanName, 0)
	name = append(name, nm)

	var telecom []d.ContactPoint

	var gender *d.Code
	gender = new(d.Code)
	*gender = "male"

	var birthDate *d.Date
	birthDate = new(d.Date)
	*birthDate = "1977-09-14"

	var deceasedBoolean *d.Boolean

	var deceasedDateTime *d.DateTime

	var address []d.Address

	var maritalStatus *d.CodeableConcept

	var multipleBirthBoolean *d.Boolean

	var multipleBirthInteger *d.Integer

	var photo []d.Attachment

	var contact []d.PatientContact

	var communication []d.PatientCommunication

	var generalPractitioner []d.Reference

	var managingOrganization *d.Reference

	var link []d.PatientLink

	pat := r.NewPatient(
		identifier,
		active,
		name,
		telecom,
		gender,
		birthDate,
		deceasedBoolean,
		deceasedDateTime,
		address,
		maritalStatus,
		multipleBirthBoolean,
		multipleBirthInteger,
		photo,
		contact,
		communication,
		generalPractitioner,
		managingOrganization,
		link,
	)

	valid, errs := pat.Validate()
	if !valid {
		fmt.Println("The document is not valid. see errors:")
		for _, e := range errs {
			fmt.Printf("- %s\n", e)
		}
	} else {
		fmt.Println("The document is valid")
	}
}

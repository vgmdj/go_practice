package simple_factory

import "testing"

func TestPhoneFactory(t *testing.T) {
	phonefactory := PhoneFactory{}
	phone := phonefactory.CreatePhone("HW")
	phone.ShowBrand()
}

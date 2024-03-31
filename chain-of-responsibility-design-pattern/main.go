package main

func main() {
	cashier := &cashier{}
	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)
	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)
	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)
	patient := &patient{name: "john"}
	//Patient visiting
	reception.execute(patient)
}

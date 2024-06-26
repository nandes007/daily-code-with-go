package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckupDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckupDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

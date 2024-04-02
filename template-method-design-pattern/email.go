package main

import "fmt"

type email struct {
	otp
}

func (e *email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (e *email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (e *email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func (e *email) publishMetric() {
	fmt.Printf("EMAIL: publishing matrics\n")
}

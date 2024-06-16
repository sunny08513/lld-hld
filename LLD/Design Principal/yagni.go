package main

/*What is YAGNI?
  Always implement things when you actually need them,
  never when you just foresee that you might need them.*/

// Over-engineered:
func process_payment1(payment_method string) {
	if payment_method == "credit_card" {
		//complex credit card handling
		return
	} else if payment_method == "paypal" {
		return
		//paypal credit card
	} else if payment_method == "bitcoin" {
		return
		//bitcoin payment handling
	}
}

// YAGNI-aligned:
func process_payment2(payment_method string) {
	if payment_method == "credit_card" {
		//credit card handling
		return
	}
}

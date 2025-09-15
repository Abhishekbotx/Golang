package main
import "fmt"

// The interface is a contract: "I can pay an amount".
type paymentBank interface{
	pay(amount float32) //interface method
}
//🌟Note:
// if you have struct with reciever type function with same name as interface method
// then that struct implements that interface
// no need to explicitly implement the interface
// this is called structural typing

type payment struct{
	gateway paymentBank //interface as a field of struct
}
// payment is a struct that contains a gateway. The gateway field is of type paymentBank,which is an interface. 
// This means that the gateway can hold any value that implements the paymentBank interface.

// In this example, the
// gateway is declared as the interface type paymentBank — so it can hold any concrete 
// implementation (e.g., stripe, razorpay, paypal) that implements pay method.

// This decouples payment from concrete gateway implementations and allows for flexibility and extensibility.
type razorpay struct{}
type stripe	 struct{}
type paypal struct{}


func(r razorpay) pay(amount float32){
	fmt.Printf("Payment of %v made using Razorpay\n",amount)
}

func (r stripe) pay(amount float32){
	fmt.Printf("Payment of %v made using Stripe\n",amount)
}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func( p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
	// p.gateway. pay is calling the pay method of the gateway (which is of type stripe here)
	// and passing the amount to it
	// this is polymorphism in action

	// explaination:
	// The makePayment method on the payment struct calls the pay method on its gateway field.
	// Since gateway is of the interface type paymentBank, it can hold any concrete implementation
	// (like stripe, razorpay, paypal) that implements the pay method.
	// When(razorpay,stripe ,etc) makePayment is called, it invokes the pay method of the specific gateway implementation
	// that was assigned to the gateway field, demonstrating polymorphism.
	
}


func main(){
	stripePaymentGw:=stripe{}
	// razorpayPaymentGw:=razorpay{}
	// paypalPaymentGw:=paypal{}
	newPayment:=payment{// need to do here with constructor function for payment struct other wise it will have compliation error
		gateway: stripePaymentGw,
	}
	newPayment.makePayment(5100)
}
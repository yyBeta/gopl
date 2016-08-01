// Package bank provides a concurrency-safe bank with one account.
package bank

var withdrawMsg struct {
	amount  int
	success chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawMsg)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func WithDraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- withdrawMsg{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case msg := <-withdraws:
			if balance >= msg.amount {
				balance -= msg.amount
				msg.success <- true
			} else {
				msg.success <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

package main

/*
Exercis e 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
prog ram. The result should indic ate whether the transaction succe e de d or fai le d due to insufficient
funds. The message sent to the monitor goroutine must contain both the amount to
withdraw and a new channel over which the monitor goroutine can send the boole an result
back to Withdraw.
*/
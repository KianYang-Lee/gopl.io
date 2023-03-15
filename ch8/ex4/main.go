package main

/*
Exercis e 8.4: Modif y the reverb2 server to use a sync.WaitGroup per connection to count
the number of active echo goroutines. When it fal ls to zero, clos e the write half of the TCP
connec tion as des cribed in Exercise 8.3. Ver ify that your modified netcat3 client from that
exercise waits for the final echoes of multiple concur rent shouts, even after the standard input
has been clos ed.
*/

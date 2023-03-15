package main

/*
Exercis e 8.13: Make the chat ser ver disconnect idle clients, such as those that have sent no
messages in the last five minutes. Hint: cal ling conn.Close() in another goroutine unblocks
ac tive Read calls such as the one done by input.Scan().
*/

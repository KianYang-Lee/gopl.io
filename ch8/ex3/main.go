package main

/*
Exercis e 8.3: In netcat3, the interface value conn has the concrete typ e *net.TCPConn, which
repres ents a TCP connec tion. A TCP connection consists of two halves that may be clos ed
indep endently using its CloseRead and CloseWrite methods. Modif y the main goroutine of
netcat3 to clos e only the write half of the connection so that the program will continue to
print the final echoes from the reverb1 server even after the standard input has been clos ed.
(D oing this for the reverb2 server is harder ; see Exercise 8.4.)
*/

package main

/*
Exercis e 8.15: Failure of any client program to read dat a in a timely manner ultimately causes
al l clients to get stuck. Modif y the broadc aster to skip a message rat her than wait if a client
wr iter is not ready to accept it. Alternat ively, add buf fer ing to each client’s outgoing message
channel so that most messages are not dropped; the broadc aster should use a non-blo cking
send to this channel.
*/

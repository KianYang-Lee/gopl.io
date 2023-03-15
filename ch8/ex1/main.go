package main

/*
Exercis e 8.1: Modif y clock2 to accept a port number, and write a program, clockwall, that
ac ts as a client of several clo ck ser vers at once, reading the times from each one and displaying
the results in a table, akin to the wal l of clo cks seen in some business offices. If you have
access to geographic ally distr ibuted computers, run instances remotely ; other wise run local
instances on dif ferent ports with fake time zones.
$ TZ=US/Eastern ./clock2 port 8010 &
$ TZ=Asia/Tokyo ./clock2 port
8020 &
$ TZ=Europe/London ./clock2 port
8030 &
$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
*/

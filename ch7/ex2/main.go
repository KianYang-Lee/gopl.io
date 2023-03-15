package main

/*
Exercis e 7.2: Write a function CountingWriter with the sig nature below that, given an
io.Writer, retur ns a new Writer that wraps the original, and a pointer to an int64 var iable
that at any moment contains the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64)
*/

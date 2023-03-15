package main

/*
Exercis e 8.10: HT TP requests may be cancelled by closing the optional Cancel channel in the
http.Request struc t. Modif y the web craw ler of Sec tion 8.6 to support cancellat ion.
Hint: the http.Get convenience function does not give you an opportunity to customize a
Request. Instead, create the request using http.NewRequest, set its Cancel field, then perform
the request by cal ling http.DefaultClient.Do(req).
*/

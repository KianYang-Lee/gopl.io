package main

/*
Exercis e 5.17: Write a var iadic function ElementsByTagName that, given an HTML node tree
and zero or more names, retur ns all the elements that match one of those names. Here are two
example cal ls:
		func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
		images := ElementsByTagName(doc, "img")
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
*/

package main

func main() {
	a := App{}

	a.Initialize("DBAdmin", "92468312#!Lex", "mydb")
	a.Run(":8000")

}

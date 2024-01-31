// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Root() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Tic Tac Toe</title><link rel=\"stylesheet\" href=\"/static/tailwind.css\"><script src=\"https://unpkg.com/htmx.org@1.9.5\" integrity=\"sha384-xcuj3WpfgjlKF+FXhSQFQ0ZNr39ln+hwjN3npfM9VBnUskLolQAcN80McRIVOPuO\" crossorigin=\"anonymous\"></script></head><body class=\"flex justify-center items-center min-h-screen bg-gray-100 p-4\"><div class=\"flex flex-col items-center space-y-4\"><div class=\"text-2xl font-bold text-center\"></div><form action=\"/reset\" method=\"post\"><input type=\"hidden\" name=\"gameID\" value=\"{{.GameID}}\"> <button class=\"w-full bg-red-800 hover:bg-red-700 text-white font-bold py-2 px-4 rounded mt-2 focus:outline-none focus:ring focus:ring-red-300\" type=\"submit\">Reset</button></form><div class=\"container mx-auto p-4 bg-white rounded-lg shadow-lg\"></div><div class=\"flex flex-col items-center space-y-4\"><form class=\"w-full max-w-xs\" action=\"/new_game\" method=\"post\"><input class=\"border border-gray-400 w-full p-2 rounded focus:border-blue-500 focus:outline-none focus:ring focus:ring-blue-200\" type=\"text\" name=\"player1\" placeholder=\"Enter Player1 name\"> <input class=\"border border-gray-400 w-full p-2 rounded focus:border-blue-500 focus:outline-none focus:ring focus:ring-blue-200\" type=\"text\" name=\"player2\" placeholder=\"Enter Player2 name\"> <button class=\"btn btn-primary\" type=\"submit\">Submit</button></form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

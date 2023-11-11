package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {

		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] token type wrong, expected=%q, got=%q", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] literal wrong, expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWithFunc(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = function(x, y) {
				x + y;
		}

		let result = add(five, ten);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, test := range tests {

		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] token type wrong, expected=%q, got=%q", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] literal wrong, expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

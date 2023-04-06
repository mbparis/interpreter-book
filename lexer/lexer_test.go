package lexer_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mbparis/interpreter-book/lexer"
	"github.com/mbparis/interpreter-book/monkey/token"
)

var _ = Describe("Lexer", func() {
	Describe("will iterate over source code; base set", Ordered, func() {
		inputBase := `=+(){},;`
		sut := lexer.New(inputBase)

		DescribeTable("outputting source code as tokens",
			func(tokenType token.TokenType, tokenLiteral string) {
				actual := sut.NextToken()
				Expect(actual.Type).To(Equal(tokenType))
				Expect(actual.Literal).To(Equal(tokenLiteral))

			},
			Entry("when ASSIGN", token.TokenType(token.ASSIGN), "="),
			Entry("when PLUS", token.TokenType(token.PLUS), "+"),
			Entry("when LPAREN", token.TokenType(token.LPAREN), "("),
			Entry("when RPAREN", token.TokenType(token.RPAREN), ")"),
			Entry("when LBRACE", token.TokenType(token.LBRACE), "{"),
			Entry("when RBRACE", token.TokenType(token.RBRACE), "}"),
			Entry("when COMMA", token.TokenType(token.COMMA), ","),
			Entry("when SEMICOLON", token.TokenType(token.SEMICOLON), ";"),
			Entry("when EOF", token.TokenType(token.EOF), ""),
		)

	})
	Describe("will iterate over source code; subset", Ordered, func() {
		input := `let five = 5;
				let ten = 10;

				let add = fn(x, y) {
				x + y;
				};

				let result = add(five, ten);
				`
		sut := lexer.New(input)

		DescribeTable("outputting source code as tokens",
			func(tokenType token.TokenType, tokenLiteral string) {
				actual := sut.NextToken()
				fmt.Printf("__actual: %+v", actual)
				Expect(actual.Type).To(Equal(tokenType))
				Expect(actual.Literal).To(Equal(tokenLiteral))
			},
			Entry("when LET", token.TokenType(token.LET), "let"),
			Entry("when IDENT", token.TokenType(token.IDENT), "five"),
			Entry("when ASSIGN", token.TokenType(token.ASSIGN), "="),
			Entry("when INT", token.TokenType(token.INT), "5"),
			Entry("when SEMICOLON", token.TokenType(token.SEMICOLON), ";"),
			Entry("when LET", token.TokenType(token.LET), "let"),
			Entry("when IDENT", token.TokenType(token.IDENT), "ten"),
			Entry("when ASSIGN", token.TokenType(token.ASSIGN), "="),
			Entry("when INT", token.TokenType(token.INT), "10"),
			Entry("when SEMICOLON", token.TokenType(token.SEMICOLON), ";"),
			Entry("when LET", token.TokenType(token.LET), "let"),
			Entry("when IDENT", token.TokenType(token.IDENT), "add"),
			Entry("when ASSIGN", token.TokenType(token.ASSIGN), "="),
			Entry("when FUNCTION", token.TokenType(token.FUNCTION), "fn"),
			Entry("when LPAREN", token.TokenType(token.LPAREN), "("),
			Entry("when IDENT", token.TokenType(token.IDENT), "x"),
			Entry("when COMMA", token.TokenType(token.COMMA), ","),
			Entry("when IDENT", token.TokenType(token.IDENT), "y"),
			Entry("when RPAREN", token.TokenType(token.RPAREN), ")"),
			Entry("when LBRACE", token.TokenType(token.LBRACE), "{"),
			Entry("when IDENT", token.TokenType(token.IDENT), "x"),
			Entry("when PLUS", token.TokenType(token.PLUS), "+"),
			Entry("when IDENT", token.TokenType(token.IDENT), "y"), //
			Entry("when SEMICOLON", token.TokenType(token.SEMICOLON), ";"),
			Entry("when RBRACE", token.TokenType(token.RBRACE), "}"),
			Entry("when SEMICOLON", token.TokenType(token.SEMICOLON), ";"),
			Entry("when LET", token.TokenType(token.LET), "let"),
			Entry("when IDENT", token.TokenType(token.IDENT), "result"),
			Entry("when ASSIGN", token.TokenType(token.ASSIGN), "="),
			Entry("when IDENT", token.TokenType(token.IDENT), "add"),
			Entry("when LPAREN", token.TokenType(token.LPAREN), "("),
			Entry("when IDENT", token.TokenType(token.IDENT), "five"),
			Entry("when COMMA", token.TokenType(token.COMMA), ","),
			Entry("when IDENT", token.TokenType(token.IDENT), "ten"),
			Entry("when RPAREN", token.TokenType(token.RPAREN), ")"),
			Entry("when SEMICOLON", token.TokenType(token.SEMICOLON), ";"),
			Entry("when EOF", token.TokenType(token.EOF), ""),
		)

	})
})

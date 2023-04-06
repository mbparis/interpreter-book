package lexer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mbparis/interpreter-book/lexer"
	"github.com/mbparis/interpreter-book/monkey/token"
)

func wrapTokenType(toWrap string) token.TokenType {
	return token.TokenType(toWrap)

}

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
				Expect(actual.Type).To(Equal(tokenType))
				Expect(actual.Literal).To(Equal(tokenLiteral))
			},
			Entry(" ", token.TokenType(token.LET), "let"),
			Entry(" ", token.TokenType(token.IDENT), "five"),
			Entry(" ", token.TokenType(token.ASSIGN), "="),
			Entry(" ", token.TokenType(token.INT), "5"),
			Entry(" ", token.TokenType(token.SEMICOLON), ";"),
			Entry(" ", token.TokenType(token.LET), "let"),
			Entry(" ", token.TokenType(token.IDENT), "ten"),
			Entry(" ", token.TokenType(token.ASSIGN), "="),
			Entry(" ", token.TokenType(token.INT), "10"),
			Entry(" ", token.TokenType(token.SEMICOLON), ";"),
			Entry(" ", token.TokenType(token.LET), "let"),
			Entry(" ", token.TokenType(token.IDENT), "add"),
			Entry(" ", token.TokenType(token.ASSIGN), "="),
			Entry(" ", token.TokenType(token.FUNCTION), "fn"),
			Entry(" ", token.TokenType(token.LPAREN), "("),
			Entry(" ", token.TokenType(token.IDENT), "x"),
			Entry(" ", token.TokenType(token.COMMA), ","),
			Entry(" ", token.TokenType(token.IDENT), "y"),
			Entry(" ", token.TokenType(token.RPAREN), ")"),
			Entry(" ", token.TokenType(token.LBRACE), "{"),
			Entry(" ", token.TokenType(token.IDENT), "x"),
			Entry(" ", token.TokenType(token.PLUS), "+"),
			Entry(" ", token.TokenType(token.IDENT), "y"), //
			Entry(" ", token.TokenType(token.SEMICOLON), ";"),
			Entry(" ", token.TokenType(token.RBRACE), "}"),
			Entry(" ", token.TokenType(token.SEMICOLON), ";"),
			Entry(" ", token.TokenType(token.LET), "let"),
			Entry(" ", token.TokenType(token.IDENT), "result"),
			Entry(" ", token.TokenType(token.ASSIGN), "="),
			Entry(" ", token.TokenType(token.IDENT), "add"),
			Entry(" ", token.TokenType(token.LPAREN), "("),
			Entry(" ", token.TokenType(token.IDENT), "five"),
			Entry(" ", token.TokenType(token.COMMA), ","),
			Entry(" ", token.TokenType(token.IDENT), "ten"),
			Entry(" ", token.TokenType(token.RPAREN), ")"),
			Entry(" ", token.TokenType(token.SEMICOLON), ";"),
			Entry(" ", token.TokenType(token.EOF), ""),
		)

	})
})

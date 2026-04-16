#pragma once

#include <string>
#include <vector>

#include "token.h"

namespace cminus {

class Lexer {
 public:
  explicit Lexer(std::string source);

  std::vector<Token> tokenize();

 private:
  std::string source_;
  size_t pos_;
  int line_;
  int column_;

  bool isAtEnd() const;
  char peek() const;
  char peekNext() const;
  char advance();
  bool match(char expected);

  void skipWhitespaceAndComments();

  Token makeToken(TokenType type,
                  const std::string& lexeme,
                  int line,
                  int column) const;
  Token identifierOrKeyword(int startLine, int startColumn);
  Token number(int startLine, int startColumn);
  Token symbol(int startLine, int startColumn);
};

}  // namespace cminus

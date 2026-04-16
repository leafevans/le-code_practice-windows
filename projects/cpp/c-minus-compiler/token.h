#pragma once

#include <string>

namespace cminus {

enum class TokenType {
  EndOfFile,
  Error,

  Identifier,
  Number,

  KeywordIf,
  KeywordElse,
  KeywordInt,
  KeywordReturn,
  KeywordVoid,
  KeywordWhile,

  Plus,
  Minus,
  Star,
  Slash,

  Less,
  LessEqual,
  Greater,
  GreaterEqual,
  EqualEqual,
  NotEqual,
  Assign,

  Semicolon,
  Comma,
  LeftParen,
  RightParen,
  LeftBracket,
  RightBracket,
  LeftBrace,
  RightBrace
};

struct Token {
  TokenType type;
  std::string lexeme;
  int line;
  int column;
};

std::string tokenTypeName(TokenType type);

}  // namespace cminus

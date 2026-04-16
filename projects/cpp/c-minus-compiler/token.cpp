#include "token.h"

namespace cminus {

std::string tokenTypeName(TokenType type) {
  switch (type) {
    case TokenType::EndOfFile:
      return "EOF";
    case TokenType::Error:
      return "ERROR";
    case TokenType::Identifier:
      return "IDENTIFIER";
    case TokenType::Number:
      return "NUMBER";
    case TokenType::KeywordIf:
      return "IF";
    case TokenType::KeywordElse:
      return "ELSE";
    case TokenType::KeywordInt:
      return "INT";
    case TokenType::KeywordReturn:
      return "RETURN";
    case TokenType::KeywordVoid:
      return "VOID";
    case TokenType::KeywordWhile:
      return "WHILE";
    case TokenType::Plus:
      return "PLUS";
    case TokenType::Minus:
      return "MINUS";
    case TokenType::Star:
      return "STAR";
    case TokenType::Slash:
      return "SLASH";
    case TokenType::Less:
      return "LESS";
    case TokenType::LessEqual:
      return "LESS_EQUAL";
    case TokenType::Greater:
      return "GREATER";
    case TokenType::GreaterEqual:
      return "GREATER_EQUAL";
    case TokenType::EqualEqual:
      return "EQUAL_EQUAL";
    case TokenType::NotEqual:
      return "NOT_EQUAL";
    case TokenType::Assign:
      return "ASSIGN";
    case TokenType::Semicolon:
      return "SEMICOLON";
    case TokenType::Comma:
      return "COMMA";
    case TokenType::LeftParen:
      return "LEFT_PAREN";
    case TokenType::RightParen:
      return "RIGHT_PAREN";
    case TokenType::LeftBracket:
      return "LEFT_BRACKET";
    case TokenType::RightBracket:
      return "RIGHT_BRACKET";
    case TokenType::LeftBrace:
      return "LEFT_BRACE";
    case TokenType::RightBrace:
      return "RIGHT_BRACE";
  }
  return "UNKNOWN";
}

}  // namespace cminus

#include "lexer.h"

#include <cctype>
#include <unordered_map>

namespace cminus {

Lexer::Lexer(std::string source)
    : source_(std::move(source)), pos_(0), line_(1), column_(1) {}

std::vector<Token> Lexer::tokenize() {
  std::vector<Token> tokens;

  while (true) {
    skipWhitespaceAndComments();

    if (isAtEnd()) {
      tokens.push_back(makeToken(TokenType::EndOfFile, "", line_, column_));
      break;
    }

    int startLine   = line_;
    int startColumn = column_;
    char c          = peek();

    if (std::isalpha(static_cast<unsigned char>(c)) || c == '_') {
      tokens.push_back(identifierOrKeyword(startLine, startColumn));
      continue;
    }

    if (std::isdigit(static_cast<unsigned char>(c))) {
      tokens.push_back(number(startLine, startColumn));
      continue;
    }

    tokens.push_back(symbol(startLine, startColumn));
  }

  return tokens;
}

bool Lexer::isAtEnd() const { return pos_ >= source_.size(); }

char Lexer::peek() const {
  if (isAtEnd()) {
    return '\0';
  }
  return source_[pos_];
}

char Lexer::peekNext() const {
  if (pos_ + 1 >= source_.size()) {
    return '\0';
  }
  return source_[pos_ + 1];
}

char Lexer::advance() {
  char c = source_[pos_++];
  if (c == '\n') {
    ++line_;
    column_ = 1;
  } else {
    ++column_;
  }
  return c;
}

bool Lexer::match(char expected) {
  if (isAtEnd() || source_[pos_] != expected) {
    return false;
  }
  advance();
  return true;
}

void Lexer::skipWhitespaceAndComments() {
  while (!isAtEnd()) {
    char c = peek();

    if (std::isspace(static_cast<unsigned char>(c))) {
      advance();
      continue;
    }

    if (c == '/' && peekNext() == '/') {
      while (!isAtEnd() && peek() != '\n') {
        advance();
      }
      continue;
    }

    if (c == '/' && peekNext() == '*') {
      advance();
      advance();
      while (!isAtEnd()) {
        if (peek() == '*' && peekNext() == '/') {
          advance();
          advance();
          break;
        }
        advance();
      }
      continue;
    }

    break;
  }
}

Token Lexer::makeToken(TokenType type,
                       const std::string& lexeme,
                       int line,
                       int column) const {
  return Token{type, lexeme, line, column};
}

Token Lexer::identifierOrKeyword(int startLine, int startColumn) {
  size_t begin = pos_;
  while (std::isalnum(static_cast<unsigned char>(peek())) || peek() == '_') {
    advance();
  }

  std::string text = source_.substr(begin, pos_ - begin);
  static const std::unordered_map<std::string, TokenType> keywordMap = {
      {"if", TokenType::KeywordIf},     {"else", TokenType::KeywordElse},
      {"int", TokenType::KeywordInt},   {"return", TokenType::KeywordReturn},
      {"void", TokenType::KeywordVoid}, {"while", TokenType::KeywordWhile},
  };

  auto it = keywordMap.find(text);
  if (it != keywordMap.end()) {
    return makeToken(it->second, text, startLine, startColumn);
  }

  return makeToken(TokenType::Identifier, text, startLine, startColumn);
}

Token Lexer::number(int startLine, int startColumn) {
  size_t begin = pos_;
  while (std::isdigit(static_cast<unsigned char>(peek()))) {
    advance();
  }
  return makeToken(TokenType::Number, source_.substr(begin, pos_ - begin),
                   startLine, startColumn);
}

Token Lexer::symbol(int startLine, int startColumn) {
  char c = advance();
  switch (c) {
    case '+':
      return makeToken(TokenType::Plus, "+", startLine, startColumn);
    case '-':
      return makeToken(TokenType::Minus, "-", startLine, startColumn);
    case '*':
      return makeToken(TokenType::Star, "*", startLine, startColumn);
    case '/':
      return makeToken(TokenType::Slash, "/", startLine, startColumn);
    case '<':
      if (match('=')) {
        return makeToken(TokenType::LessEqual, "<=", startLine, startColumn);
      }
      return makeToken(TokenType::Less, "<", startLine, startColumn);
    case '>':
      if (match('=')) {
        return makeToken(TokenType::GreaterEqual, ">=", startLine, startColumn);
      }
      return makeToken(TokenType::Greater, ">", startLine, startColumn);
    case '=':
      if (match('=')) {
        return makeToken(TokenType::EqualEqual, "==", startLine, startColumn);
      }
      return makeToken(TokenType::Assign, "=", startLine, startColumn);
    case '!':
      if (match('=')) {
        return makeToken(TokenType::NotEqual, "!=", startLine, startColumn);
      }
      return makeToken(TokenType::Error, "!", startLine, startColumn);
    case ';':
      return makeToken(TokenType::Semicolon, ";", startLine, startColumn);
    case ',':
      return makeToken(TokenType::Comma, ",", startLine, startColumn);
    case '(':
      return makeToken(TokenType::LeftParen, "(", startLine, startColumn);
    case ')':
      return makeToken(TokenType::RightParen, ")", startLine, startColumn);
    case '[':
      return makeToken(TokenType::LeftBracket, "[", startLine, startColumn);
    case ']':
      return makeToken(TokenType::RightBracket, "]", startLine, startColumn);
    case '{':
      return makeToken(TokenType::LeftBrace, "{", startLine, startColumn);
    case '}':
      return makeToken(TokenType::RightBrace, "}", startLine, startColumn);
    default:
      return makeToken(TokenType::Error, std::string(1, c), startLine,
                       startColumn);
  }
}

}  // namespace cminus

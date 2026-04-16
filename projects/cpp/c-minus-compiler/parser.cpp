#include "parser.h"

#include <cstdlib>
#include <sstream>
#include <stdexcept>
#include <utility>

namespace cminus {

Parser::Parser(std::vector<Token> tokens)
    : tokens_(std::move(tokens)), current_(0) {}

std::unique_ptr<Program> Parser::parseProgram() {
  auto program = std::make_unique<Program>();

  while (!isAtEnd()) {
    if (check(TokenType::EndOfFile)) {
      break;
    }
    program->declarations.push_back(parseDeclaration());
  }

  return program;
}

bool Parser::isAtEnd() const {
  return current_ >= tokens_.size() ||
         tokens_[current_].type == TokenType::EndOfFile;
}

const Token& Parser::peek() const { return tokens_[current_]; }

const Token& Parser::previous() const { return tokens_[current_ - 1]; }

const Token& Parser::lookahead(size_t offset) const {
  size_t index = current_ + offset;
  if (index >= tokens_.size()) {
    return tokens_.back();
  }
  return tokens_[index];
}

bool Parser::check(TokenType type) const {
  if (isAtEnd()) {
    return type == TokenType::EndOfFile;
  }
  return peek().type == type;
}

bool Parser::match(TokenType type) {
  if (!check(type)) {
    return false;
  }
  advance();
  return true;
}

const Token& Parser::advance() {
  if (!isAtEnd()) {
    ++current_;
  }
  return previous();
}

const Token& Parser::consume(TokenType type, const std::string& message) {
  if (check(type)) {
    return advance();
  }
  parseError(message);
}

[[noreturn]] void Parser::parseError(const std::string& message) const {
  std::ostringstream oss;
  const Token& t = peek();
  oss << "Parse error at line " << t.line << ", column " << t.column << ": "
      << message << ". Got token " << tokenTypeName(t.type) << " (\""
      << t.lexeme << "\")";
  throw std::runtime_error(oss.str());
}

std::string Parser::parseTypeSpecifier() {
  if (match(TokenType::KeywordInt)) {
    return "int";
  }
  if (match(TokenType::KeywordVoid)) {
    return "void";
  }
  parseError("expected type specifier int or void");
}

std::unique_ptr<Decl> Parser::parseDeclaration() {
  std::string typeName = parseTypeSpecifier();
  Token id = consume(TokenType::Identifier, "expected identifier after type");

  if (match(TokenType::LeftParen)) {
    return parseFunDeclAfterTypeAndName(typeName, id.lexeme);
  }

  return std::make_unique<GlobalVarDecl>(
      parseVarDeclAfterTypeAndName(typeName, id.lexeme));
}

std::unique_ptr<VarDecl> Parser::parseVarDeclAfterTypeAndName(
    std::string typeName, std::string name) {
  bool isArray  = false;
  int arraySize = 0;

  if (match(TokenType::LeftBracket)) {
    Token sizeToken =
        consume(TokenType::Number, "expected number in array declaration");
    arraySize = std::atoi(sizeToken.lexeme.c_str());
    consume(TokenType::RightBracket, "expected ] in array declaration");
    isArray = true;
  }

  consume(TokenType::Semicolon, "expected ; after variable declaration");
  return std::make_unique<VarDecl>(std::move(typeName), std::move(name),
                                   isArray, arraySize);
}

std::unique_ptr<FunDecl> Parser::parseFunDeclAfterTypeAndName(
    std::string returnType, std::string name) {
  std::vector<std::unique_ptr<Param>> params = parseParams();
  consume(TokenType::RightParen, "expected ) after function parameters");
  auto body = parseCompoundStmt();
  return std::make_unique<FunDecl>(std::move(returnType), std::move(name),
                                   std::move(params), std::move(body));
}

std::vector<std::unique_ptr<Param>> Parser::parseParams() {
  std::vector<std::unique_ptr<Param>> params;

  if (check(TokenType::KeywordVoid) &&
      lookahead(1).type == TokenType::RightParen) {
    advance();
    return params;
  }

  params.push_back(parseParam());
  while (match(TokenType::Comma)) {
    params.push_back(parseParam());
  }

  return params;
}

std::unique_ptr<Param> Parser::parseParam() {
  std::string typeName = parseTypeSpecifier();
  Token id = consume(TokenType::Identifier, "expected parameter name");

  bool isArray = false;
  if (match(TokenType::LeftBracket)) {
    consume(TokenType::RightBracket,
            "expected ] in parameter array declaration");
    isArray = true;
  }

  return std::make_unique<Param>(typeName, id.lexeme, isArray);
}

std::unique_ptr<CompoundStmt> Parser::parseCompoundStmt() {
  consume(TokenType::LeftBrace, "expected { to begin compound statement");
  auto compound = std::make_unique<CompoundStmt>();

  while (check(TokenType::KeywordInt) || check(TokenType::KeywordVoid)) {
    std::string typeName = parseTypeSpecifier();
    Token id = consume(TokenType::Identifier, "expected local variable name");

    if (check(TokenType::LeftParen)) {
      parseError("nested function declaration is not allowed");
    }

    compound->localDecls.push_back(
        parseVarDeclAfterTypeAndName(typeName, id.lexeme));
  }

  while (!check(TokenType::RightBrace) && !isAtEnd()) {
    compound->statements.push_back(parseStatement());
  }

  consume(TokenType::RightBrace, "expected } after compound statement");
  return compound;
}

std::unique_ptr<Stmt> Parser::parseStatement() {
  if (check(TokenType::LeftBrace)) {
    return parseCompoundStmt();
  }
  if (check(TokenType::KeywordIf)) {
    return parseSelectionStmt();
  }
  if (check(TokenType::KeywordWhile)) {
    return parseIterationStmt();
  }
  if (check(TokenType::KeywordReturn)) {
    return parseReturnStmt();
  }
  return parseExpressionStmt();
}

std::unique_ptr<Stmt> Parser::parseExpressionStmt() {
  if (match(TokenType::Semicolon)) {
    return std::make_unique<ExprStmt>(nullptr);
  }

  auto expr = parseExpression();
  consume(TokenType::Semicolon, "expected ; after expression");
  return std::make_unique<ExprStmt>(std::move(expr));
}

std::unique_ptr<Stmt> Parser::parseSelectionStmt() {
  consume(TokenType::KeywordIf, "expected if");
  consume(TokenType::LeftParen, "expected ( after if");
  auto condition = parseExpression();
  consume(TokenType::RightParen, "expected ) after if condition");

  auto thenBranch                  = parseStatement();
  std::unique_ptr<Stmt> elseBranch = nullptr;
  if (match(TokenType::KeywordElse)) {
    elseBranch = parseStatement();
  }

  return std::make_unique<IfStmt>(std::move(condition), std::move(thenBranch),
                                  std::move(elseBranch));
}

std::unique_ptr<Stmt> Parser::parseIterationStmt() {
  consume(TokenType::KeywordWhile, "expected while");
  consume(TokenType::LeftParen, "expected ( after while");
  auto condition = parseExpression();
  consume(TokenType::RightParen, "expected ) after while condition");
  auto body = parseStatement();
  return std::make_unique<WhileStmt>(std::move(condition), std::move(body));
}

std::unique_ptr<Stmt> Parser::parseReturnStmt() {
  consume(TokenType::KeywordReturn, "expected return");

  if (match(TokenType::Semicolon)) {
    return std::make_unique<ReturnStmt>(nullptr);
  }

  auto expr = parseExpression();
  consume(TokenType::Semicolon, "expected ; after return expression");
  return std::make_unique<ReturnStmt>(std::move(expr));
}

std::unique_ptr<Expr> Parser::parseExpression() {
  if (check(TokenType::Identifier)) {
    size_t saved = current_;
    Token id     = advance();
    auto lhs     = parseVarWithName(id.lexeme);

    if (match(TokenType::Assign)) {
      auto rhs = parseExpression();
      return std::make_unique<AssignExpr>(std::move(lhs), std::move(rhs));
    }

    current_ = saved;
  }

  return parseSimpleExpression();
}

std::unique_ptr<Expr> Parser::parseSimpleExpression() {
  auto left = parseAdditiveExpression();

  if (isRelOp(peek().type)) {
    std::string op = advance().lexeme;
    auto right     = parseAdditiveExpression();
    return std::make_unique<BinaryExpr>(op, std::move(left), std::move(right));
  }

  return left;
}

std::unique_ptr<Expr> Parser::parseAdditiveExpression() {
  auto expr = parseTerm();

  while (isAddOp(peek().type)) {
    std::string op = advance().lexeme;
    auto right     = parseTerm();
    expr = std::make_unique<BinaryExpr>(op, std::move(expr), std::move(right));
  }

  return expr;
}

std::unique_ptr<Expr> Parser::parseTerm() {
  auto expr = parseFactor();

  while (isMulOp(peek().type)) {
    std::string op = advance().lexeme;
    auto right     = parseFactor();
    expr = std::make_unique<BinaryExpr>(op, std::move(expr), std::move(right));
  }

  return expr;
}

std::unique_ptr<Expr> Parser::parseFactor() {
  if (match(TokenType::LeftParen)) {
    auto expr = parseExpression();
    consume(TokenType::RightParen, "expected ) after expression");
    return expr;
  }

  if (match(TokenType::Number)) {
    int value = std::atoi(previous().lexeme.c_str());
    return std::make_unique<NumExpr>(value);
  }

  if (match(TokenType::Identifier)) {
    std::string name = previous().lexeme;

    if (match(TokenType::LeftParen)) {
      auto args = parseArgs();
      consume(TokenType::RightParen, "expected ) after arguments");
      return std::make_unique<CallExpr>(name, std::move(args));
    }

    return parseVarWithName(name);
  }

  parseError("expected expression factor");
}

std::unique_ptr<Expr> Parser::parseVarWithName(std::string name) {
  if (match(TokenType::LeftBracket)) {
    auto idx = parseExpression();
    consume(TokenType::RightBracket, "expected ] after array index");
    return std::make_unique<VarExpr>(std::move(name), std::move(idx));
  }

  return std::make_unique<VarExpr>(std::move(name), nullptr);
}

std::vector<std::unique_ptr<Expr>> Parser::parseArgs() {
  std::vector<std::unique_ptr<Expr>> args;

  if (check(TokenType::RightParen)) {
    return args;
  }

  args.push_back(parseExpression());
  while (match(TokenType::Comma)) {
    args.push_back(parseExpression());
  }

  return args;
}

bool Parser::isRelOp(TokenType type) const {
  return type == TokenType::Less || type == TokenType::LessEqual ||
         type == TokenType::Greater || type == TokenType::GreaterEqual ||
         type == TokenType::EqualEqual || type == TokenType::NotEqual;
}

bool Parser::isAddOp(TokenType type) const {
  return type == TokenType::Plus || type == TokenType::Minus;
}

bool Parser::isMulOp(TokenType type) const {
  return type == TokenType::Star || type == TokenType::Slash;
}

}  // namespace cminus

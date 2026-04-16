#pragma once

#include <memory>
#include <string>
#include <vector>

#include "ast.h"
#include "token.h"

namespace cminus {

class Parser {
 public:
  explicit Parser(std::vector<Token> tokens);

  std::unique_ptr<Program> parseProgram();

 private:
  std::vector<Token> tokens_;
  size_t current_;

  bool isAtEnd() const;
  const Token& peek() const;
  const Token& previous() const;
  const Token& lookahead(size_t offset) const;

  bool check(TokenType type) const;
  bool match(TokenType type);
  const Token& advance();
  const Token& consume(TokenType type, const std::string& message);

  [[noreturn]] void parseError(const std::string& message) const;

  std::string parseTypeSpecifier();
  std::unique_ptr<Decl> parseDeclaration();
  std::unique_ptr<VarDecl> parseVarDeclAfterTypeAndName(std::string typeName,
                                                        std::string name);
  std::unique_ptr<FunDecl> parseFunDeclAfterTypeAndName(std::string returnType,
                                                        std::string name);

  std::vector<std::unique_ptr<Param>> parseParams();
  std::unique_ptr<Param> parseParam();

  std::unique_ptr<CompoundStmt> parseCompoundStmt();
  std::unique_ptr<Stmt> parseStatement();
  std::unique_ptr<Stmt> parseExpressionStmt();
  std::unique_ptr<Stmt> parseSelectionStmt();
  std::unique_ptr<Stmt> parseIterationStmt();
  std::unique_ptr<Stmt> parseReturnStmt();

  std::unique_ptr<Expr> parseExpression();
  std::unique_ptr<Expr> parseSimpleExpression();
  std::unique_ptr<Expr> parseAdditiveExpression();
  std::unique_ptr<Expr> parseTerm();
  std::unique_ptr<Expr> parseFactor();
  std::unique_ptr<Expr> parseVarWithName(std::string name);
  std::vector<std::unique_ptr<Expr>> parseArgs();

  bool isRelOp(TokenType type) const;
  bool isAddOp(TokenType type) const;
  bool isMulOp(TokenType type) const;
};

}  // namespace cminus

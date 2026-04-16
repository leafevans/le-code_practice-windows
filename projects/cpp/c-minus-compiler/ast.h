#pragma once

#include <memory>
#include <ostream>
#include <string>
#include <vector>

namespace cminus {

inline void printIndent(std::ostream& out, int indent) {
  for (int i = 0; i < indent; ++i) {
    out << "  ";
  }
}

struct Node {
  virtual ~Node()                                         = default;
  virtual void print(std::ostream& out, int indent) const = 0;
};

struct Expr : Node {};
struct Stmt : Node {};
struct Decl : Node {};

struct VarExpr : Expr {
  std::string name;
  std::unique_ptr<Expr> index;

  VarExpr(std::string n, std::unique_ptr<Expr> idx)
      : name(std::move(n)), index(std::move(idx)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "VarExpr name=" << name;
    if (index) {
      out << " [\n";
      index->print(out, indent + 1);
      out << "\n";
      printIndent(out, indent);
      out << "]";
    }
  }
};

struct NumExpr : Expr {
  int value;

  explicit NumExpr(int v) : value(v) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "NumExpr value=" << value;
  }
};

struct CallExpr : Expr {
  std::string callee;
  std::vector<std::unique_ptr<Expr>> args;

  CallExpr(std::string c, std::vector<std::unique_ptr<Expr>> a)
      : callee(std::move(c)), args(std::move(a)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "CallExpr callee=" << callee << " args=[\n";
    for (size_t i = 0; i < args.size(); ++i) {
      args[i]->print(out, indent + 1);
      if (i + 1 < args.size()) {
        out << "\n";
      }
    }
    out << "\n";
    printIndent(out, indent);
    out << "]";
  }
};

struct BinaryExpr : Expr {
  std::string op;
  std::unique_ptr<Expr> left;
  std::unique_ptr<Expr> right;

  BinaryExpr(std::string o, std::unique_ptr<Expr> l, std::unique_ptr<Expr> r)
      : op(std::move(o)), left(std::move(l)), right(std::move(r)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "BinaryExpr op=" << op << " (\n";
    left->print(out, indent + 1);
    out << "\n";
    right->print(out, indent + 1);
    out << "\n";
    printIndent(out, indent);
    out << ")";
  }
};

struct AssignExpr : Expr {
  std::unique_ptr<Expr> lhs;
  std::unique_ptr<Expr> rhs;

  AssignExpr(std::unique_ptr<Expr> l, std::unique_ptr<Expr> r)
      : lhs(std::move(l)), rhs(std::move(r)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "AssignExpr (\n";
    lhs->print(out, indent + 1);
    out << "\n";
    rhs->print(out, indent + 1);
    out << "\n";
    printIndent(out, indent);
    out << ")";
  }
};

struct VarDecl : Node {
  std::string typeName;
  std::string name;
  bool isArray;
  int arraySize;

  VarDecl(std::string t, std::string n, bool isArr, int size)
      : typeName(std::move(t)),
        name(std::move(n)),
        isArray(isArr),
        arraySize(size) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "VarDecl type=" << typeName << " name=" << name;
    if (isArray) {
      out << " size=" << arraySize;
    }
  }
};

struct CompoundStmt;

struct Param : Node {
  std::string typeName;
  std::string name;
  bool isArray;

  Param(std::string t, std::string n, bool isArr)
      : typeName(std::move(t)), name(std::move(n)), isArray(isArr) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "Param type=" << typeName << " name=" << name;
    if (isArray) {
      out << "[]";
    }
  }
};

struct ExprStmt : Stmt {
  std::unique_ptr<Expr> expr;

  explicit ExprStmt(std::unique_ptr<Expr> e) : expr(std::move(e)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "ExprStmt";
    if (expr) {
      out << " (\n";
      expr->print(out, indent + 1);
      out << "\n";
      printIndent(out, indent);
      out << ")";
    }
  }
};

struct IfStmt : Stmt {
  std::unique_ptr<Expr> condition;
  std::unique_ptr<Stmt> thenBranch;
  std::unique_ptr<Stmt> elseBranch;

  IfStmt(std::unique_ptr<Expr> c,
         std::unique_ptr<Stmt> t,
         std::unique_ptr<Stmt> e)
      : condition(std::move(c)),
        thenBranch(std::move(t)),
        elseBranch(std::move(e)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "IfStmt (\n";
    printIndent(out, indent + 1);
    out << "Condition:\n";
    condition->print(out, indent + 2);
    out << "\n";
    printIndent(out, indent + 1);
    out << "Then:\n";
    thenBranch->print(out, indent + 2);
    if (elseBranch) {
      out << "\n";
      printIndent(out, indent + 1);
      out << "Else:\n";
      elseBranch->print(out, indent + 2);
    }
    out << "\n";
    printIndent(out, indent);
    out << ")";
  }
};

struct WhileStmt : Stmt {
  std::unique_ptr<Expr> condition;
  std::unique_ptr<Stmt> body;

  WhileStmt(std::unique_ptr<Expr> c, std::unique_ptr<Stmt> b)
      : condition(std::move(c)), body(std::move(b)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "WhileStmt (\n";
    condition->print(out, indent + 1);
    out << "\n";
    body->print(out, indent + 1);
    out << "\n";
    printIndent(out, indent);
    out << ")";
  }
};

struct ReturnStmt : Stmt {
  std::unique_ptr<Expr> expr;

  explicit ReturnStmt(std::unique_ptr<Expr> e) : expr(std::move(e)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "ReturnStmt";
    if (expr) {
      out << " (\n";
      expr->print(out, indent + 1);
      out << "\n";
      printIndent(out, indent);
      out << ")";
    }
  }
};

struct CompoundStmt : Stmt {
  std::vector<std::unique_ptr<VarDecl>> localDecls;
  std::vector<std::unique_ptr<Stmt>> statements;

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "CompoundStmt {\n";
    printIndent(out, indent + 1);
    out << "LocalDecls:\n";
    for (const auto& decl : localDecls) {
      decl->print(out, indent + 2);
      out << "\n";
    }
    printIndent(out, indent + 1);
    out << "Statements:\n";
    for (const auto& stmt : statements) {
      stmt->print(out, indent + 2);
      out << "\n";
    }
    printIndent(out, indent);
    out << "}";
  }
};

struct FunDecl : Decl {
  std::string returnType;
  std::string name;
  std::vector<std::unique_ptr<Param>> params;
  std::unique_ptr<CompoundStmt> body;

  FunDecl(std::string rt,
          std::string n,
          std::vector<std::unique_ptr<Param>> p,
          std::unique_ptr<CompoundStmt> b)
      : returnType(std::move(rt)),
        name(std::move(n)),
        params(std::move(p)),
        body(std::move(b)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "FunDecl returnType=" << returnType << " name=" << name << " (\n";
    printIndent(out, indent + 1);
    out << "Params:\n";
    for (const auto& param : params) {
      param->print(out, indent + 2);
      out << "\n";
    }
    printIndent(out, indent + 1);
    out << "Body:\n";
    body->print(out, indent + 2);
    out << "\n";
    printIndent(out, indent);
    out << ")";
  }
};

struct GlobalVarDecl : Decl {
  std::unique_ptr<VarDecl> var;

  explicit GlobalVarDecl(std::unique_ptr<VarDecl> v) : var(std::move(v)) {}

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "Global";
    out << " (\n";
    var->print(out, indent + 1);
    out << "\n";
    printIndent(out, indent);
    out << ")";
  }
};

struct Program : Node {
  std::vector<std::unique_ptr<Decl>> declarations;

  void print(std::ostream& out, int indent) const override {
    printIndent(out, indent);
    out << "Program [\n";
    for (size_t i = 0; i < declarations.size(); ++i) {
      declarations[i]->print(out, indent + 1);
      if (i + 1 < declarations.size()) {
        out << "\n";
      }
      out << "\n";
    }
    printIndent(out, indent);
    out << "]";
  }
};

}  // namespace cminus

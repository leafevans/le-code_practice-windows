#include <fstream>
#include <iostream>
#include <sstream>
#include <stdexcept>
#include <string>
#include <vector>

#include "lexer.h"
#include "parser.h"

static std::string readWholeFile(const std::string& path) {
  std::ifstream in(path);
  if (!in) {
    throw std::runtime_error("cannot open input file: " + path);
  }
  std::ostringstream ss;
  ss << in.rdbuf();
  return ss.str();
}

int main(int argc, char** argv) {
  try {
    std::string source;

    if (argc >= 2) {
      source = readWholeFile(argv[1]);
    } else {
      source =
          "int x;\n"
          "int arr[10];\n"
          "int sum(int n) {\n"
          "  int i;\n"
          "  int s;\n"
          "  i = 0;\n"
          "  s = 0;\n"
          "  while (i < n) {\n"
          "    s = s + i;\n"
          "    i = i + 1;\n"
          "  }\n"
          "  return s;\n"
          "}\n";
    }

    cminus::Lexer lexer(source);
    std::vector<cminus::Token> tokens = lexer.tokenize();

    std::cout << "=== Token Stream ===\n";
    bool hasLexError = false;
    for (const auto& token : tokens) {
      std::cout << cminus::tokenTypeName(token.type) << "\t'" << token.lexeme
                << "'\t@" << token.line << ":" << token.column << "\n";
      if (token.type == cminus::TokenType::Error) {
        hasLexError = true;
      }
    }

    if (hasLexError) {
      std::cerr
          << "Lexical analysis failed: token stream contains ERROR tokens.\n";
      return 1;
    }

    cminus::Parser parser(tokens);
    auto program = parser.parseProgram();

    std::cout << "\n=== AST ===\n";
    program->print(std::cout, 0);
    std::cout << "\n";
    return 0;
  } catch (const std::exception& ex) {
    std::cerr << ex.what() << "\n";
    return 1;
  }
}

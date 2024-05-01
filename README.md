# Monkey Interpreter

## Summary

This is a interpreter written in Go that I wrote from a resource called "Writing an Interpreter in Go" by Thorsten Ball. It covers the process from:

    1. Lexical Analysis
    2. Parsing
    3. Walking the Abstract Syntax Tree
    4. Evaluating

## Lexical Analysis

In this process we transform the given source code into tokens. These tokens will then be separated into their respective types which could be a datatype such as an integer or a keyword that will later enact a process. This tokenizing stage will ignore all whitespace characters and look at each token to be prepared to be parsed for next stages.

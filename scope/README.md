# golangtutorials - Scope
* The scope of a predeclared identifier is the universe block.
* The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.
* The scope of the package name of an imported package is the file block of the file containing the import declaration.
* The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
* The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
* The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

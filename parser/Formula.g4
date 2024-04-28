grammar Formula;

equation: 


expr: OPENPAREN expr CLOSEPAREN
    | <assoc=right>expr POW expr
    | expr (MULT | DIV) expr
    | expr (PLUS | MINUS) expr
    | expr (EQUAL | NEQUAL | LESS | GREAT | LESSEQ | GREATEQ) expr
    | expr AND expr
    | expr OR expr
    | constant | variable;
    
constant: INTLIT | FLOATLIT;
binaryOperator: PLUS | MINUS | MULT | DIV | POW |
                EQUAL | NEQUAL | LESS | GREAT | LESSEQ | GREATEQ | AND | OR;

fraction: BACKSLASH 'frac' OPENCURLY ID CLOSECURLY OPENCURLY ID CLOSECURLY
mathSymbol: BACKSLASH ID mathSymbolTail | fraction
mathSymbolTail: 

variable: ID variableTail;
variableTail: SUBSCRIPT OPENCURLY  CLOSECURLY
        
OPENPAREN: '(';
CLOSEPAREN: ')';
OPENCURLY: '{';
CLOSECURLY: '}';
SUBSCRIPT: '_';
SUPERSCRIPT: '^';
BACKSLASH: '\';

PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';
POW: '**';
EQUAL: '==';
NEQUAL: '!=';
LESS: '<';
GREAT: '>';
LESSEQ: '<=';
GREATEQ: '>=';
AND: '&';
OR: '|';

EQUAL: '=';



SYMBOL: [a-z|A-Z];
ID: [a-z|A-Z]*;
LINE_COMMENT : '//' .*? '\r'? '\n' -> skip;
COMMENT: '/*' .*? '*/' -> skip;
INTLIT: [0]|[1-9][0-9]*;
FLOATLIT: INTLIT'.'[0-9]*;
WS: [ \t\r\n]+ -> skip;

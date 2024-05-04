grammar Formula;

equation: expr EQUAL expr;

expr: OPENPAREN expr CLOSEPAREN
    | <assoc=right>expr POW expr
    | expr (MULT | DIV) expr
    | expr (PLUS | MINUS) expr
    | constant | variable | fraction;
    
binaryOperator: PLUS | MINUS | MULT | DIV | POW | EQUAL ; 

fraction: '\\frac' OPENCURLY expr CLOSECURLY OPENCURLY expr CLOSECURLY;

constant: generalIntLit | FLOATLIT;

variable: generalId subscriptTail argumentTail;
subscriptTail: SUBSCRIPT OPENCURLY generalId CLOSECURLY | SUBSCRIPT OPENCURLY generalIntLit CLOSECURLY | SUBSCRIPT SINGLEID | SUBSCRIPT SINGLEINTLIT | ;
argumentTail: OPENPAREN argumentList CLOSEPAREN | OPENPAREN argumentList SEMICOLON argumentList CLOSEPAREN | ;

argumentList: expr argumentListTail;
argumentListTail: COMMA expr argumentListTail | ;

generalId: SINGLEID | ID;
generalIntLit: SINGLEINTLIT | INTLIT;

COMMA: ',';
DOT: '.';
SEMICOLON: ';';
OPENPAREN: '(';
CLOSEPAREN: ')';
OPENCURLY: '{';
CLOSECURLY: '}';
SUBSCRIPT: '_';
SUPERSCRIPT: '^';
BACKSLASH: '\\';

PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';
POW: '**';
EQUAL: '=';

SINGLEID: [a-z|A-Z];
ID: BACKSLASH [a-z|A-Z]+ | [a-z|A-Z][a-z|A-Z]+;
LINE_COMMENT: '//' .*? '\r'? '\n' -> skip;
COMMENT: '/*' .*? '*/' -> skip;

SINGLEINTLIT: [0-9];
INTLIT: [1-9][0-9]+;
FLOATLIT: INTLIT'.'[0-9]*;
WS: [ \t\r\n]+ -> skip;

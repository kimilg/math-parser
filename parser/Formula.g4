grammar Formula;

equation: expr EQUAL expr;

expr: OPENPAREN expr CLOSEPAREN
    | <assoc=right>expr POW expr
    | expr (MULT | DIV) expr
    | expr (PLUS | MINUS) expr
    | constant | variable | fraction;
    
binaryOperator: PLUS | MINUS | MULT | DIV | POW | EQUAL ; 

fraction: BACKSLASH 'frac' OPENCURLY expr CLOSECURLY OPENCURLY expr CLOSECURLY;

constant: INTLIT | FLOATLIT;

variable: ID subscriptTail argumentTail | BACKSLASH ID subscriptTail argumentTail ;
subscriptTail: SUBSCRIPT OPENCURLY ID CLOSECURLY | SUBSCRIPT OPENCURLY INTLIT CLOSECURLY | SUBSCRIPT ID | SUBSCRIPT INTLIT | ;
argumentTail: OPENPAREN argumentList CLOSEPAREN | OPENPAREN argumentList SEMICOLON argumentList CLOSEPAREN | ;

argumentList: expr argumentListTail;
argumentListTail: COMMA expr argumentListTail | ;


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


ID: [a-z|A-Z]+;
LINE_COMMENT: '//' .*? '\r'? '\n' -> skip;
COMMENT: '/*' .*? '*/' -> skip;
INTLIT: [0]|[1-9][0-9]*;
FLOATLIT: INTLIT'.'[0-9]*;
WS: [ \t\r\n]+ -> skip;

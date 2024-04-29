#!/bin/sh

alias antlr4='java -Xmx500M -cp "./antlr-4.13.1-complete.jar" org.antlr.v4.Tool'

antlr4 -Dlanguage=Go -no-visitor -package parsing -o ../parsing Formula.g4

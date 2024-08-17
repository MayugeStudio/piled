#!/usr/bin/env python3
from enum import Enum, auto
from typing import Optional

Location = tuple[int, int]


class TokenType(Enum):
    KEYWORD = auto()
    LITERAL = auto()


class Token:
    type: TokenType
    loc: Location

    def __init__(self, type: TokenType, loc: Location) -> None:
        self.type = type
        self.loc = loc


def lex_token(s: str) -> Optional[Token]:
    return None


def lex_line(line: str) -> list[Token]:
    return []


def lex_program(program: list[str]) -> list[Token]:
    return []


print("Hello, World")

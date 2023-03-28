"""
 SPDX-FileCopyrightText: 2023-present Intel Corporation

 SPDX-License-Identifier: Apache-2.0
"""

"""YANG usage guidelines plugin
This extends the validation by preventing go keywords to be used as identifiers
in YANG models.
"""

import optparse

from pyang import plugin
from pyang import statements
from pyang import error
from pyang import grammar

def pyang_plugin_init():
    plugin.register_plugin(GoPlugin())


class GoPlugin(plugin.PyangPlugin):
    def __init__(self):
        plugin.PyangPlugin.__init__(self)

    def setup_ctx(self, ctx):
        if not ctx.opts.lint:
            return
        self._setup_ctx(ctx)

    def _setup_ctx(self, ctx):
        "Should be called by any derived plugin's setup_ctx() function."

        ctx.strict = True
        ctx.canonical = True
        ctx.max_identifier_len = 64
        ctx.implicit_errors = False

        statements.add_validation_fun(
            'grammar', ['*'],
            lambda ctx, s: v_chk_go_keywords(ctx, s))

        error.add_error_code(
            'LINT_GO_KEYWORD', 3,
            '%s is a Go keyword, e.g. ' + ', '.join(_go_keywords))

        error.add_error_code(
            'LINT_GO_TYPE', 3,
            '%s is a Go type, e.g. ' + ', '.join(_go_types))

        error.add_error_code(
            'LINT_GO_CONSTANT', 3,
            '%s is a Go constant, e.g. ' + ', '.join(_go_constants))


_go_keywords = [
    "break",
    "case",
    "chan",
    "const",
    "continue",
    "default",
    "defer",
    "else",
    "fallthrough",
    "for",
    "func",
    "go",
    "goto",
    "if",
    "import",
    "interface",
    "map",
    "package",
    "range",
    "return",
    "select",
    "struct",
    "switch",
    "type",
    "var"
]

_go_types = [
    "bool",
    "byte",
    "complex64",
    "complex128",
    "float32",
    "float64",
    "int",
    "int8",
    "int16",
    "int32",
    "int64",
    "rune",
    "string",
    "uint",
    "uint8",
    "uint16",
    "uint32",
    "uint64",
    "uintptr",
]
_go_constants = [
    "false",
    "true"
]

def condition(substmt):
    return substmt.keyword == 'key'

def v_chk_go_keywords(ctx, stmt):
    if stmt.keyword in grammar.stmt_map:
        arg_type, subspec = grammar.stmt_map[stmt.keyword]

        """Find out if stmt is a list key - go to its parent and find child 'key'"""
        stmt_is_a_key = False
        if stmt.parent:
            keystmt = next(filter(condition, stmt.parent.substmts), None)
            if keystmt is not None and stmt.arg in keystmt.arg:
                stmt_is_a_key = True

        if arg_type in 'identifier' and stmt_is_a_key and not_go_keyword(stmt.arg):
            """only reject go keyword in list keys"""
            error.err_add(ctx.errors, stmt.pos, 'LINT_GO_KEYWORD', stmt.arg)
        elif arg_type in 'identifier' and not_go_type(stmt.arg):
            error.err_add(ctx.errors, stmt.pos, 'LINT_GO_TYPE', stmt.arg)
        elif arg_type in 'identifier' and not_go_constant(stmt.arg):
            error.err_add(ctx.errors, stmt.pos, 'LINT_GO_CONSTANT', stmt.arg)


def not_go_keyword(name: str):
    """Returns True if is not keyword"""
    if name is None:
        return False
    return name.lower() in _go_keywords

def not_go_type(name: str):
    """Returns True if is not type"""
    if name is None:
        return False
    return name.lower() in _go_types

def not_go_constant(name: str):
    """Returns True if is not constant"""
    if name is None:
        return False
    return name.lower() in _go_constants

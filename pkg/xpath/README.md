# XPath support with YANG

[XPath 1.0] is a referenced standard as part of the [YANG] specification. It is
used as a query language for specifying many inter-node references and
dependencies.

To bring XPath support to [YGOT], this project uses a Go XPath implementation
from [Antchfx]. Specifically it implements the [NodeNavigator] interface, as
`YangNodeNavigator` thereby allowing it to reuse the `Select()` and `Evaluate()`
methods on parsed XPath statements.

The implementation consists of 3 major parts:

## Creating the YangNodeNavigator
A `YangNodeNavigator` is created by merging together the YANG `Schema` with 
a `ValidatedGoStruct`.

The `Schema` (created by YGOT) is a tree of `yang.Entry`s reflecting the
hierarchy of the YANG model. In building the `YangNodeNavigator`, a recursive
function steps through every part of the tree, adding new `Annotation`
entries to hold the related Go Struct.

List entries are treated specially with each instance of a list, being given
a new `Dir` entry in the tree with the name format `<listname>__<listindex>`.

> To ensure the consistent navigation of the tree, a slice of `orderedkeys` is
> added to the `yang.Entry` as an `Annotation`. If a leaf has no value, it will
> be omitted from the `orderedkeys` list. 

## Querying the YangNodeNavigator
Once created, the `YangNodeNavigator` can be traversed using the interface
methods (from `NodeNavigator`) like `MoveToChild()`, `MoveToNext()` etc.

> The `orderedkeys` list ensures a consistent order of traversal.

XPath queries (once `Compile()`ed) can be executed on the `YangNodenavigator` in
2 ways:

1) `Evaluate()` - returns the result of the expression - this can handle XPath
    functions like `count`, `sum` etc. and can return scalar values (number,
    boolean, string) or a NodeNavigator.
2) `Select()` - selects a node set using the specified XPath expression - this
    node set can be then be navigated to extract values.

There are many examples of both types of query in the
[unit test](../../modelplugin/testdevice-2.0.0/testdevice_2_0_0/xpath_test.go).

For example
```
count(/t1:cont1a/t1:list2a[number(t1:tx-power) < number(t1:rx-power)])
```
This can be interpreted as: count the number of instances of `list2a` whose
attribute `tx-power` is less than its `rx-power`.

> Each attribute and container should be prefixed with its namespace prefix.
> Also any comparison will be a text comparison by default, and hence numbers
> must be cast explicitly.

Another example is:
```
sum(/t1:cont1a/t1:list2a[contains(@t1:name,'l2')]/t1:rx-power)
```
This can be interpreted as: add together all of the `rx-power` values from
`list2a` entries that have a `name` attribute that contains the substring `l2`.

> Because `name` is the key of the `list2a` list in YANG, it must be prefixed
> with `@` when referring to it. Also text strings should be wrapped in
> single quotes. Predicates are given in `[]` and any number can be given,
> allowing querying of double keyed lists.

## Evaluating 'must' statements in YANG model
One of the main use cases for XPath within YANG is the `must` statement.
This is specified as an XPath query that must evaluate to `true`.

To validate a `ValidatedGoStruct`, it is transformed to a `YangNodeNavigator`
and the `WalkAndValidateMust()` method can then be used to iterate through
the tree, finding `must` statements and `Evaluate()`ing them. If any `must`
statement does not evaluate to `true`, then an error is thrown and the
validation stops.

An example of a must statement, is given below as a statement on a `container`
that enforces some relationship between leafs in that container, or its
parent, peers or children:
```
container cont2a {
  must "string(t1:leaf2g) = 'true' or number(./t1:leaf2a) < 4" {
    error-message "If Leaf 2a is above 4 then leaf2g must be true for the validation to pass";
  }
  ...
```

This can be interpreted as:
either `leaf2g` (which is a boolean) must be `true` or else the value of `leaf2a`
must be less than `4`. When this is not the case, the `must` validation will fail,
as the configuration is not valid.

> Care must be taken with data types when doing comparisons - above the `leaf2a`
> is a number and must be converted before comparison with the number 4. Also
> comparing a boolean leaf (leaf2g) is best done by conversion to a string
> before comparison.

In the above XPath query, the relative operator `.` is used to refer to the
container context - it could be omitted from the `./t1:leaf2a`, as it is assumed
by default (as happens in `t1:leaf2g`). Similarly

* The `..` operator can be used to refer to the parent of the container
* The `/` refers to the root of the tree
* The `//` refers to a child at any level beneath the root


[XPath 1.0]: https://www.w3.org/TR/1999/REC-xpath-19991116/
[YANG]: https://datatracker.ietf.org/doc/html/rfc6020#section-6.4
[YGOT]: https://github.com/openconfig/ygot
[Antchfx]: github.com/antchfx/xpath
[NodeNavigator]: https://github.com/antchfx/xpath/blob/696d1234f878e2c59321bb58cbc838250b1191e0/xpath.go#L32

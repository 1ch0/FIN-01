package stdlib

import "list"

l1: [1, 2, 3, 4, 5, 6]
l2: ["c", "b", "a"]

// constrain.cue length
l2: list.MinItems(1)
l2: list.MaxItems(3)

// slice a lists.cue into 左闭右开
l3: list.Slice(l1, 2, 4)

// get the sum and product
sum: list.Sum(l1)
prd: list.Product(l1)
// a: 1*2*3*4*5*6

// liner search for lists.cue (no binary)
lc: list.Contains(l1, 2)

// sort the lists.cue
ls:  list.Sort(l2, list.Ascending)
l2s: list.IsSorted(l2, list.Ascending)
lss: list.IsSorted(ls, list.Ascending)

// Flatten a lists.cue
ll: [1, [2, 3], [4, [5]]]
lf: list.FlattenN(ll, 1)

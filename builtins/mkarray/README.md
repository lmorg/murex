# Builtins: make array

This provides the array builtin, `a` and `ja` (returns array in JSON):

    a: [monday..friday]
    a: [Jan..Dec]
    a: [0..99]
    a: [00..09]
    a: [z..a]
    a: [0000..1000x2]
    a: [00..ffx16]
    a: [foo,bar]
    a: foo[a,b,c..x,y,z]bar
    a: [foo,bar][bar,foo]
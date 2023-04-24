# Builtins: HCL

This provides support for HCL data format: https://github.com/hashicorp/hcl

It is an optional builtin and has an additional dependency:

    go get -u github.com/hashicorp/hcl

Because the above library doesn't support marshalling and because JSON
is fully compatible with HCL, Murex will default to marshalling HCL
data types as JSON.
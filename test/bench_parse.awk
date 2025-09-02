#!/usr/bin/awk -f
# Parse `go test -bench` output into CSV with derived metrics.
# Outputs: suite,name,type,size,parallel,mode,ns_per_op,bytes_per_op,allocs_per_op,ns_per_item,items_per_sec

BEGIN {
    FS=" "; OFS=",";
    print "suite","name","type","size","parallel","mode","ns_per_op","bytes_per_op","allocs_per_op","ns_per_item","items_per_sec";
}

/^BenchmarkForeachParallel(Matrix|ExecTrue)\// {
    line=$0;
    # First token contains the benchmark name with sub-benchmark path
    name=$1;
    # Extract metrics by searching tokens
    ns=bytes=allocs="";
    for (i=1; i<=NF; i++) {
        if ($i=="ns/op" && i>1) ns=$(i-1);
        if ($i=="B/op" && i>1) bytes=$(i-1);
        if ($i=="allocs/op" && i>1) allocs=$(i-1);
    }

    # Strip leading 'Benchmark' and trailing thread suffix '-8'
    sub(/^Benchmark/, "", name);
    sub(/-[0-9]+$/, "", name);

    # Split path segments
    n=split(name, seg, "/");
    suite=seg[1];
    benchType="foreach";
    idx=2;
    if (suite ~ /ExecTrue$/) {
        # ExecTrue format includes an extra 'exec' segment
        if (seg[2] == "exec") { benchType="exec"; idx=3; } else { benchType="exec"; }
    }

    size=""; par=""; mode="";
    for (j=idx; j<=n; j++) {
        if (seg[j] ~ /^size=/) { sub(/^size=/, "", seg[j]); size=seg[j]; }
        else if (seg[j] ~ /^parallel=/) { sub(/^parallel=/, "", seg[j]); par=seg[j]; }
        else if (seg[j] ~ /^(ordered|unordered)$/) { mode=seg[j]; }
    }

    # Derive per-item metrics when size is known and numeric
    ns_item=""; items_sec="";
    if (ns != "" && size ~ /^[0-9]+$/) {
        ns_item = ns / size;
        if (ns_item > 0) items_sec = 1e9 / ns_item; else items_sec="";
    }

    print suite,name,benchType,size,par,mode,ns,bytes,allocs,ns_item,items_sec;
}

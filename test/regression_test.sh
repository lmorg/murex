#!/usr/bin/env bash

# Regression tests
cd test
JSH="../murex"
timeout="1.1"
nreps=250
repstimeout="15"

i=1
failed=0
TTY=$(tty)

NC='\033[0m'
PASSED="\033[0;32m[PASSED]$NC"
UNEXPECTED="\033[0;31m[UNEXPECTED OUTPUT]$NC"
ONLYMANAGED="\033[0;31m[ONLY MANAGED"
TIMEDOUT="\033[0;35m[TIMED OUT]$NC"
DESC="\033[0;34m"

trap ctrl_c INT

if [[ $(which timeout >/dev/null 2>&1; echo $?) != 0 ]]; then
    printf "\033[0;31m$0 requires 'timeout' installed!$NC\n"
    exit 1
fi

ctrl_c() {
        printf "\n\033[0;31m[TESTS TERMINATED BY USER]$NC\n"
        exit 1
}

shell() {
    printf "[%3d] $DESC%-50s$NC " $i "$1" >$TTY
    timeout $timeout $JSH -c "$1"
    if [ $? -eq 124 ]; then
        #echo "timed out"
        printf "$TIMEDOUT " >$TTY
    fi
}

check() {
    in=$(cat)
    comp="$(echo -e "$1")"
    if [[ "$in" == "$comp" ]]; then
        printf "$PASSED\n"
    else
        printf "$UNEXPECTED\n"
        return 1
    fi
}

reps() {
    printf "[%3d] $DESC%-50s $NC" $i "$2 REPS: $1" >$TTY
    timeout $repstimeout bash -c "( for i in {1..$2}; do echo -n '->';	$JSH -c '$1'; done ) | wc -l | sed -r 's/ //g'"
    if [ $? -eq 124 ]; then
        printf "$TIMEDOUT " >$TTY
    fi
}

checkreps() {
    in=$(cat)
    if [[ "$in" == "$1" ]]; then
        printf "$PASSED\n"
    else
        printf "$ONLYMANAGED $in]$NC\n"
        return 1
    fi
}

while true; do
    case $i in
        # testing basic syntax with out and err streams
        0)shell '' 2>&1                             | check "";;
        1)shell 'null' 2>&1                         | check "";;
        2)shell '# out: out' 2>&1                   | check "";;
        3)shell 'out: out # comment' 2>&1           | check "out";;
        4)shell 'out: out; err: err' 2>/dev/null     | check "out";;
        5)shell 'out: out; err: err' 2>&1 >/dev/null | check "err";;
        6)shell 'out: out; err: err' 2>&1            | check "out\nerr";;
        7)shell 'out: o u t' 2>&1   | check "o u t";;
        8)shell 'err: o u t' 2>&1   | check "o u t";;
        9)shell 'out: "o u t"' 2>&1 | check "o u t";;
        10)shell 'err: "o u t"' 2>&1 | check "o u t";;
        11)shell "out: 'o u t'" 2>&1 | check "o u t";;
        12)shell "err: 'o u t'" 2>&1 | check "o u t";;
        13)shell 'out: `o u t`' 2>&1 | check '`o u t`';;
        14)shell 'err: `o u t`' 2>&1 | check '`o u t`';;
        15)shell 'out: o,u,t' 2>&1   | check "o,u,t";;
        16)shell 'err: o,u,t' 2>&1   | check "o,u,t";;
        17)shell 'out: "o,u,t"' 2>&1 | check "o,u,t";;
        18)shell 'err: "o,u,t"' 2>&1 | check "o,u,t";;
        19)shell "out: 'o,u,t'" 2>&1 | check "o,u,t";;
        20)shell "err: 'o,u,t'" 2>&1 | check "o,u,t";;
        21)shell 'out: `o,u,t`' 2>&1 | check '`o,u,t`';;
        22)shell 'err: `o,u,t`' 2>&1 | check '`o,u,t`';;

        # testing out and err pipes
        #23)shell 'printf: out\n' 2>&1                  | check "out";;                      # `printf` without a pty sends \n
        #24)shell 'printf: out1\n; printf: out2\n' 2>&1    | check "echo -e 'out1\nout2";;   # `printf` without a pty sends \n
        23)shell 'printf: out\n' 2>&1                  | check "out\r";;                     # `printf` with a pty sends \r\n
        24)shell 'printf: out1\n; printf: out2\n' 2>&1 | check "out1\r\nout2\r";;            # `printf` with a pty sends \r\n
        25)shell 'printf: out\n | grep: out' 2>&1      | check "out";;
        26)shell 'out: out | grep: out' 2>&1         | check "out";;
        27)shell 'err: err | grep: err' 2>/dev/null  | check "";;
        28)shell 'err: err | grep: err' 2>&1         | check "err";;
        29)shell 'err: err | grep: out' 2>&1         | check "err";;
        30)shell 'sleep: 60; out: awake # this should timeout' 2>&1 | check "";;
        31)shell 'sleep: 1; out: awake' 2>&1          | check "awake";;
        32)shell 'sleep: 1 | out: awake' 2>&1        | check "awake";;
        33)shell 'sleep: 60 | out: awake # this should timeout' 2>&1 | check "awake";;

        # more complex testing of pipes
        34)shell 'out: out->match: out' 2>&1                     | check "out";;
        35)shell 'out: aaout->match: out' 2>&1                   | check "aaout";;
        36)shell 'out: outaa->match: out' 2>&1                   | check "outaa";;
        37)shell 'out: aaoutaa->match: out' 2>&1                 | check "aaoutaa";;
        38)shell 'out: out->match: out; err: err' 2>/dev/null    | check "out";;
        39)shell 'out: out->match: out; err: err' 2>&1 >/dev/null | check "err";;
        40)shell 'out: out->match: out; err: err->match: out' 2>/dev/null | check "out";;
        41)shell 'out: out->match: out; err: err->match: out' 2>&1 | check "Methodable function \`match\` does not exist for \`err.(null)\`\nout\nerr";;
        42)shell 'out: out->match: out; err: err' 2>&1            | check "out\nerr";;
        43)shell 'out: out1->match: out1; out: out2->match: out2' 2>&1 | check "out1\nout2";;
        44)shell 'out: out->match: noout' 2>&1                   | check "";;
        45)shell 'out: out->match: noout' 2>/dev/null            | check "";;

        # more complex testing of pipes originating from PTYs
        46)shell 'printf: out\n->match: out' 2>&1                     | check "out";;
        47)shell 'printf: aaout\n->match: out' 2>&1                   | check "aaout";;
        48)shell 'printf: outaa\n->match: out' 2>&1                   | check "outaa";;
        49)shell 'printf: aaoutaa\n->match: out' 2>&1                 | check "aaoutaa";;
        50)shell 'printf: out\n->match: out; err: err' 2>/dev/null     | check "out";;
        51)shell 'printf: out\n->match: out; err: err' 2>&1 >/dev/null | check "err";;
        52)shell 'printf: out\n->match: out; err: err->match: out' 2>/dev/null | check "out";;
        53)shell 'printf: out\n->match: out; err: err' 2>&1            | check "out\nerr";;
        54)shell 'printf:out1\n->match:out1;printf:out2\n->match:out2' 2>&1 | check "out1\nout2";;
        55)shell 'printf: out\n->match: noout' 2>&1                   | check "";;
        56)shell 'printf: out\n->match: noout' 2>/dev/null            | check "";;

        # testing err pipes
        57)shell 'out: out->match: out; err: err; out: 2' 2>&1        | check "out\nerr\n2";;
        58)shell 'out: out ? grep: out' 2>&1 >/dev/null | check "";;
        59)shell 'err: err ? grep: err' 2>/dev/null     | check "err";;
        60)shell 'err: err ? grep: err' >/dev/null      | check "";;
        61)shell 'err: err ? grep: out' 2>&1            | check "";;
        62)shell 'sleep: 1 ? out: awake' 2>&1           | check "awake";;
        63)shell 'sleep: 60 ? out: awake # this should timeout' 2>&1 | check "awake";;

        # test `match` without ending \n
        64)shell 'text: fox.txt->match: jumped' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        65)shell 'text: fox.txt->match: jumped' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        66)shell 'text: fox.txt->match' 2>&1 | check 'Error in `match`: No parameters supplied.';;
        67)shell 'text: fox.txt->!match:e->!match:o' 2>&1| check "quick\nlazy";;
        68)shell 'text: fox.txt->!match' 2>&1| check 'Error in `!match`: No parameters supplied.';;

        # test `regex` without ending \n
        69)shell 'text: fox.txt->regex: m,jumped' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        70)shell 'text: fox.txt->regex: m,jumped' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        71)shell 'text: fox.txt->regex: `m,jumped`' 2>/dev/null | check "";;
        72)shell 'text: fox.txt->regex: `m,jumped`' 2>&1| check 'Error in `regex`: Invalid regexp.';;
        73)shell 'text: fox.txt->regex: m jumped' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        74)shell 'text: fox.txt->regex: m jumped' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        75)shell 'text: fox.txt->regex: "m,jumped"' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        76)shell 'text: fox.txt->regex: "m,jumped"' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        77)shell 'text: fox.txt->regex: m/jumped/' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        78)shell 'text: fox.txt->regex: m/jumped/' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        79)shell 'text: fox.txt->regex: "m#jumped#"' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        80)shell 'text: fox.txt->regex: "m#jumped#"' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        81)shell 'text: fox.txt->regex: ' 2>&1 | check 'Error in `regex`: No parameters supplied.';;
        82)shell 'text: fox.txt->!regex: m,[eo]' 2>&1| check "quick\nlazy";;
        83)shell 'text: fox.txt->!regex: ' 2>&1| check 'Error in `!regex`: No parameters supplied.';;
        84)shell 'text: fox.txt->!regex: m,[eo]->regex: s/[ai]/x/' 2>&1| check "quxck\nlxzy";;

        # test `match` with ending \n
        85)shell 'text: fox_crlf.txt->match: jumped' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        86)shell 'text: fox_crlf.txt->match: jumped' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        87)shell 'text: fox_crlf.txt->match: ' 2>&1 | check 'Error in `match`: No parameters supplied.';;
        88)shell 'text: fox_crlf.txt->!match: e->!match: o' 2>&1| check "quick\nlazy";;
        89)shell 'text: fox_crlf.txt->!match: ' 2>&1| check 'Error in `!match`: No parameters supplied.';;

        # test `regex` with ending \n
        90)shell 'text: fox_crlf.txt->regex: m,jumped' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        91)shell 'text: fox_crlf.txt->regex: m,jumped' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        92)shell 'text: fox_crlf.txt->regex: "m,jumped"' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        93)shell 'text: fox_crlf.txt->regex: "m,jumped"' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        94)shell 'text: fox_crlf.txt->regex: m jumped' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        95)shell 'text: fox_crlf.txt->regex: m jumped' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        96)shell 'text: fox_crlf.txt->regex: "m,jumped"' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        97)shell 'text: fox_crlf.txt->regex: "m,jumped"' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        98)shell 'text: fox_crlf.txt->regex: m/jumped/' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        99)shell 'text: fox_crlf.txt->regex: m/jumped/' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        100)shell 'text: fox_crlf.txt->regex: "m#jumped#"' 2>/dev/null | check "fox jumped over\njumped over\nfox jumped\njumped";;
        101)shell 'text: fox_crlf.txt->regex: "m#jumped#"' 2>&1| check "fox jumped over\njumped over\nfox jumped\njumped";;
        102)shell 'text: fox_crlf.txt->regex: ' 2>&1 | check 'Error in `regex`: No parameters supplied.';;
        103)shell 'text: fox_crlf.txt->!regex: m,[eo]' 2>&1| check "quick\nlazy";;
        104)shell 'text: fox_crlf.txt->!regex: ' 2>&1| check 'Error in `!regex`: No parameters supplied.';;
        105)shell 'text:fox_crlf.txt->!regex:m,[eo]->regex:s/[ai]/x/' 2>&1| check "quxck\nlxzy";;

        # test whitespaces
        106)shell 'out: out|grep: out' 2>&1 | check "out";;
        107)shell 'out: out | grep: out' 2>&1 | check "out";;
        108)shell 'out: out  |  grep: out' 2>&1 | check "out";;
        109)shell "$(echo -e 'out: out\t|\tgrep: out')" 2>&1 | check "out";;
        110)shell "$(echo -e 'out: out\r|\ngrep: out')" 2>&1 | check "out";;

        # test builtins don't get called when executing a PTY
        111)shell 'out: out|match: out' 2>&1 | check 'exec: "match": executable file not found in $PATH';;

        # test blocking in concurrent pipes
        112)shell 'exec: printf out\n' 2>&1 | check "out";;
        113)shell 'exec: sleep 5; out: awake # this should timeout' 2>&1 | check "";;
        114)shell 'exec: sleep 1;  out: awake' 2>&1          | check "awake";;
        115)shell 'exec: sleep 5 | out: awake # this should timeout' 2>&1        | check "awake";;
        116)shell 'exec: sleep 1  | out: awake' 2>&1        | check "awake";;
        117)shell 'exec: sh -c "sleep 5; echo out" | out: awake # this should timeout' 2>&1 | check "awake";;
        118)shell 'exec: sh -c "sleep 1; echo out 1>&2" | out: awake' 2>&1                       | check "awake\nout";;
        119)shell 'exec: sh -c "sleep 5; echo out" | grep: out # this should timeout' 2>&1  | check "";;
        120)shell 'exec: sh -c "sleep 5; echo out"->match: out # this should timeout' 2>&1   | check "";;

        # testing `get`
        121)shell 'get: http://laurencemorgan.co.uk->json: Status Code' 2>&1 | check "200";;

        # repetition tests for consistency
        122)reps 'out: out | grep: out->match: out' $nreps 2>&1 | checkreps $nreps;;
        123)reps 'printf: out\n->match: out' $nreps 2>&1 | checkreps $nreps;;
        124)reps 'out: out->match: out' $nreps 2>&1 | checkreps $nreps;;
        125)reps 'text: fox.txt->foreach: line {out: $line} ->match: "This is just some dummy text for regression testing"' $nreps 2>&1 | checkreps $nreps;;
        126)reps 'text: fox_crlf.txt->foreach: line {out: $line} ->match: "This is just some dummy text for regression testing"' $nreps 2>&1 | checkreps $nreps;;

        # checking booleans and `if` structures
        127)shell 'out: true->if: {out: match}' 2>&1 | check "match";;
        128)shell 'out: true->!if: {out: match}' 2>&1 | check "";;
        129)shell 'out: true->if: {out: true} {out: match}' 2>&1 | check "true";;
        130)shell 'out: true->!if: {out: false} {out: match}' 2>&1 | check "match";;
        131)shell 'if: {out: false} {out: match}' 2>&1 | check "";;
        132)shell '!if: {out: false} {out: match}' 2>&1 | check "match";;
        133)shell 'if: {out: true} {out: match}' 2>&1 | check "match";;
        134)shell '!if: {out: true} {out: match}' 2>&1 | check "";;
        135)shell 'if: {out: true} {out: positive} {out: negative}' 2>&1 | check "positive";;
        136)shell '!if: {out: true} {out: positive} {out: negative}' 2>&1 | check "negative";;
        137)shell 'if: {out: false} {out: positive} {out: negative}' 2>&1 | check "negative";;
        138)shell '!if: {out: false} {out: positive} {out: negative}' 2>&1 | check "positive";;
        139)shell 'exec: true->if: {out: match} # returns false because no stdout' 2>&1 | check "";;
        140)shell 'exec: true->!if: {out: match}# returns false because no stdout' 2>&1 | check "match";;
        141)shell 'exec: false->if: {out: match}' 2>&1 | check "";;
        142)shell 'exec: false->!if: {out: match}' 2>&1 | check "match";;
        143)shell 'out: qwerty->if: {out: match}' 2>&1 | check "match";;
        144)shell 'out: qwerty->!if: {out: match}' 2>&1 | check "";;
        145)shell 'out: ->if: {out: match}' 2>&1 | check "";;
        146)shell 'out: ->!if: {out: match}' 2>&1 | check "match";;
        147)shell 'out: 1->if: {out: match}' 2>&1 | check "match";;
        148)shell 'out: 1->!if: {out: match}' 2>&1 | check "";;
        149)shell 'out: 0->if: {out: match}' 2>&1 | check "";;
        150)shell 'out: 0->!if: {out: match}' 2>&1 | check "match";;
        151)shell 'out: yes->if: {out: match}' 2>&1 | check "match";;
        152)shell 'out: yes->!if: {out: match}' 2>&1 | check "";;
        153)shell 'out: no->if: {out: match}' 2>&1 | check "";;
        154)shell 'out: no->!if: {out: match}' 2>&1 | check "match";;
        155)shell 'out: on->if: {out: match}' 2>&1 | check "match";;
        156)shell 'out: on->!if: {out: match}' 2>&1 | check "";;
        157)shell 'out: off->if: {out: match}' 2>&1 | check "";;
        158)shell 'out: off->!if: {out: match}' 2>&1 | check "match";;
        159)shell 'out: pass->if: {out: match}' 2>&1 | check "match";;
        160)shell 'out: pass->!if: {out: match}' 2>&1 | check "";;
        161)shell 'out: fail->if: {out: match}' 2>&1 | check "";;
        162)shell 'out: fail->!if: {out: match}' 2>&1 | check "match";;
        163)shell 'out: passed->if: {out: match}' 2>&1 | check "match";;
        164)shell 'out: passed->!if: {out: match}' 2>&1 | check "";;
        165)shell 'out: failed->if: {out: match}' 2>&1 | check "";;
        166)shell 'out: failed->!if: {out: match}' 2>&1 | check "match";;
        167)shell 'true ->if: {out: match}' 2>&1 | check "match";;
        168)shell 'true ->!if: {out: match}' 2>&1 | check "";;
        169)shell 'false ->if: {out: match}' 2>&1 | check "";;
        170)shell 'false ->!if: {out: match}' 2>&1 | check "match";;
        171)shell 'false ->if: {out: match}' 2>&1 | check "";;
        172)shell 'false ->!if: {out: match}' 2>&1 | check "match";;
        173)shell 'true ->!' 2>&1 | check "false";;
        174)shell 'false ->!' 2>&1 | check "true";;

        # testing `try` and `catch` structures
        175)shell 'true         ->catch  {out: catch}' 2>&1 | check "true";;
        176)shell 'false        ->catch  {out: catch}' 2>&1 | check "false\ncatch";;
        177)shell 'true         ->!catch {out: catch}' 2>&1 | check "true\ncatch";;
        178)shell 'false        ->!catch {out: catch}' 2>&1 | check "false";;
        179)shell 'try { true  }->catch  {out: catch}' 2>&1 | check "true";;
        180)shell 'try { false }->catch  {out: catch}' 2>&1 | check "false\ncatch";;
        181)shell 'try { true  }->!catch {out: catch}' 2>&1 | check "true\ncatch";;
        182)shell 'try { false }->!catch {out: catch}' 2>&1 | check "false";;

        # testing `try` and `catch` structures with secondary `catch` for else conditions
        183)shell 'true         ->catch {out:catch}->!catch{out:else}' 2>&1 | check "true\nelse";;
        184)shell 'false        ->catch {out:catch}->!catch{out:else}' 2>&1 | check "false\ncatch";;
        185)shell 'true         ->!catch{out:catch}->catch {out:else}' 2>&1 | check "true\ncatch";;
        186)shell 'false        ->!catch{out:catch}->catch {out:else}' 2>&1 | check "false\nelse";;
        187)shell 'try { true  }->catch {out:catch}->!catch{out:else}' 2>&1 | check "true\nelse";;
        188)shell 'try { false }->catch {out:catch}->!catch{out:else}' 2>&1 | check "false\ncatch";;
        189)shell 'try { true  }->!catch{out:catch}->catch {out:else}' 2>&1 | check "true\ncatch";;
        190)shell 'try { false }->!catch{out:catch}->catch {out:else}' 2>&1 | check "false\nelse";;

        # testing new regex function: find
        191)shell 'text: fox_crlf.txt->regex: f/fox/' | check "fox\nfox\nfox\nfox";;

        # some more advanced testing of `try` and `catch`
        192)shell 'try: { out: 1; out: 2 | grep: false; out: 3 }' 2>&1 | check "1";;
        193)shell 'try{out:1;out:2|grep:no; out:3}->catch{out:failed}' 2>&1 | check "1\nfailed";;
        194)shell 'try: { out: 1; out: 2 | grep: 2; out: 3 }' 2>&1 | check "1\n2\n3";;
        195)shell 'try {out 1;out 2|grep 2;out 3}->catch{out failed}' 2>&1 | check "1\n2\n3";;
        196)shell 'try {out 1;out 2|grep 2;out 3}->!catch{out failed}' 2>&1 | check "1\n2\n3\nfailed";;
        197)shell 'out:1;out:2|grep:false->catch{out:failed}' 2>&1 | check '1\nfailed';;
        198)shell 'out:1;out:2|grep:2->catch{out:failed}' 2>&1 | check '1\n2';;
        199)shell 'out:1;out:2|grep:false->!catch{out:success}' 2>&1 | check '1';;
        200)shell 'out:1;out:2|grep:2->!catch{out:success}' 2>&1 | check '1\n2\nsuccess';;

        # testing encoder builtins
        201)shell 'out: test->base64->!base64' 2>&1 | check "test\n";;
        202)reps 'out: test->base64->!base64->match: test' $nreps 2>&1 | checkreps $nreps;;
        203)reps 'out: test\n->base64->!base64->match:test' $nreps 2>&1 | checkreps $nreps;;
        204)shell 'out: test->escape->!escape' 2>&1 | check "test";;
        205)shell 'out: test->gz->!gz' 2>&1         | check "test";;

        # variables and subshells
        206)shell 'set foo=bar; out $foo' 2>&1      | check "bar";;
        207)shell 'set foo=bar; out "$foo"' 2>&1    | check "bar";;
        208)shell "set foo=bar; out '\$foo'" 2>&1   | check '$foo';;
        209)shell 'set foo=bar; out "\$foo"' 2>&1   | check '$foo';;
        210)shell 'out bar->set foo; out $foo' 2>&1      | check "bar";;
        211)shell 'out bar->set foo; out "$foo"' 2>&1    | check "bar";;
        212)shell "out bar->set foo; out '\$foo'" 2>&1   | check '$foo';;
        213)shell 'out bar->set foo; out "\$foo"' 2>&1   | check '$foo';;
        214)shell 'out ${out foo}' 2>&1      | check "foo";;
        215)shell 'out "${out foo}"' 2>&1    | check "foo";;
        216)shell "out '\${out foo}'" 2>&1   | check '${out foo}';;
        217)shell 'out "\${out foo}"' 2>&1   | check '${out foo}';;

        *) break
    esac
    let failed=$failed+$?
    let i++
done

echo -e "\nAll tests have been run. $failed failed."

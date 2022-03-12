# Proposal: ModSecurity compatible web application firewall plugin for Apache APISIX

## Background

https://github.com/apache/apisix/issues/1868

## ModSecurity and the new Coraza WAF

## Execution flow

## OWASP Core Ruleset

## wasm

https://github.com/tinygo-org/tinygo/issues/447

```
 > [5/5] RUN tinygo build -o ./main.go.wasm -scheduler=none -target=wasi ./main.go:
#9 0.907 # encoding/xml
#9 0.907 /usr/local/go/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
#9 0.907 /usr/local/go/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
```

## References

- [Talking about ModSecurity and the new Coraza WAF](https://coreruleset.org/20211222/talking-about-modsecurity-and-the-new-coraza-waf/)
- [Tinygo need serialization tools #447](https://github.com/tinygo-org/tinygo/issues/447)

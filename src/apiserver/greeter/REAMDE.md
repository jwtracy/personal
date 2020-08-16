# Greeter Service

Test with
```shell
grpcurl -protoset src/apiserver/greeter/pb/greeter.pb -plaintext -d '{"name": "John"}' 159.89.222.167:8080 app.greeter.Greeter/Hello
```

or

```shell
curl -X POST -H 'Content-type:application/json' -d '{"name": "John"}' 68.183.249.43/twirp/app.greeter.Greeter/Hello | jq
```


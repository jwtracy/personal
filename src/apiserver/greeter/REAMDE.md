# Greeter Service

Test with
```shell
grpcurl -protoset src/apiserver/greeter/pb/greeter.pb -plaintext -d '{"name": "John"}' 159.89.222.167:8080 app.greeter.Greeter/Hello
```

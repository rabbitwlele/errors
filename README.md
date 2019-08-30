## 生成一个error

和原生error方法相同，会自动记录stack

```go
err :=  errors.New()
```

## 给error添加stack或者message

场景一般是已经接受一个error，需要保存栈信息

```go
//adding stack to an error
err := io.EOF
err = errors.WithStack(err)
err = errors.WithMessage(err)
```

### 获取根因

会获取生成链最底端的error

```go
switch err := errors.Cause(err) {
case io.EOF:
        // handle specifically
default:
        // unknown error
}
```

## 给error追加code信息

### 创建code

```go
import 'github.com/rabbitwlele/errors/code'
UserNotFound := code.New(1001,"user not found")
RecordNotFound := code.New(1002,"record not found")  
```

### 追加和获取code

```go
err = errors.WithCode(err,UserNotFound)
fmt.Pritntln(errors.Code(err).Code() , errors.Code().Message()) //1001 user not found
err = errors.WithCode(err,RecordNotFound)
fmt.Pritntln(errors.Code(err).Code() , errors.Code().Message()) //1002 record not found
```

## 建议

1. 第三方包返回的error包装stack和code
2. 内部的包不直接不做任何包装
3. 被调用者函数内部不做任何error打印日志处理，将error抛给调用者
4. 错误日志统一由最上层的调用者打印
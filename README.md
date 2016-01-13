
##说明
StackErr记录错误发生时filename,line,stack等信息

##使用方法
stackerr.New(err interface{}) *StackErr

StackErr.Error() 错误的描述信息
StackErr.Detail() 会返回带有文件名及文件函数的错误描述信息 
StackErr.Stack() 会返回错误发生时的堆栈信息

# Lec7a 元循环求值器  

7a,7b主要讲的就是一个解释器了。
- 解释器特性；
- 元解释器有什么用-探究语言特性；

```
      PROC,ARGS
       ----->
EVAL           APPLY
       <-----
      EXP,ENV

即：
EXP,ENV   --> EVAL   --> PROC,ARGS (EVAL就是将表达式+上下文环境 转化为 可执行函数与参数)
PROC,ARGS --> APPLY  --> EXP,ENV (APPLY将可以计算的基本过程计算，然后又得到不能算的表达式了)
```

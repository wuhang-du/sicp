# 通用运算符  

这篇是很精彩的。  

1. 首先基于一个复数的系统设计加减计算。  
George: a+bi Mary:a+cosb 即：两个人使用了两种不同的方式去表示复数。  
同时两个人又分别给出了自己的函数组，如：获取实部，获取虚部，等。  
2. 为了方便识别，又分别加上了类型的东西。  
3. 做加减法时，由经理去判断类型，然后调用对应的函数。  

那为什么要有经理呢？如何优化？  

3. 有一个公共的函数表，存储着 类型-运算符-对应的方法，例如： George_Rect-> '+ -> func(recta,rectb)。  
George和Mary都去那里注册。注意：大家都是注册在一个地方，且都是 '+ 这个方法。
4. 做运算时，定义统一的ADD,SUB。在ADD内去 4.1这个地方找函数，4.2当然也可以更加去中心化，每个带着对象带着自己的函数。  
5. 然后基于这个东西，靠组合就有了很强大的处理能力：多项式，分数多项式，甚至于矩阵，因为总是能递归到可操作的地方。


结论： 
- golang的interface就是这种定义的一种实现。然后基于组合去构建能力更强大的系统。  
- 在实际的生产工作中，可能下层的设计有各种方式，在不可控的情况下，那么我们可以加一层抽象，将表示与使用分开。这样上层就比较舒服了。  

几个问题：  
- 层次之间的界限问题即复杂度？
底层修改，上层怎么知道：必须得有一定的约定。
- 普通用户不需要知道有这么多类型，但是需要自己从零开始构建。  
- 这个系统非常完美，这个东西是怎么来的？ 长期演化+精心设计。
例如：3+7/2 怎么办？ 谁去做 3转化为 3/1, 在哪里做，多少人需要知道这个事情？ 这是系统后面的复杂之处。
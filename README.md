# 树形结构模型
树形结构在实际编程中应用广泛，如 省市区、组织架构、代理关系、目录结构、分类、分组等。
基于关系型数据结构存储树状模型的设计主要有邻接列表模式、物化路径、左右值编码、闭包表等几种方案，本模型选用最通用的邻接列表模式、物化路径相结合。
设计方案通用，抽象后节点接口为TreeNodeI,仓储接口为TreeRepositoryI
## 功能列表
1. 包内结构体treeNode 主要提供路径(path)和深度(depth)计算,以及和这2个值强相关的常见操作:增加节点、删除节点、移动节点、获取父节点、获取子节点
2. 包内结构体treeNodeIs 提供节点集合操作,主要包括:实现了接口TreeNodeI集合转treeNodeIs结构体、treeNodeIs结构体转目标结构体、重新计算所有的路径(path)和深度(depth)、二维数组转树结构、计算子节点数量等
## 辅助工具列表
1. EmptyTreeNode、EmptyTreeRpository 是TreeNodeI,TreeRepositoryI 的空实现,建议其它实现继承它,主要减少其它实现非必要接口,以及方便后续TreeNodeI,TreeRepositoryI扩展方法
2. SimpleTree 是一个简单的TreeNodeI 实现,主要用于案例、内存数据方式操作以及自我测试使用
xgbx
## 扩展
分层聚合其实也类似于树状结构，设计中将融合分层聚合的模型，以电商的商品分类为例，某个商品按商品管理分类，属于手机，按活动分类可以属于热门商品，基于这种多维度分类，模型提供树交叉功能
## 参考资料
https://www.cnblogs.com/goloving/p/13570067.html

按照DDD设计原理,实体（Entity）是在相同限界上下文中具有唯一标识的领域模型，可变，通过标识判断同一性
限界上下文: 映射关系像一颗树,有根节点,节点可以有任意多个子节点,叶子节点无子节点,实现节点的增、删、改(子节点移动)、查等功能
实体的唯一标识:nodeId是树节点模型实体的唯一标识,任何两个nodeId不同的node 实例,都表示不同的实体.
可变: 树节点模型的所有属性(nodeId、parentId、depth、path)在发生某些行为时都可能发生变化,label属性虽然在行为过程中不会变,但是会影响行为(目前设定lable=leaf时,不能增加子节点,关于这个业务规则是否有更好的变通实现方式,可以探讨)
 新增行为: 新增节点,会修改入参的depth、path属性
 移动节点行为: 移动节点会修改当前节点的 parentId、depth、path属性以及所有子节点的depth、path属性
 查询行为: 查询节点,不会修改属性,但是在新增、移动节点时,需要查询行为,查询行为可能觉得新增、移动行为对属性修改的结果.另外查询行为有自己内部逻辑(查询子树),所以也封装到entity中
 删除行为: 删除节点,并不会改变nodeId、parentId、depth、path中的任何属性,删除行为导致的结果比较简洁(删除字段标记即可),将删除属性移交给repository实现更具扩展性(数据表可以定义任何字段为删除字段)
 实体 access repository : repository 强IO操作,具有强副作用(相同输入,根据外部设备的不同——如网路、数据库不可用,会得到不同的结果)entity模型通过构造函数传入repository interface 实现依赖倒置效果,并且值依赖repository 的查询接口，entity 对写入数据不直接依赖repository，而是返回修改后的数据，从而对外部实现零约束，提升扩展性
 entity的创建: 统一采用 NewNodeEntity 工厂创建,nodeEntity对包外不可见,包外限定只能通过NewNodeEntity方法创建
 entity 的聚合: 包内不提供聚合，因为值对象不固定，聚合对象属性不固定，所以entity 聚合交由外部自主实现，这得益于entity对外部数据结构的零约束特性
 实体发布领域事件: 暂未实现,如果后续需要，则可以在包内实现生产领域事件，交由外部发布
 实体关联值对象: 不计划实现，实体可以和任意结构值对象组合生成聚合跟满足业务，因此无需在包内实现值对象关联


## 使用
1. 外部实现 TreeNodeI 接口,需要存储媒介实现TreeRepositoryI
2. 调用包方法ConvertToTreeNodes、NewTreeNode 转换成当前包内对象
3. 调用相应方法完成功能或者数据整理
4. 使用treeNodeIs相关功能后,调用treeNodeIs.Convert转换为目标结构体
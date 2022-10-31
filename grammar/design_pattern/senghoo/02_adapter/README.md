# 适配器模式

适配器模式用于转换一种接口适配另一种接口。

实际使用中 Adapter 一般为接口，并且使用工厂函数生成实例。

在 Adapter 中匿名组合 Adapter 接口，所以Adapter类也拥有 SpecificRequest 实例方法，又因为Go语言中非入侵式接口特征，其实 Adapter 也适配 Adapter 接口。

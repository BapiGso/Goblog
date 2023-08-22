# 使用sync.pool复用对象

## handler仅get对象和根据req参数调用方法，方便解耦

## struct分为不同属性模块进行分类，比如comment、content等

## 不同方法来更新属性，可以返回err到handler

## new函数在池里新建对象，get、put用于取出和放回

## 多次查询也不要关联查询，对我来说方便维护，性能不是很重要

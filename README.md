# gosearch
基于ElasticSearch封装的通用型搜索引擎，方便从外部批量导入数据，指定搜索字段，准精确模式与模糊模式进行中文、简拼混合搜索，未采用IK分词模式
本搜索引擎微服务模块基于自主的mgin微服务框架，使用到的模块
- maczh/mgconfig
- maczh/logs
- gin
- maczh/gintool
- maczh/mgtrace
- olivere/elastic
  - ElasticSearch 7.1.12
- mgo
  - MongoDB 3.4
- go-redis/redis
  - redis 4.x
- Nacos 1.2
- swagger
    
## API接口
本模块提供以下接口
 接口uri | 接口功能 
 ---|---:
 /gosearch/add | 插入文档到全文索引库
 /gosearch/add/batch | 批量插入文档到全文索引库
 /gosearch/del | 从全文索引库中删除文档
 /gosearch/del/batch | 从全文索引库中批量删除文档  
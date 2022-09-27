# 脚手架项目(CLD)

更新: 在使用struct读取配置时, 一定要注意结构相对应,`mapstructure`字段进行映射

-------------------

1. CLD简介:

   参考连接: [Go语言的CLD分层（MVC架构） - 简书 (jianshu.com)](https://www.jianshu.com/p/403f3316a5fb?u_atoken=c586f3f0-6db4-4e01-a522-249505496174&u_asession=01w3v6K38bujHaz_2GKrr2EyYrSW0yf7vpTBwG9HDZzwBWCV6lsb8n_XrC6o0KDK_xX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K8pSRB2CYcxuwcfoIbQ41w2hVaMhQoQvxDmSarRnKyUVWBkFo3NEHBv0PZUm6pbxQU&u_asig=05Il5fzBncc_3GGVCLVu67Y16sEb55MxYEY2oU4ZIPlTJQL57Ac86VnHtni2Jr3Q1PxpnNUOUAZ6r3whbSbWC8BhNVCRMaOUVkpSxYovclLnquT32CiXHsL6VJG571KiolKVE7sm_hYrUiFAKOL9dsSZhlmKpCAlvy_A_XR2m7JCT9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzb3mAd6LyQsZt_kcuSEksvlDUHVI0TAYcJPDwmkCx3d-CNkNfzbmXGypMDlqpKRYAO3h9VXwMyh6PgyDIVSG1W9a9y2wTUhS6dth2rzTG-BxHANw4_8BpU31TkZPiPBinlMMxKtIHlgiQ7whl39aSDMYMsCN7izKX97tlv6knaYHmWspDxyAEEo4kbsryBKb9Q&u_aref=edyiz5J6CXZFUh6u0Ii%2B2gzs1is%3D)

​	`Controller`，控制层，与上述类似，服务入口，负责处理路由，参数校验，请求转发。

​	`Logic/Service`，业务逻辑（服务）层，一般是业务逻辑的入口，可以认为从这里开始，所有的请求参数一定是合法的。业务逻辑和业务流程也都在这一层中。常见的设计中会将该层称为 Business Rules。

​	`DAO/Repository`，DAO层，这一层主要负责和数据、存储打交道。将下层存储以更简单的函数、接口形式暴露给 Logic 层来使用。负责数据的持久化工作。

2. controller-控制层

3. logic-业务逻辑层

4. logger-日志相关

5. dao-数据访问层

6. models-数据模型

7. pkg-第三方包

8. routers-路由层

9. setting-配置层

10. confing.yaml-配置文件

11. web_app.log-日志文件

12. 参考链接:

    - [Go Web开发进阶实战（gin框架） - 网易云课堂 (163.com)](https://study.163.com/course/courseLearn.htm?courseId=1210171207#/learn/video?lessonId=1281046428&courseId=1210171207)

      ...待补充


#README

## 1. 项目启动

1. 生成swagger 文档

   ```she
   swag init
   ```

	2. 导入组件

    ```she
    go mod tidy
    ```

	3. 打包

    ```she
    go build
    ```

	4. 启动

    ```she
    go run main.go
    ```



注意事项:

- 因当前开发环境为windows,导致以下问题:

  - 项目中包含了 Makefile 编译命令,导致编译无法正常运转, 还没找到合适方法, 相关命令可在`Makefile `中查看

    ```shell
    make build
    ```

  - 项目中包含了 air 热部署命令,air 的兼容性存在问题, 还没找到合适方法,相关命令可在 `.air.conf`文件中查看

    ```she
    air //可执行该命令, 启动项目
    ```

  - swagger 启动命令添加至air中未生效, 后续在进行测试验证

- 前端项目: `bells_of_ireland_frontend` 的相关命令

  - Project setup

    ```
    npm install
    ```

  - Compiles and hot-reloads for development

    ```
    npm run serve
    ```

  - Compiles and minifies for production

    ```
    npm run build
    ```

  - Lints and fixes files

    ```
    npm run lint
    ```

- 项目开发时,多参考官方文档, 用户随时会有变更
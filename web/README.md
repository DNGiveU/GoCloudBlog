# Ant Design Pro

This project is initialized with [Ant Design Pro](https://pro.ant.design). Follow is the quick guide for how to use.

## Environment Prepare

Install `node_modules`:

```bash
npm install
```

or

```bash
yarn
```

## Provided Scripts

Ant Design Pro provides some useful script to help you quick start and build with web project, code style check and test.

Scripts provided in `package.json`. It's safe to modify or add additional script:

### Start project

```bash
npm start
```

### Build project

```bash
npm run build
```

### Check code style

```bash
npm run lint
```

You can also use script to auto fix some lint error:

```bash
npm run lint:fix
```

### Test code

```bash
npm test
```

## More

You can view full document on our [official website](https://pro.ant.design). And welcome any feedback in our [github](https://github.com/ant-design/ant-design-pro).

## Landing Ant Design

[https://landing.ant.design/index-cn](https://landing.ant.design/index-cn)

在`ant design pro`的基础上集成.

### 文件路径

将`Landing Ant Design`打包的`Home`压缩包解压, 拷贝至`src/pages`下以`home`重命名

### 路由添加

```json
{ "path":  "/", "component": "./home"}
```

### 依赖安装

```shell script
$ npm install antd enquire-js rc-queue-anim rc-scroll-anim rc-tween-one --save;
$ npm install rc-banner-anim --save;// 如果用的是多屏滑动型的 banner，加上这条。
# 或者
$ yarn add antd enquire-js rc-queue-anim rc-scroll-anim rc-tween-one --save;
$ yarn add rc-banner-anim --save;
```

### Css Module

这里推荐使用`css-module`的`global`

在`home/less/antMotionStyle.less`中使用`global`:

```less
:global {
  @import './common.less';
  @import './custom.less';
  @import './content.less';
  @import './nav0.less';
  @import './banner0.less';
  ...
}
```

### 暂时先删除换肤插件

由于换肤插件需要重新 build 全部的 less, 暂时不支持 `Landing Ant Design` 的 less，所以我们先暂时删除换肤插件。

文件目录：`config/config.js`;

删除`webpackPlugin`相关的代码：

```
export default {
  // add for transfer to umi
  ...
  manifest: {
    basePath: '/',
  },

  // 注释掉
- chainWebpack: webpackPlugin,
};
```

## 三方库

[ant 社区精选组件](https://ant.design/docs/react/recommendation-cn)

[简图 mermaidjs](https://github.com/DNGiveU/mermaidjs.github.io)

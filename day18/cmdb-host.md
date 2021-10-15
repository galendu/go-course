# 主机列表页面

我们先完成主机的列表页面, 列表页面大致布局如下:

![](./images/list-page-layout.jpg)

我们列表页面布局大概如下:
+ 提示信息
+ 表格操作: 左边: 搜索区域， 右边: 操作区域
+ 表格数据: 一个Box存放, 底部是分页信息


当前的页面如下:
```html
<template>
  <div class="host-container">
    Host 页面
  </div>
</template>

<script>
export default {
  name: 'CmdbHost',
  data() {
    return {}
  }
}
</script>
```

## Tips组件

tips大概布局:
+ 一个box 容器, 带背景色, Box有一个关闭按钮
+ 第一行: icon + title文字描述
+ 下面是列表文字说明

先写组件模版:
```html
<template>
  <div v-if="!hidden" class="tips">
    <div class="titile">
        <div class="tips-icon">
            <svg-icon icon-class="tips-info"></svg-icon>
        </div>
        <span class="title-content">
            温馨提示
        </span>
        <span class="close-btn">
          <i class="el-icon-close" @click="handleClose" />
        </span>
        
    </div>
    <div class="content">
        <li class="tip-item"> 具体提示</li>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tips',
  data() {
    return {
      hidden: false,
    }
  },
  methods: {
    handleClose() {
      this.hidden = true
    }
  }
}
</script>
```

然后我们在Host列表页面 应用Tips组件:
```html
<template>
  <div class="host-container">
    <tips />
    Host 页面
  </div>
</template>

<script>
import Tips from '@/components/Tips'

export default {
  name: 'CmdbHost',
  components: { Tips },
  data() {
    return {}
  }
}
</script>
```


然后开始补充样式:
```scss
<style lang="scss" scoped>
.tips {
  width: 100%;
  background-color: rgba(48, 210, 190, 0.42);
  color: rgb(20, 105, 105);
  font-size: 12px;
  padding: 8px 16px;
}

.titile {
  display: flex;

  .tips-icon {
    font-size: 16px;
  }

  .title-content {
    margin-left: 12px;
    font-size: 13px;
    font-weight: 600;
  }

  .close-btn {
    margin-left: auto;
    cursor: pointer;
  }
}

.content {
  font-size: 12px;
  padding: 6px 26px 0px 26px;
}
</style>
```

然后我们来定义这个组件需要传入的变量: 
```js
<script>
export default {
  name: 'Tips',
  props: {
    title: {
      type: String,
      default: '温馨提示',
    },
    tips: {
      type: Array,
      default() {
        return []
      }
    }
  },
  // ...
}
</script>
```

最后给组件传递参数 进行测试:
```html
<tips :tips="['主机列表页面']" />
```

为了代码干净, 我们把tips定义在 data里面:
```html
<template>
  <div class="host-container">
    <tips :tips="tips" />
    Host 页面
  </div>
</template>

<script>
import Tips from '@/components/Tips'

const tips = [
  '现在仅同步了阿里云主机资产'
]

export default {
  name: 'CmdbHost',
  components: { Tips },
  data() {
    return {
      tips: tips
    }
  }
}
</script>
```

![](./images/tips.jpg)

## 表格数据

表格数据，我们采用Element UI 的Table组件: [Table 表格](https://element.eleme.cn/#/zh-CN/component/table)

表格我们采用卡片风格, [Border 边框](https://element.eleme.cn/#/zh-CN/component/border)里提供了box-shadow的样式:
```
基础投影 box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
浅色投影 box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
```

我们采用浅色投影, 因此补充一个全局样式: box-shadow
```css
.box-shadow {
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
}
```

然后我们copy一个样例过来看看效果:
```html
<template>
  <div class="host-container">
    <tips :tips="tips" />
    <div class="box-shadow">
      <el-table
        :data="tableData"
        style="width: 100%">
        <el-table-column
          prop="date"
          label="日期"
          width="180">
        </el-table-column>
        <el-table-column
          prop="name"
          label="姓名"
          width="180">
        </el-table-column>
        <el-table-column
          prop="address"
          label="地址">
        </el-table-column>
      </el-table>
    </div>
    Host 页面
  </div>
</template>

<script>
import Tips from '@/components/Tips'

const tips = [
  '现在仅同步了阿里云主机资产'
]

export default {
  name: 'CmdbHost',
  components: { Tips },
  data() {
    return {
      tips: tips,
      tableData: [{
          date: '2016-05-02',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄'
        }, {
          date: '2016-05-04',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1517 弄'
        }, {
          date: '2016-05-01',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1519 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1516 弄'
        }]
    }
  }
}
</script>
```

我们之前是把获取数据的代码直接写在该模块内部: 
```js
import axios from 'axios';

// ...

getHosts() {
  axios
    .get('http://localhost:8050/hosts', {params: this.query})
    .then(response => {
      this.tableData = response.data.data.items
      this.total = response.data.data.total
      console.log(this.tableData)
    })
    .catch(function (error) { // 请求失败处理
      console.log(error);
    });
},
```

这样并不方便互用和维护, 因此我们把这个路径都放到API模块下: api/cmdb/host.js

如果每个API都直接使用 axios, 那么后面做一些 中间件处理异常的逻辑就没发做了, 因此我们需要构造一个全局的axios实例



## 搜索框





## 资源同步



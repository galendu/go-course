# 博客后台

![](./images/backend-blog-list.png)

## 博客列表页

### 列表页

列表页由3部分组成:
+ 页头
+ 表格操作区
+ 表格

博客列表主要使用表格组件: [表格 Table](https://arco.design/vue/component/table)

```vue
<template>
  <div>
    <!-- 页头 -->
    <a-breadcrumb>
      <a-breadcrumb-item>文章管理</a-breadcrumb-item>
      <a-breadcrumb-item>文章列表</a-breadcrumb-item>
    </a-breadcrumb>
    <!-- 表格的操作区 -->
    <div class="operate">
      <div>
        <a-button size="small" type="primary">
          <template #icon>
            <icon-plus />
          </template>
          新建
        </a-button>
      </div>
      <div class="search">
        <a-input-search
          :style="{ width: '320px' }"
          placeholder="输入文章标题进行搜索"
        />
      </div>
    </div>
    <!-- 表格内容区 -->
    <div class="table">
      <a-table :data="data">
        <template #columns>
          <a-table-column title="标题" data-index="name"></a-table-column>
          <a-table-column title="作者" data-index="salary"></a-table-column>
          <a-table-column title="概要" data-index="address"></a-table-column>
          <a-table-column title="状态" data-index="email"></a-table-column>
          <a-table-column title="操作" align="center">
            <template #cell="{ record }">
              <a-space>
                <a-button
                  size="small"
                  status="danger"
                  @click="$modal.info({ title: 'Name', content: record.name })"
                  >删除
                </a-button>
                <a-button
                  size="small"
                  type="primary"
                  @click="$modal.info({ title: 'Name', content: record.name })"
                  >编辑
                </a-button>
                <a-button
                  size="small"
                  type="primary"
                  @click="$modal.info({ title: 'Name', content: record.name })"
                  >发布
                </a-button>
                <a-button
                  size="small"
                  type="primary"
                  @click="$modal.info({ title: 'Name', content: record.name })"
                  >查看
                </a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup>
import { reactive } from "vue";

const data = reactive([
  {
    key: "1",
    name: "Jane Doe",
    salary: 23000,
    address: "32 Park Road, London",
    email: "jane.doe@example.com",
  },
  {
    key: "2",
    name: "Alisa Ross",
    salary: 25000,
    address: "35 Park Road, London",
    email: "alisa.ross@example.com",
  },
  {
    key: "3",
    name: "Kevin Sandra",
    salary: 22000,
    address: "31 Park Road, London",
    email: "kevin.sandra@example.com",
  },
  {
    key: "4",
    name: "Ed Hellen",
    salary: 17000,
    address: "42 Park Road, London",
    email: "ed.hellen@example.com",
  },
  {
    key: "5",
    name: "William Smith",
    salary: 27000,
    address: "62 Park Road, London",
    email: "william.smith@example.com",
  },
]);
</script>

<style scoped>
.operate {
  display: flex;
  margin: 8px 0 8px 0;
}

.search {
  margin-left: auto;
}

.table {
  width: 100%;
}
</style>
```


### 跳转到详情页



### 编辑页



### 发布与删除


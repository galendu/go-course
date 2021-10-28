# CMDB前端页面 

## 调整script

之前我们使用 npm run serve, 是因为package.json配置的脚手架如下:
```json
"scripts": {
"serve": "vue-cli-service serve",
"build": "vue-cli-service build",
"lint": "vue-cli-service lint",
"svgo": "svgo -f src/icons/svg --config=src/icons/svgo.yml"
}
```

一般我们习惯调整成为: npm run dev
```json
"scripts": {
"dev": "vue-cli-service serve",
"build": "vue-cli-service build",
"lint": "vue-cli-service lint",
"svgo": "svgo -f src/icons/svg --config=src/icons/svgo.yml"
}
```

## 主机列表

根据我们的数据结构:
```json
{
    "code":0,
    "data":{
        "items":[
            {
                "id":"c5pu17p3n7phb445t73g",
                "sync_at":1634984095854,
                "secret_id":"",
                "vendor":"VENDOR_TENCENT",
                "resource_type":"HOST_RESOURCE",
                "region":"ap-shanghai",
                "zone":"ap-shanghai-3",
                "create_at":1593580796000,
                "instance_id":"ins-7sgo2va9",
                "resource_hash":"231eb53f386be0ad4a502ddfff663e107b0e0405",
                "describe_hash":"e168c71a705013413e83705dfe3ff3c4504b6a98",
                "expire_at":1657171197000,
                "category":"",
                "type":"S2.SMALL2",
                "name":"nbtuan-web",
                "description":"",
                "status":"RUNNING",
                "tags":null,
                "update_at":0,
                "sync_accout":"",
                "public_ip":[
                    "49.234.114.127"
                ],
                "private_ip":[
                    "172.17.0.7"
                ],
                "pay_type":"PREPAID",
                "resource_id":"c5pu17p3n7phb445t73g",
                "cpu":1,
                "memory":2,
                "gpu_amount":0,
                "gpu_spec":"",
                "os_type":"",
                "os_name":"CentOS 7.6 64bit",
                "serial_number":"f191197c-c009-4a08-9a52-e6f08bacebbf",
                "image_id":"img-9qabwvbn",
                "internet_max_bandwidth_out":1,
                "internet_max_bandwidth_in":0,
                "security_groups":[
                    "sg-05url5pe"
                ]
            }
        ],
        "total":1
    }
}
```

在列表页做对应的展示:

```html
<el-table :data="hosts" style="width: 100%">
    <el-table-column prop="name" label="名称">
        <template slot-scope="{ row }">
        {{ row.resource_id }} <br />
        {{ row.name }}
        </template>
    </el-table-column>
    <el-table-column prop="name" label="资产来源">
        <template slot-scope="{ row }">
        {{ row.vendor }} <br />
        {{ row.region }}
        </template>
    </el-table-column>
    <el-table-column prop="name" label="内网IP/外网IP">
        <template slot-scope="{ row }">
        {{ row.private_ip }} <br />
        {{ row.public_ip }}
        </template>
    </el-table-column>
    <el-table-column prop="name" label="系统类型">
        <template slot-scope="{ row }">
        {{ row.os_name }}
        </template>
    </el-table-column>
    <el-table-column prop="sync_at" label="创建时间">
        <template slot-scope="scope">
        {{ scope.row.create_at | parseTime }}
        </template>
    </el-table-column>
    <el-table-column prop="expire_at" label="过期时间">
        <template slot-scope="scope">
        {{ scope.row.expire_at | parseTime }}
        </template>
    </el-table-column>
    <el-table-column prop="name" label="规格">
        <template slot-scope="{ row }">
        {{ row.cpu }} / {{ row.memory }}
        </template>
    </el-table-column>
    <el-table-column prop="name" label="状态">
        <template slot-scope="{ row }">
        {{ row.status }}
        </template>
    </el-table-column>
    <el-table-column prop="操作" align="center" label="状态">
        <template slot-scope="{ row }">
        <el-button type="text" disabled>归档</el-button>
        <el-button type="text" disabled>监控</el-button>
        </template>
    </el-table-column>
</el-table>
```

## 主机搜索框

我们使用一个关键字输入框进行搜索支持:
+ instance_id
+ name
+ public_ip
+ pravite_ip

调整后端API, 关键字支持这些字段
```go
if req.Keywords != "" {
    query.Where("r.name LIKE ? OR r.id = ? OR r.instance_id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
        "%"+req.Keywords+"%",
        req.Keywords,
        req.Keywords,
        req.Keywords+"%",
        req.Keywords+"%",
    )
}
```

使用Postman测试关键字搜索是否正常

添加前端搜索框，对接关键字搜索
```html
<div class="search">
    <el-input
        v-model="query.keywords"
        placeholder="请输入实例ID|名称|IP 敲回车进行搜索"
        @keyup.enter.native="get_hosts"
    ></el-input>
</div>
```

测试关键字搜索是否都能正常


## CMDB详情页面





## CMDB搜索页面


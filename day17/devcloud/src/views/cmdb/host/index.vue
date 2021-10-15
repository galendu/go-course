<template>
  <div class="host-container">
    <tips :tips="tips" />
    <div class="box-shadow">
      <el-table
        :data="hosts"
        style="width: 100%">
        <el-table-column
          prop="name"
          label="名称"
          width="180">
        </el-table-column>
        <el-table-column
          prop="sync_at"
          label="同步时间"
          width="180">
          <template slot-scope="scope">
            {{ scope.row.sync_at | parseTime}}
          </template>
          
        </el-table-column>
        <el-table-column
          prop="description"
          label="描述">
        </el-table-column>
      </el-table>

      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="query.page_number"
        :page-sizes="[2,10, 20, 30, 50]"
        :page-size="query.page_size"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
      </el-pagination>
    </div>
    Host 页面
  </div>
</template>

<script>
import Tips from '@/components/Tips'
import { LIST_HOST } from '@/api/cmdb/host.js'

const tips = [
  '现在仅同步了阿里云主机资产'
]

export default {
  name: 'CmdbHost',
  components: { Tips },
  data() {
    return {
      tips: tips,
      query: {page_size: 20, page_number: 1},
      total: 0,
      hosts: []
    }
  },
  created() {
    this.get_hosts()
  },
  methods: {
    async get_hosts() {
      const resp = await LIST_HOST(this.query)
      console.log(resp)
      this.hosts = resp.data.items
      this.total = resp.data.total
    },
    handleSizeChange(val) {
      this.query.page_size = val
      this.get_hosts()
    },
    handleCurrentChange(val) {
      this.query.page_number = val
      this.get_hosts()
    }
  }
}
</script>

<style lang="scss" scoped>
.box-shadow {
  margin: 12px 0;
}
</style>
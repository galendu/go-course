  <template>
    <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column
        prop="id"
        label="资产ID"
        width="180">
      </el-table-column>
      <el-table-column
        prop="name"
        label="主机名称"
        width="180">
      </el-table-column>
      <el-table-column
        prop="region"
        label="地址">
      </el-table-column>
    </el-table>
  </template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      tableData: []
    }
  },
  mounted() {
    axios
      .get('http://localhost:8050/hosts')
      .then(response => {
        this.tableData = response.data.data.items
        console.log(this.tableData)
      })
      .catch(function (error) { // 请求失败处理
        console.log(error);
      });
  }
}
</script>
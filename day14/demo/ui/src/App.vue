  <template>
    <div>
      <el-table
        :data="tableData"
        style="width: 100%">
        <el-table-column
          prop="id"
          label="ID"
          width="180">
        </el-table-column>
        <el-table-column
          prop="name"
          label="姓名"
          width="180">
        </el-table-column>
        <el-table-column
          prop="region"
          label="Region">
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
      
      <h1 class="f12">基础标签</h1>
      <h1>这是一个标题</h1>
      <h2>这是一个标题</h2>
      <h3>这是一个标题</h3>
      <h4>这是一个标题</h4>

      <p>这是一个段落</p>
      <p>这是一个段落</p>

      <hr>
      <p>这是一个 <br />  段落</p>


      <h1>文本标签</h1>
        <!-- <del> 和 <ins> 一起使用，描述文档中的更新和修正。浏览器通常会在已删除文本上添加一条删除线，在新插入文本下添加一条下划线 -->
      <p></p>
      <del>test</del> <br>
      <i>定义斜体文本。</i> <br>
      <ins>定义被插入文本</ins> <br>
      <strong>加粗文本</strong>

      <p>这个文本包含 <sub>下标</sub>文本。</p>
      <p>这个文本包含 <sup>上标</sup> 文本。</p>

      <p>This is a <u>parragraph</u>.</p>

      <h1>表单标签</h1>
      <form action="demo_form.php">
        <label for="male">Male</label>
        <input type="radio" name="sex" id="male" value="male"><br>
        <label for="female">Female</label>
        <input type="radio" name="sex" id="female" value="female"><br><br>
        <input type="submit" value="提交">
      </form>

      <h1>内联框架</h1>
      <iframe src="//www.runoob.com">
        <p>您的浏览器不支持  iframe 标签。</p>
      </iframe>

      <h1>列表</h1>
      <ul id="list_menu" class="ul_class">
          <li id="coffee">Coffee</li>
          <li>Tea</li>
          <li>Milk</li>
          <div>
            <li>In Div</li>
          </div>
      </ul>

      <h1>表格</h1>
      <table border="1">
      <tr>
      <th>Month</th>
      <th>Savings</th>
      </tr>
      <tr>
      <td>January</td>
      <td>$100</td>
      </tr>
      </table>


      <p>
      <a href="#C4">查看章节 4</a>
      </p>

      <h2>章节 1</h2>
      <p>这边显示该章节的内容……</p>

      <h2>章节 2</h2>
      <p>这边显示该章节的内容……</p>

      <h2>章节 3</h2>
      <p>这边显示该章节的内容……</p>

      <h2><a>章节 4</a></h2>
      <p id="C4">这边显示该章节的内容……</p>
    </div>
  </template>

<script>
  import axios from 'axios';

  export default {
    data() {
      return {
        query: {
          page_size: 20,
          page_number: 1,
        },
        tableData: []
      }
    },
    mounted() {
      this.getHosts()
    },
    methods: {
      getHosts() {
          axios
            .get('http://localhost:8050/hosts', {params: this.query})
            .then(response => {
              console.log(response)
              this.tableData = response.data.data.items
              this.total = response.data.data.total
              console.log(this.tableData)
            })
            .catch(function (error) { // 请求失败处理
              console.log(error);
            });
        },
        handleSizeChange(val) {
          this.query.page_size = val
          this.getHosts()
        },
        handleCurrentChange(val) {
          this.query.page_number = val
          this.getHosts()
        }
    }
  }
</script>

<style scoped>

ul>li:first-child {
  font-weight: 600;
}

</style>
<template>
  <div class="app-container">
    <el-select
      v-model="namespaceList"
      filterable @change="refreshData"
      clearable
    >
      <el-option
        v-for="item in namespaceList"
        :key="item.index"
        :label="item.index"
        :value="item"
      >
      </el-option>
    </el-select>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="状态" width="95">
        <template slot-scope="scope">
          {{ scope.$index+1 }}
        </template>
      </el-table-column>

      <el-table-column label="名称" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="镜像">
        <template slot-scope="scope">
          {{ scope.row.image }}
        </template>
      </el-table-column>
      <el-table-column label="pod副本数" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.poddetail.currentrs }}/{{ scope.row.poddetail.disiredrs }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="Status" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="操作" width="300">
        <template slot-scope="scope">
          <i class="el-icon-time"/>
          <span>{{ scope.row.DisplayTime }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

  import {getDeployments} from '@/api/workloads'
  import {getNameSpaces} from '../../../api/workloads'

  export default {
    filters: {
      statusFilter(status) {
        const statusMap = {
          health: 'success',
          // true: 'gray',
          unhealth: 'danger'
        }
        return statusMap[status]
      }
    },
    data() {
      return {
        list: null,
        //这个是进入就请求服务器获取的, option 选择了也要刷新namespaceList
        namespaceList: "",
        listLoading: true,
        namespace: "default",
        select:'default'
      }
    },
    created() {
      this.fetchData()
    },
    watch: {
      // 如果路由有变化，会再次执行该方法
      '$route': 'fetchData'
    },
    methods: {
      fetchData() {
        this.listLoading = true
        // 先写死是这个,后面改成动态获取
        getDeployments(this.namespace).then(response => {
          this.list = response.data.lists
          this.listLoading = false
        })
        getNameSpaces().then(response => {
          this.namespaceList = response.data.lists
        })
      },
      refreshData(val) {
        this.$forceUpdate
        getNameSpaces().then(response => {
          this.namespaceList = response.data.lists
        })
        console.log("namespaceList:" + this.namespaceList)       
        this.listLoading = true
        getDeployments(val).then(response => {
          this.list = response.data.lists
          this.listLoading = false
        })
      }
    }
  }
</script>

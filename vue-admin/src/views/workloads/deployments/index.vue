<template>
  <div class="app-container">
    <el-select
      v-model="namespacelist"
      filterable
    >
      <el-option
        v-for="item in namespacelist"
        :key="item.name"
        :label="item.name"
        :value="item.name"
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
          <i class="el-icon-time" />
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
      //这个是进入就请求服务器获取的
      namespacelist: null,
      listLoading: true,
      namespace: "default"
    }
  },
  created() {
    this.fetchData()
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
        this.namespacelist = response.data.lists
      })
    }
  }
}
</script>

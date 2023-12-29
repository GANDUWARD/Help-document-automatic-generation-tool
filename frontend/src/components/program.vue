<template>
  <div class="all-box">
    <div class="left-tree">
      <p>左部区域用于展示文件结构</p>
      <ul>
        <li v-for="file in files" :key="file.path">
          {{ file.name }}
          <ul v-if="file.children && file.children.length">
            <file-tree :files="file.children"></file-tree>
          </ul>
        </li>
      </ul>
    </div>
    <div class="mid-area">
      <p>主区域用于文本查看编辑</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      files: [],
    };
  },
  async mounted() {
    // 异步请求获取文件结构数据
    try {
      const response = await this.fetchFileStructure('/api/getFileStructure');
      this.files = response.data;
    } catch (error) {
      console.error('Error fetching file structure:', error);
    }
  },
  methods: {
    async fetchFileStructure(apiEndpoint) {
      // 这里使用了 axios 库，你需要确保项目中已经安装了 axios
      return this.$axios.get(apiEndpoint);
    },
  },
};
</script>

<style scoped>
.all-box{
  display: flex;
  height: 100vh;
}
.left-tree {
  width: 20%;
  color: #53b4f5;
  border: #01080e;
  border-style: solid;
  height: 100%;
}
.mid-area {
  flex-grow: 1;
  color: bisque;
  border-style: dashed;
  height: 100;
}
</style>

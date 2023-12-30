<!-- program.vue -->
<template>
  <div class="all-box">
    <div class="left-tree">
      <p>文件结构</p>
      <!-- 监听自定义事件，更新右侧区域的数据 -->
      <file-tree-node :file="files" @file-clicked="handleFileClicked"></file-tree-node>
    </div>
    <div class="mid-area">
      <p>文本查看编辑</p>
      <div v-if="isEditing" class="editing-container">
        <!-- 新建状态，显示文本框和保存按钮 -->
        <input v-model="editedFileName" placeholder="文件名" />
        <textarea v-model="editedFileContent" placeholder="文件内容"></textarea>
        <button @click="saveChanges">保存</button>
      </div>
      <div v-else>
        <!-- 显示文件内容 -->
        <div class="code-content">{{ midAreaData }}</div>
        <!-- 显示编辑按钮 -->
        <button @click="startEditing">编辑</button>
        <button @click="exportFile">导出</button> 
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import FileTreeNode from './FileTreeNode.vue';

export default {
  name: 'HelloWorld',
  components: {
    'file-tree-node': FileTreeNode,
  },
  data() {
    return {
      files: {},
      midAreaData: null, // 添加用于存储右侧区域的数据
      isEditing: true, // 初始化为 true，即一开始就显示新建状态
      editedFileName: '', // 添加编辑时的文件名
      editedFileContent: '', // 添加编辑时的文件内容
      filename:''   //用于更新文件名称
    };
  },
  async mounted() {
    try {
      this.files = await this.fetchFileStructure('http://127.0.0.1:8888/filelist/fl');
      console.log(this.files);
    } catch (error) {
      console.error('Error fetching file structure:', error);
    }
  },
  methods: {
    async fetchFileStructure(apiEndpoint) {
      try {
        const response = await axios.get(apiEndpoint);
        const responseData = response.data.data;

        // 如果返回的数据是字符串，将其转换为对象
        if (typeof responseData === 'string') {
          return JSON.parse(responseData);
        }

        return responseData;
      } catch (error) {
        console.error('Error fetching file structure:', error);
        throw error; // 重新抛出错误，以便在调用方处理
      }
    },
    async handleFileClicked(filePath) {
      try {

        console.log('click on ', filePath);
        const response = await axios.post('http://127.0.0.1:8888/file', { path: filePath });
        console.log(response.data.data);

        // 判断文件类型是否为 HTML
        if (filePath.endsWith('.html')) {
          // 在新标签页中打开 HTML 内容
          const newTab = window.open();
          newTab.document.write(response.data.data);
        } else {
          // 在右侧区域显示其他文件内容
          this.midAreaData = response.data.data;
       // 更新 this.filename，仅保留文件名部分
       this.filename = filePath.split('/').pop().split('\\').pop();
        }

        // 退出编辑状态
        this.isEditing = false;
      } catch (error) {
        console.error('Error sending request:', error);
      }
    },
    startEditing() {
      // 进入编辑状态，保存当前文件名和文件内容
      this.isEditing = true;
      this.editedFileName = this.filename; // 使用当前文件的名字
      this.editedFileContent = this.midAreaData; // 使用当前文件的内容
    },
    saveChanges() {
      // 保存修改后的文件名和文件内容，以及文件路径
      const editedFilePath = 'C:\\Users\\ASUS\\Documents\\GitHub\\Help-document-automatic-generation-tool\\backend\\taskarea'; 
      const editedFileName = this.editedFileName;
      const editedFileContent = this.editedFileContent;

      // 发送请求保存修改
      axios.post('http://127.0.0.1:8888/save', {
        path: editedFilePath,
        name: editedFileName,
        content: editedFileContent,
      })
      .then(response => {
        console.log('Save successful:', response.data);
      })
      .catch(error => {
        console.error('Error saving changes:', error);
      });

      // 退出编辑状态
      this.isEditing = false;
    },
    exportFile() {
      // 导出文件，发送请求到后端
      const exportedFileName = this.filename;
      axios.post('http://127.0.0.1:8888/export', {
        path: 'C:\\Users\\ASUS\\Documents\\GitHub\\Help-document-automatic-generation-tool\\backend\\taskarea',
        name: exportedFileName,
      })
      .then(response => {
        console.log('Export successful:', response.data);
        // 可以根据需要处理导出后的响应
      })
      .catch(error => {
        console.error('Error exporting file:', error);
      });
    },
  },
};
</script>

<style scoped>
.all-box {
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
  color: rgb(0, 0, 0);
  background-color: antiquewhite;
  border-style: dashed;
  height: 100%;
  overflow: auto; /* 添加 overflow 属性以处理过长的文本 */
}


.editing-container {
  display: flex;
  flex-direction: column;
  height: 100%; /* 设置编辑容器高度为100% */
}

.editing-container input,
.editing-container textarea {
  margin-bottom: 10px; /* 设置文本框之间的垂直间距 */
  padding: 10px; /* 设置文本框内边距 */
  width: 100%; /* 设置文本框宽度为容器宽度 */
  box-sizing: border-box; /* 包含 padding 在内的盒模型 */
}

/* 修改文件内容文本框高度为70% */
.editing-container textarea {
  flex: 1; /* 使用 flex 属性填充编辑容器的剩余空间 */
  height: 70%;
}

.code-content {
  white-space: pre-wrap;
  height: 50%; /* 设置文件内容文本框高度 */
  overflow: auto; /* 添加 overflow 属性以处理过长的文本 */
}
</style>
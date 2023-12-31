<template>
  <div class="all-box">
    <div class="left-tree">
      <p>文件结构</p>
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
      midAreaData: null,
      isEditing: true,
      editedFileName: '',
      editedFileContent: '',
      filename: '',
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

        if (typeof responseData === 'string') {
          return JSON.parse(responseData);
        }

        return responseData;
      } catch (error) {
        console.error('Error fetching file structure:', error);
        throw error;
      }
    },
    async handleFileClicked(filePath) {
      try {
        console.log('click on ', filePath);
        const response = await axios.post('http://127.0.0.1:8888/file', { path: filePath });
        console.log(response.data.data);

        if (filePath.endsWith('.html')) {
          const newTab = window.open();
          newTab.document.write(response.data.data);
        } else {
          this.midAreaData = response.data.data;
          this.filename = filePath.split('/').pop().split('\\').pop();
        }

        this.isEditing = false;
      } catch (error) {
        console.error('Error sending request:', error);
      }
    },
    startEditing() {
      this.isEditing = true;
      this.editedFileName = this.filename;
      this.editedFileContent = this.midAreaData;
    },
    saveChanges() {
      const editedFilePath = 'C:\\Users\\ASUS\\Documents\\GitHub\\Help-document-automatic-generation-tool\\backend\\taskarea';
      const editedFileName = this.editedFileName;
      const editedFileContent = this.editedFileContent;

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

      this.isEditing = false;
    },
    exportFile() {
      const exportedFileName = this.filename;
      axios.post('http://127.0.0.1:8888/export', {
        path: 'C:\\Users\\ASUS\\Documents\\GitHub\\Help-document-automatic-generation-tool\\backend\\taskarea',
        name: exportedFileName,
      })
      .then(response => {
        console.log('Export successful:', response.data);
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
  overflow: auto;
}

.editing-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.editing-container input,
.editing-container textarea {
  margin-bottom: 10px;
  padding: 10px;
  width: 100%;
  box-sizing: border-box;
}

.editing-container textarea {
  flex: 1;
  height: 70%;
}

.code-content {
  white-space: pre-wrap;
  height: 50%;
  overflow: auto;
}
</style>
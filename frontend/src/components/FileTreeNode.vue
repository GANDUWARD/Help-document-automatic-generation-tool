<template>
  <li>
    <span @click="handleFileClick">{{ file.name }}</span>
    <ul v-if="file.firstchild">
      <file-tree-node :file="file.firstchild" @file-clicked="handleFileClicked"></file-tree-node>
    </ul>
    <file-tree-node v-if="isObject(file.nextsibling)" :file="file.nextsibling" @file-clicked="handleFileClicked"></file-tree-node>
  </li>
</template>

<script>
export default {
  name: 'FileTreeNode',
  props: {
    file: {
      type: Object,
      required: true,
    },
  },
  methods: {
    isObject(obj) {
      return obj !== null && typeof obj === 'object' && !Array.isArray(obj);
    },
    handleFileClick() {
      this.$emit('file-clicked', this.file.path);
    },
    handleFileClicked(filePath) {
      this.$emit('file-clicked', filePath);
    },
  },
};
</script>
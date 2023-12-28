<template>
  <div class="left-tree">
    <p>左部区域用于展示文件结构</p>
  </div>
  <div class="mid-area">
    <p>主区域用于文本查看编辑</p>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
  },
  data(){
    return{
      files:""
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.left-tree {
  width: 20%;
  color: #53b4f5;
  border: #01080e;
  border-style:solid;
}

.mid-area {
  flex-grow: 1;
  color: bisque;
  border-style: dashed;
}
</style>

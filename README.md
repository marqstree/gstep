# gstep
a workflow engine  
show you how to use gstep


<p align="center">
    <img width="200" height="200" src="https://www.bqdnao.com/faceroop-static/gstep_logo.png">
</p>

# 1. go backend project startup
## create database
create mysql8 database name: gstep  
run mysql database script: (gstep project path)/mysql/mysql8.sql

## configure
config file: (gstep project path)/config.json
```
{
  "port": "9900",
  "auth": {
    "secret": "abcd"
  },
  "db":{
    "database": "gstep",
    "host": "localhost",
    "port": "3306",
    "user": "root",
    "password": "password"
  },
  "notify": {
    "taskStateChange": "http://localhost:9900/notify/task_state_change"
  },
  "department": {
    "rootParentDepartmentId": "0"
  }
}
```

## run
go run main.go



# 2. vue3 project setup
## install vue3 components
```
npm install gstep-vue --save
npm install element-plus --save
```

## import vue3 components in main.js
```js
import { createApp } from 'vue'
import App from './App.vue'

// import element-plus css style
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
// import gstep-vue workflow eidtor component css style
import 'gstep-vue/style.css'
import Editor from 'gstep-vue'

// use vue3 components
createApp(App)
    .use(ElementPlus)
    .use(Editor)
    .mount('#app')
```

## use gstep-vue component in your vue3 page
```js
<template>
  <div class="page">
    <!-- gstep editor -->
    <Editor class="editor" 
      baseUrl="http://localhost:9900" 
      :template="template"
      @cancel="onCancel"
      @save="onSave" />
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
// workflow chart data
const template = ref({})

onMounted(() => {
  // init workflow data
  template.value = {
    id: 0,          //new workflow set 0
    groupId: null   //new workflow set null
  }
})

// toolbar cancel callback
const onCancel=()=>{
  console.log('cancel')
}

// toolbar save callback
const onSave=()=>{
  console.log('save')
}

</script>

<style scoped>
.page {
  width: 100%;
  height: 100vh;
}

.editor {
  width: 100%;
  height: 100%;
}
</style>
```

## run vue3 project
```
npm run dev
```

## view your page
<p align="center">
    <img width="200" height="200" src="https://www.bqdnao.com/faceroop-static/gstep_vue.jpeg">
</p>

# 3. handle workflow process api path
1.workflow template  
+ save  
/template/save  
+ query  
/template/query  
+ detail  
/template/detail  

2.workflow process  
+ start  
/process/start  
+ audit pass  
/task/pass  
+ audit go back  
/task/retreat  
+ process cease  
/task/cease  

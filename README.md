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
  // 1)new template
  template.value = {
  }

  // 2)exist template version
  // template.value = {
  //   id: 1,
  // }
  
  // 3)exist template new version
  // template.value = {
  //   templateId: 1
  // }
})

// toolbar cancel callback
const onCancel=()=>{
  console.log('cancel')
}

// toolbar save callback
const onSave=()=>{
  console.log('save')
  console.log('+++ saved template ++++++')
  console.log(template.value)
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
<img width="400"  src="https://www.bqdnao.com/faceroop-static/gstep_vue.jpg">

# 3. handle workflow process api path
1.workflow template  
+ save template basic info  
path:/template/save_info  
method:post  
Content-Type:application/json  
json body:  
```
  {
      "id": 64,
      "title": "请假"
  }
```


+ get template basic info  
/template/info  
  path:/template/info  
  method:post  
  Content-Type:application/json  
  json body:
```
{
    "templateId": 1,
    "versionId": 65
}
```

+ get template detail  
  /template/info  
  path:/template/detail  
  method:post  
  Content-Type:application/json  
  json body:
```
{
    "versionId": 81
}
```

+ query template by page
  /template/query  
  path:/template/query  
  method:post  
  Content-Type:application/json  
  json body:
```
{
    "limit":10,
    "page":1
}
```

2.workflow process  
+ start new process 
    path:/process/start  
    method:post  
    Content-Type:application/json  
    json body:
```
{
    "templateId": 2,
    "startUserId": "103",
    "form": {
        "day": 5
    }
}
```

+ approve a process task   
  path:/task/pass  
  method:post  
  Content-Type:application/json  
  json body:
```
{
    "taskId": 75,
    "form": {
        "day": 11
    },
    "userId": "001"
}
```


+ retreat process task  
  path:/task/retreat  
  method:post  
  Content-Type:application/json  
  json body:
```
{
    "taskId": 74,
    "userId": "301"
}
```

+ process cease  
  path:/task/cease  
  method:post  
  Content-Type:application/json  
  json body:
```
{
    "taskId": 74,
    "userId": "301"
}
```

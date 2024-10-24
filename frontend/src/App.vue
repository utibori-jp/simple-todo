<template>
  <div>
    <h1>Simple Todo List</h1>
    <ul>
      <li v-for="(todo, index) in todos" :key="todo.No" class="todo-item">
        <span v-if="!todo.isEditing">{{ index+1 }}: {{ todo.task }}</span>
        <input v-if="todo.isEditing" v-model="changedTask" :placeholder="todo.task" />

        <div class="button-group">
          <button v-if="!todo.isEditing" @click="startEditing(todo)">編集</button>
          <button v-if="!todo.isEditing" @click="deleteTodo(todo)">削除</button>
          <button v-if="todo.isEditing" @click="updateTodo(todo)">完了</button>
          <button v-if="todo.isEditing" @click="resetEditingMode()">キャンセル</button>
        </div>
      </li>
    </ul>
    <div class="add-todo">
      <input v-model="newTodo" placeholder="新しいタスクを入力" />
      <button @click="addTodo">追加</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      todos: [],
      newTodo: '',
      changedTask: ''
    };
  },
  methods: {
    initialTodo() {
      this.todos = [{No: 1, Task: "Hello"}];
    },
    fetchTodos() {
      fetch('http://localhost:8081/todos') // バックエンドAPIのエンドポイント
        .then(response => response.json())
        .then(data => {
          console.log(data)
          this.todos = data;
        })
      this.resetEditingMode()
    },
    addTodo() {
      const task = this.newTodo.trim();
      console.log(task)
      if (task) {
        fetch('http://localhost:8081/todos', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({Task: task})
        })
        .then(response => response.json())
        .then(data => {
          console.log(data)
          this.fetchTodos()
          this.newTodo = ''; // 入力フォームをクリア
        })
        .catch(error => {
          console.error('Error:', error);
        });
      }
    },
    updateTodo(todo) {
      fetch(`http://localhost:8081/todos/${todo.no}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({Task: this.changedTask}),
      })
      .then(response => response.json())
      .then(data => {
        this.fetchTodos()
      })
      .catch(error => {
        console.error('Error:', error)
        this.resetEditingMode()
      })
    },
    deleteTodo(todo) {
      fetch(`http://localhost:8081/todos/${todo.no}`,{
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      })
      .then(data => {
        this.fetchTodos()
      })
      .catch(error => {
        console.error('Error:', error)
        this.resetEditingMode()
      })
    },
    startEditing(todo) {
      todo.isEditing = true
    },
    resetEditingMode() {
      for (const todo of this.todos) {
        todo.isEditing = false
      }
      this.changedTask = ''
    },
  },
  created() {
    this.fetchTodos();
  },
};
</script>

<style scoped>
.todo-item {
  display: flex;
  justify-content: space-between; /* リスト項目の内容とボタンを両端に配置 */
  align-items: center; /* 垂直方向に中央揃え */
  padding: 10px; /* パディングを追加してスペースを作成 */
  border-bottom: 1px solid #ddd; /* 下にボーダーを追加 */
}

.button-group {
  display: flex; /* ボタンを横並びにする */
  gap: 10px; /* ボタンの間にスペースを追加 */
}

.add-todo {
  display: flex; /* 入力フィールドと追加ボタンを横並びにする */
  margin-top: 20px; /* 上にマージンを追加 */
}

input {
  padding: 10px; /* 入力フィールドのパディングを調整 */
  border: 1px solid #ccc; /* 入力フィールドのボーダー */
  border-radius: 5px; /* 入力フィールドの角を丸める */
  flex-grow: 1; /* 残りのスペースを占める */
  margin-right: 10px; /* ボタンとの間にマージンを追加 */
}


button {
  padding: 5px 10px; /* ボタンのパディングを調整 */
  background-color: #007bff; /* ボタンの背景色を設定 */
  color: white; /* ボタンの文字色を白に設定 */
  border: none; /* ボタンのボーダーを無くす */
  border-radius: 5px; /* ボタンの角を丸める */
  cursor: pointer; /* マウスカーソルをポインターに変更 */
}

button:hover {
  background-color: #0056b3; /* ホバー時の背景色 */
}
</style>

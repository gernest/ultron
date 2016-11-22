Vue.http.options.root = '/root';
Vue.http.headers.common['Content-Type'] = 'application/json';

new Vue({
    el: '#app',
    data: {
      task: { title: '', detail: '', date: '' },
      tasks: []
    },
    delimiters: ['${', '}'],
    mounted: function () {
      this.fetchTasks();
    },
    methods: {
      fetchTasks: function () {
        var tasks = [];
        this.$http.get('/tasks/', "", {emulateJSON: true})
        .then(response => response.json())
        .then(result => {
            Vue.set(this.$data, 'tasks', result);
            console.log("success in getting tasks");
        })
        .catch(error => {
            console.log(error);
        });
      },
      addTask: function () {
        if (this.task.title.trim()) {
          this.$http.post('/tasks/', this.task,{emulateJSON: true})
          .then(response => response)
          .then( result => {
              this.tasks.push(this.task);
              console.log('Task added!');
              this.task = { title: '', detail: '', date: '' };
          }).catch( error => {
              console.log(error);
          });
        }
      },
      deleteTask: function (index) {
        if (confirm('Really want to delete？')) {
          console.log(index);
          this.$http.delete('/tasks/' + index)
          .then(response => response)
          .then( result => {
              console.log(result);
              this.tasks.splice(index, 1);
          }).catch( error => {
              console.log(error);
              alert("unable to delete")
          });
        }
      }
    }
});

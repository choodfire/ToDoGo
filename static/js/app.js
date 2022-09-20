const App = {
	data() {
		return {
			inputValue: "",
			tasksArr: [],
			doneTasksArr: []
		}
	},
	methods: {
		addNewTask() {
			if (this.inputValue) {
				this.tasksArr.push(this.inputValue);
				this.inputValue = "";
			}
		},
		deleteTask(index) {
			this.tasksArr.splice(index, 1);
		},
		doneTask(index) {
			let temp = this.tasksArr.splice(index, 1);
			this.doneTasksArr.push(temp[0]);
		},
		deleteDoneTask(index) {
			this.doneTasksArr.splice(index, 1);
		},
		undoTask(index) {
			let temp = this.doneTasksArr.splice(index, 1);
			this.tasksArr.push(temp[0]);
		}
	},
}

Vue.createApp(App).mount('#app');
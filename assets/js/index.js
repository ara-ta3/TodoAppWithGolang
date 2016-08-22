(function() {
    "use strict";

    class TodoModel {
        constructor(id, title, description, done) {
            this.id = id;
            this.title = title;
            this.description = description;
            this.done = done;
        }
    }

    class TodoList extends React.Component {
        render() {
            const todos = this.props.todos.map((todo, idx) => {
                return (
                        <Todo key={idx + "-" + Date.now()} origin={todo} onTodoDone={this.props.onTodoDone}/>
                       )
            });
            return (
                    <div className="todolist">
                    <h2>残タスク</h2>
                    {todos}
                    </div>
                   );
        }
    }
    class Todo extends React.Component {
        componentDidMount() {
            const checked = ReactDOM.findDOMNode(this.refs.checkbox).checked;
            this.applyStyle(checked);
        }

        applyStyle(checked) {
            const textForm = ReactDOM.findDOMNode(this.refs.text);
            if(checked) {
                textForm.className = "text-muted";
                textForm.style.textDecoration = "line-through";
            } else {
                textForm.className = "text-primary";
                textForm.style.textDecoration = "";
            }
        }

        handleTodo() {
            const checked = ReactDOM.findDOMNode(this.refs.checkbox).checked;
            ReactDOM.findDOMNode(this.refs.checkbox).checked = !checked;
            this.props.onTodoDone(this.props.origin.id, !checked, () => {
                this.applyStyle(!checked);
            });
        }

        render() {
            return (
                    <div className="todo">
                    <label>
                    <input ref="checkbox" type="checkbox" className="todoCheckbox" checked={this.props.origin.done} />
                    </label>
                    <span ref="text" onClick={this.handleTodo.bind(this)}>  {this.props.origin.title}  ( {this.props.origin.description} ) </span>
                    </div>
                   );
        }
    }

    class TodoForm extends React.Component {
        handleSubmit(e) {
            e.preventDefault();

            const text = ReactDOM.findDOMNode(this.refs.text).value.trim();
            const description = ReactDOM.findDOMNode(this.refs.description).value.trim();
            text && this.props.onTodoSubmit(text, description, () => {
                ReactDOM.findDOMNode(this.refs.text).value = "";
                ReactDOM.findDOMNode(this.refs.description).value = "";
            });
        }
        render() {
            return (
                    <form className="todoForm" onSubmit={this.handleSubmit.bind(this)}>
                    <div className="form-group">
                    <input type="text" className="form-control" placeholder="Todo Title" ref="text" required/>
                    </div>
                    <div className="form-group">
                    <input type="text" className="form-control" placeholder="Todo Description" ref="description" required/>
                    </div>

                    <input type="submit" className="btn btn-primary" value="追加" />
                    </form>
                   );
        }
    }

    class TodoApp extends React.Component {
        constructor(props) {
            super(props);
            this.state = {
                todos: []
            }
        }

        componentDidMount() {
            this.loadTodos();
        }

        loadTodos() {
            const req = new XMLHttpRequest();
            req.addEventListener('load', (event) => {
                const res = JSON.parse(event.target.response);
                if (res.error != null) {
                    console.error(res.error);
                    return;
                }
                const todos = Object.keys(res.todos).map((key) => {
                    const t = res.todos[key];
                    return new TodoModel(t.id, t.title, t.description, t.done);
                });
                this.setState({
                    todos: todos
                });
            });
            req.addEventListener('error', function(event) {
                console.error(event);
            });
            req.open('GET', '/api/todo');
            req.send();
        }

        onTodoDone(id, done, callback) {
            const req = new XMLHttpRequest();
            req.addEventListener('load', (event) => {
                const res = JSON.parse(event.target.response);
                if (res.error != null) {
                    console.error(res.error);
                    return;
                }
                callback();
            });
            req.addEventListener('error', (event) => {
                console.error(event.error);
            });
            req.open('PUT', `/api/todo/${id}/`);
            req.setRequestHeader( 'Content-Type', 'application/x-www-form-urlencoded' );
            req.send(`done=${encodeURIComponent(done)}`);
        }

        onTodoSubmit(title, description, callback) {
            const req = new XMLHttpRequest();
            req.addEventListener('load', (event) => {
                const res = JSON.parse(event.target.response);
                if (res.error != null) {
                    console.error(res.error);
                    return;
                }
                callback();
                this.loadTodos();
            });
            req.addEventListener('error', (event) => {
                console.error(event.error);
            });
            req.open('POST', '/api/todo');
            req.setRequestHeader( 'Content-Type', 'application/x-www-form-urlencoded' );
            req.send([
                `title=${encodeURIComponent(title)}`,
                `description=${encodeURIComponent(description)}`
            ].join("&"));
        }

        render() {
            return (
                    <div>
                        <div className="row">
                        <h1>俺が作った最高のTodoApp</h1>
                        </div>
                        <hr />
                        <div className="row">
                            <TodoList todos={this.state.todos} onTodoDone={this.onTodoDone.bind(this)}/>
                            <hr />
                            <TodoForm onTodoSubmit={this.onTodoSubmit.bind(this)}/>
                        </div>
                    </div>
                   );
        }
    }

    ReactDOM.render(
            <TodoApp />,
            document.getElementById("contents")
            );
})()


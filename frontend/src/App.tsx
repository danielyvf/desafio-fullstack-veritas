// src/App.tsx
import { useEffect, useState } from 'react';
import type { Task, TaskStatus } from './types';
import { getTasks, createTask, updateTask, deleteTask } from './services/api';
import { Column } from './components/Column';
import { TaskModal } from './components/TaskModal';
import './App.css';

export function App() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    loadTasks();
  }, []);

  async function loadTasks() {
    try {
      setIsLoading(true);
      setError('');
      const data = await getTasks();
      setTasks(data);
    } catch (err: any) {
      setError(err.message || 'Falha ao carregar as tarefas da API');
    } finally {
      setIsLoading(false);
    }
  }

  async function handleCreateTask(title: string, description: string) {
    const newTask = await createTask(title, description);
    setTasks((prev) => [...prev, newTask]);
  }

  async function handleMoveTask(task: Task, newStatus: TaskStatus) {
    const updatedTask: Task = { ...task, status: newStatus };
    try {
      setTasks((prev) =>
        prev.map((t) => (t.id === task.id ? updatedTask : t))
      );
      await updateTask(task.id, updatedTask);
    } catch (err) {
      await loadTasks();
    }
  }

  async function handleDeleteTask(id: string) {
    try {
      setTasks((prev) => prev.filter((t) => t.id !== id));
      await deleteTask(id);
    } catch (err) {
      await loadTasks();
    }
  }

  return (
    <div className="app-container">
      <header className="app-header">
        <h1 className='titulo1'>TASK.</h1>
        <button className="primeiro-btn" onClick={() => setIsModalOpen(true)}>
          + Nova Tarefa
        </button>
      </header>

      {error && (
        <div className="error-banner">
          {error} <button onClick={loadTasks}>Tentar Novamente</button>
        </div>
      )}

      {isLoading ? (
        <div className="loading-state">Carregando tarefas...</div>
      ) : (
        <main className="kanban-board">
          <Column
            title="A Fazer"
            status="TODO"
            tasks={tasks}
            onDeleteTask={handleDeleteTask}
            onMoveTask={handleMoveTask}
          />
          <Column
            title="Em Progresso"
            status="IN_PROGRESS"
            tasks={tasks}
            onDeleteTask={handleDeleteTask}
            onMoveTask={handleMoveTask}
          />
          <Column
            title="Concluído"
            status="DONE"
            tasks={tasks}
            onDeleteTask={handleDeleteTask}
            onMoveTask={handleMoveTask}
          />
        </main>
      )}

      <TaskModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onCreate={handleCreateTask}
      />
    </div>
  );
}

export default App;
import type { Task, TaskStatus } from '../types';
import '../App.css';

interface TaskCardProps {
  task: Task;
  onDelete: (id: string) => void;
  onMove: (task: Task, newStatus: TaskStatus) => void;
}

export function TaskCard({ task, onDelete, onMove }: TaskCardProps) {
  return (
    <div className="task-card">
      <div className="task-card-header">
        <h3>{task.title}</h3>
        <button 
          className="delete-btn" 
          onClick={() => onDelete(task.id)}
          title="Excluir tarefa"
        >
          ✕
        </button>
      </div>

      {task.description && <p>{task.description}</p>}

      <div className="task-card-actions">
        {task.status !== 'TODO' && (
          <button onClick={() => onMove(task, task.status === 'TODO' ? 'IN_PROGRESS' : 'DONE')}>
            ← Mover
          </button>
        )}
        {task.status !== 'DONE' && (
          <button onClick={() => onMove(task, task.status === 'TODO' ? 'IN_PROGRESS' : 'DONE')}>
            Mover →
          </button>
        )}
      </div>
    </div>
  );
}
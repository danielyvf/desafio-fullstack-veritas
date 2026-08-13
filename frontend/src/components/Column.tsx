import type { Task, TaskStatus } from '../types';
import { TaskCard } from './TaskCard';

interface ColumnProps {
  title: string;
  status: TaskStatus;
  tasks: Task[];
  onDeleteTask: (id: string) => void;
  onMoveTask: (task: Task, newStatus: TaskStatus) => void;
}

export function Column({ title, status, tasks, onDeleteTask, onMoveTask }: ColumnProps) {
  const filteredTasks = tasks.filter((task) => task.status === status);

  return (
    <div className="kanban-column">
      <div className="column-header">
        <h2>{title}</h2>
        <span className="task-count">{filteredTasks.length}</span>
      </div>

      <div className="column-body">
        {filteredTasks.length === 0 ? (
          <div className="empty-state">Nenhuma tarefa</div>
        ) : (
          filteredTasks.map((task) => (
            <TaskCard
              key={task.id}
              task={task}
              onDelete={onDeleteTask}
              onMove={onMoveTask}
            />
          ))
        )}
      </div>
    </div>
  );
}
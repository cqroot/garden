type Task = {
  id: number;
  title: string;
  note: string;
  project: number;
  due: number;
  status: number;
};

export type { Task };

export function copyTask(src: Task, dst: Task) {
  src.title = dst.title;
  src.note = dst.note;
  src.project = dst.project;
  src.due = dst.due;
  src.status = dst.status;
}

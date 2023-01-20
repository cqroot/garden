export function formattedDate(timestamp: number): string {
  const date = new Date(timestamp);
  return (
    date.getFullYear() + "-" + (date.getMonth() + 1) + "-" + date.getDate()
  );
}

export function todayTimestamp(): number {
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  return today.valueOf();
}

export function tomorrowTimestamp(): number {
  const tomorrow = new Date();
  tomorrow.setHours(0, 0, 0, 0);
  tomorrow.setDate(tomorrow.getDate() + 1);
  return tomorrow.valueOf();
}

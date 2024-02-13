export type Activity = {
  id: string;
  goalId: string;
};

export type CreateActivityData = {
  goalId: string;
  dur: string;
  durType: string;
};
export async function createActivity(
  data: CreateActivityData
): Promise<Activity> {
  const res = await fetch('http://localhost:8000/activity', {
    method: 'POST',
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error('failed to create activity');

  return await res.json();
}

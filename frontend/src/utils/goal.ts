export type Goal = {
  id: string;
  name: string;
  dur: string;
  durType: string;
  durPer: string;
};

export async function getGoals(): Promise<Goal[]> {
  const res = await fetch(`http://localhost:8000/goals`, { credentials: 'include' });

  if (!res.ok) throw new Error('failed to fetch goals');

  return await res.json();
}

export type CreateGoalData = Pick<Goal, 'name' | 'dur' | 'durPer' | 'durType'>;
export async function createGoal(data: CreateGoalData): Promise<Goal> {
  const res = await fetch(`http://localhost:8000/goals`, {
    method: 'POST',
    body: JSON.stringify(data),
		credentials: 'include',
  });

  if (!res.ok) throw new Error('failed to create new goal');

  return await res.json();
}

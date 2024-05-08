import { BACKEND_URL } from './host';

export type Goal = {
  id: string;
  name: string;
  dur: string;
  durType: string;
  durPer: string;
};

export async function getGoals(): Promise<{ goals: Goal[] }> {
  const res = await fetch(`${BACKEND_URL}/goals`, {
    credentials: 'include',
  });

  if (!res.ok) throw new Error('failed to fetch goals');

  return await res.json();
}

export type CreateGoalData = Pick<Goal, 'name' | 'dur' | 'durPer' | 'durType'>;
export async function createGoal(data: CreateGoalData): Promise<Goal> {
  const res = await fetch(`${BACKEND_URL}/goals`, {
    method: 'POST',
    body: JSON.stringify(data),
    credentials: 'include',
  });

  if (!res.ok) throw new Error('failed to create new goal');

  return await res.json();
}

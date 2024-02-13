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
  // TODO: point to backend
  return await new Promise((resolve, reject) => {
    setTimeout(() => resolve({ id: 'jskdjkf', goalId: 'kfjkdjf' }));
  });
}

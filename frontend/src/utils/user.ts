export type User = {
  id: string;
  name: string;
  username: string;
};

export async function getUser(id: string): Promise<User> {
  const res = await fetch(`http://localhost:8000/User/${id}`, {
    headers: { 'Content-Type': 'application/json' },
  });
  if (!res.ok) throw new Error('failed to fetch user');
  return await res.json();
}

export type CreateUserData = {
  name: string;
  username: string;
  password: string;
};

export async function createUser(data: CreateUserData): Promise<User> {
  const res = await fetch('http://localhost:8000/CreateUser', {
    body: JSON.stringify(data),
    method: 'POST',
  });

  if (!res.ok) throw new Error('failed to fetch user');

  return await res.json();
}

export type LoginUserData = {
  username: string;
  password: string;
};

export async function loginUser(data: LoginUserData): Promise<User> {
  const res = await fetch('http://localhost:8000/login', {
    body: JSON.stringify(data),
    method: 'POST',
  });

  if (!res.ok) throw new Error('failed to log user in');

  return await res.json();
}

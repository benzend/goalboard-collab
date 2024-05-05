import { BACKEND_API_HOST } from './host';

export type User = {
  id: string;
  name: string;
  username: string;
};

export async function getUser(id: string): Promise<User> {
  const res = await fetch(`${BACKEND_API_HOST}/User/${id}`, {
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

export type CreateUserReturnData = {
  user: User;
  token: string;
};

export async function createUser(
  data: CreateUserData
): Promise<CreateUserReturnData> {
  console.log('you know, youre supposed to refresh');
  const res = await fetch(`${BACKEND_API_HOST}/register`, {
    mode: 'cors',
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

export type LoginUserReturnData = {
  user: User;
  token: string;
};

export async function loginUser(
  data: LoginUserData
): Promise<LoginUserReturnData> {
  const res = await fetch(`${BACKEND_API_HOST}/login`, {
    body: JSON.stringify(data),
    method: 'POST',
  });

  if (!res.ok) throw new Error('failed to log user in');

  return await res.json();
}

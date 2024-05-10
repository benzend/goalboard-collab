export type User = {
  id: string;
  name: string;
  username: string;
};

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
  const res = await fetch('/api/register', {
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
  const res = await fetch('/api/login', {
    body: JSON.stringify(data),
    method: 'POST',
  });

  if (!res.ok) throw new Error('failed to log user in');

  return await res.json();
}

export async function logoutUser() {
  const res = await fetch('/api/logout', {
    method: 'POST',
  });

  if (!res.ok) throw new Error('failed to log user out');
}

export type GetCurrentUserReturnData = {
  user: User;
};

export async function getCurrentUser(): Promise<GetCurrentUserReturnData> {
  const res = await fetch('/api/current_user');

  if (!res.ok) throw new Error('failed to grab user');

  return await res.json();
}

import { useContext, createContext, useState } from 'react';
import { useNavigate, Outlet } from 'react-router-dom';
import { User, createUser, loginUser } from '../utils/user';

export type LoginActionData = {
  username: string;
  password: string;
};

export type RegisterActionData = {
  username: string;
  password: string;
  name: string;
};

export type AuthContextType = null | {
  token: null | string;
  user: null | User;
  loginAction: Function;
  logOut: Function;
  register: Function;
};

const AuthContext = createContext<AuthContextType>(null);

export const AuthProvider = () => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState(localStorage.getItem('site') || '');
  const navigate = useNavigate();
  const loginAction = async (data: LoginActionData) => {
    try {
      console.log('before res');
      const res = await loginUser({
        username: data.username,
        password: data.password,
      });
      console.log({ res });
      if (res) {
        setUser(res.user);
        setToken(res.token);
        localStorage.setItem('site', res.token);
        navigate(`/${res.user.id}/goals`);
        return;
      } else {
        throw new Error('failed to sign in');
      }
    } catch (err) {
      console.error(err);
    }
  };

  const logOut = () => {
    setUser(null);
    setToken('');
    localStorage.removeItem('site');
    navigate('/login');
  };

  const register = async (data: RegisterActionData) => {
    try {
      const res = await createUser({
        username: data.username,
        name: data.name,
        password: data.password,
      });

      if (res) {
        setUser(res.user);
        setToken(res.token);
        localStorage.setItem('site', res.token);
        navigate(`/${res.user.id}/goals/new`);
        return;
      }

      throw new Error('failed to register user');
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <AuthContext.Provider
      value={{ token, user, loginAction, logOut, register }}
    >
      <Outlet />
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  return useContext(AuthContext);
};

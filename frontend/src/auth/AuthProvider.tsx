import { useContext, createContext, useState, useEffect } from 'react';
import { useNavigate, Outlet } from 'react-router-dom';
import {
  User,
  createUser,
  getCurrentUser,
  loginUser,
  logoutUser,
} from '../utils/user';
import { deleteCookie, setCookie } from '../utils/cookie';

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
      const res = await loginUser({
        username: data.username,
        password: data.password,
      });
      if (res) {
        setUser(res.user);
        setToken(res.token);
        // setCookie('jwt_token', res.token);
        navigate(`/goals`);
      } else {
        throw new Error('failed to sign in');
      }
    } catch (err) {
      console.error(err);
    }
  };

  const logOut = async () => {
    setUser(null);
    setToken('');
    // deleteCookie('jwt_token');
    await logoutUser();
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
        // setCookie('jwt_token', res.token);
        navigate('/goals');
      } else {
        throw new Error('failed to register user');
      }
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    if (!user) {
      getCurrentUser()
        .then((res) => {
          setUser(res.user);
        })
        .catch((err) => {
          console.error(err);
        });
    }
  }, []);

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

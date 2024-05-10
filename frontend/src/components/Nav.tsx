import { Link } from 'react-router-dom';
import { useAuth } from '../auth/AuthProvider';

export const Nav = () => {
  const auth = useAuth();

  const logout = () => {
    auth?.logOut();
  };
  return (
    <nav className="bg-primary p-4 flex justify-between align-middle">
      <section className="flex align-middle">
        <Link to="/">
          <span className="text-lg font-bold text-white">Goalboard</span>
        </Link>
      </section>
      <section className="flex align-middle">
        {auth?.user ? (
          <button onClick={logout} className="font-bold text-white">
            Logout
          </button>
        ) : (
          <Link to="/login">
            <span className="font-bold text-white">Login</span>
          </Link>
        )}
      </section>
    </nav>
  );
};

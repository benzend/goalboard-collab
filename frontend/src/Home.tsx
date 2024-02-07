import 'react';
import { Link } from 'react-router-dom';

export const Home = () => {
  return (
    <>
      <h1>Goalboard</h1>

      <Link to="/login">Sign in</Link>
      <Link to="/register">Sign up</Link>
    </>
  );
};

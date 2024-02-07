import 'react';
import { Link } from 'react-router-dom';
import { Nav } from './components/Nav';

export const Home = () => {
  return (
    <div className="min-h-screen">
      <Nav />
      <main className="pt-20 text-center">
        <h1 className="text-4xl font-bold mb-20">Goalboard</h1>

        <section>
          <Link to="/login">
            <span className="rounded-2xl bg-cyan-600 text-white inline-block w-20 py-2">
              Sign in
            </span>
          </Link>
          <br className="mb-6" />
          <Link to="/register">
            <span className="rounded-2xl border border-cyan-600 text-cyan-600 inline-block w-20 py-2">
              Sign up
            </span>
          </Link>
        </section>
      </main>
    </div>
  );
};

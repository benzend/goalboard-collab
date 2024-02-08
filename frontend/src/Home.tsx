import 'react';
import { Link } from 'react-router-dom';
import { Nav } from './components/Nav';
import cn from 'classnames';
import { getButtonClasses } from './components/Button';

export const Home = () => {
  return (
    <div className="min-h-screen">
      <Nav />
      <main className="pt-20 text-center">
        <h1 className="text-4xl font-bold mb-20">Goalboard</h1>

        <section>
          <Link to="/login">
            <span
              className={cn(
                getButtonClasses('primary'),
                'inline-block w-20 py-2 px-0'
              )}
            >
              Sign in
            </span>
          </Link>
          <br className="mb-6" />
          <Link to="/register">
            <span
              className={cn(
                getButtonClasses('outline'),
                'inline-block w-20 py-2 px-0'
              )}
            >
              Sign up
            </span>
          </Link>
        </section>
      </main>
    </div>
  );
};

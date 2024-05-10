import 'react';
import { Link } from 'react-router-dom';
import { Nav } from './components/Nav';
import cn from 'classnames';
import { getButtonClasses } from './components/Button';
import { Heading } from './components/Heading';

export const Home = () => {
  return (
    <div className="min-h-screen">
      <Nav />
      <main className="pt-20 text-center max-w-lg mx-auto">
        <div className="flex justify-center items-center h-[60vh]">
          <div>
            <Heading el="h1" type="h1" className="mb-20">
              Goalboard
            </Heading>

            <section className="flex gap-4 justify-center flex-col md:flex-row">
              <Link to="/login">
                <span
                  className={cn(
                    getButtonClasses('primary'),
                    'inline-block w-24 py-2 px-0'
                  )}
                >
                  Sign in
                </span>
              </Link>
              <Link to="/register">
                <span
                  className={cn(
                    getButtonClasses('outline'),
                    'inline-block w-24 py-2 px-0'
                  )}
                >
                  Sign up
                </span>
              </Link>
            </section>
          </div>
        </div>
      </main>
    </div>
  );
};

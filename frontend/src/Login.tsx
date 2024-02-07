import { Form, Link, useNavigate } from 'react-router-dom';
import { Nav } from './components/Nav';
import { FormEvent, FormEventHandler } from 'react';

export const Login = () => {
  const navigate = useNavigate();
  const handleSubmit: FormEventHandler = (e: FormEvent) => {
    e.preventDefault();

    const target = e.target as typeof e.target & {
      email: { value: string };
      password: { value: string };
    };

    const email = target.email.value;
    const password = target.password.value;

    console.log({ email, password });

    // TODO: have an actual ID
    navigate('/ksjldfksjf/goals');

    // TODO: Handle this in our backend
  };
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10">
        <h1 className="text-xl font-semibold mb-10">Sign in</h1>

        <Form onSubmit={handleSubmit}>
          <label htmlFor="email">
            <span className="font-thin">Email</span>
            <br />
            <input
              type="email"
              name="email"
              id="email"
              className="border border-black w-full px-2 py-1 rounded shadow-lg"
            />
          </label>

          <br className="mb-10" />

          <label htmlFor="password">
            <span className="font-thin">Password</span>
            <br />
            <input
              type="password"
              name="password"
              id="password"
              className="border border-black w-full px-2 py-1 rounded shadow-lg"
            />
          </label>

          <br className="mb-10" />

          <Link to="/register">
            <span className="text-cyan-600 underline">Sign up</span>
          </Link>

          <button
            type="submit"
            className="fixed bottom-2 right-2 bg-cyan-500 text-white px-4 py-2 rounded-2xl"
          >
            Sign in
          </button>

          {/**
           * TODO: Handle for password forgot
           */}
        </Form>
      </main>
    </div>
  );
};

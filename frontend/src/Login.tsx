import { Form, Link, useNavigate } from 'react-router-dom';
import { Nav } from './components/Nav';
import { FormEvent, FormEventHandler } from 'react';
import { Button } from './components/Button';
import { Input } from './components/Input';

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
            <Input
              type="email"
              name="email"
              id="email"
              inputStyleType="w-full"
            />
          </label>

          <br className="mb-10" />

          <label htmlFor="password">
            <span className="font-thin">Password</span>
            <br />
            <Input
              type="password"
              name="password"
              id="password"
              inputStyleType="w-full"
            />
          </label>

          <br className="mb-10" />

          <Link to="/register">
            <span className="text-cyan-600 underline">Sign up</span>
          </Link>

          <Button type="submit" className="fixed bottom-2 right-2">
            Sign in
          </Button>

          {/**
           * TODO: Handle for password forgot
           */}
        </Form>
      </main>
    </div>
  );
};

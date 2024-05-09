import { Form, Link, useNavigate } from 'react-router-dom';
import { Nav } from './components/Nav';
import { FormEvent, FormEventHandler } from 'react';
import { Button } from './components/Button';
import { Input } from './components/Input';
import { Heading } from './components/Heading';
import { useAuth } from './auth/AuthProvider';

export const Login = () => {
  const navigate = useNavigate();
  const auth = useAuth();

  if (auth?.user) {
    navigate('/goals');
  }

  const handleSubmit: FormEventHandler = (e: FormEvent) => {
    e.preventDefault();

    const target = e.target as typeof e.target & {
      // email: { value: string };
      username: { value: string };
      password: { value: string };
    };

    // const email = target.email.value;
    const username = target.username.value;
    const password = target.password.value;

    console.debug({ username, password });

    /**
     * TODO: Figure out how to store user validation internally
     * TODO: so that we can move around the app freely once
     * TODO: logged in
     */

    auth?.loginAction({ username, password });
    // loginUser({ username, password })
    //   .then((user) => {
    //     console.debug({ user });
    //     navigate(`/${user.id}/goals`);
    //   })
    //   .catch((err) => {
    //     console.error({ err });
    //     alert('failed to login');
    //   });
  };
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10">
        <Heading el="h1" type="h2" className="mb-10">
          Sign in
        </Heading>

        <Form onSubmit={handleSubmit}>
          {/* <label htmlFor="email">
            <span className="font-thin">Email</span>
            <br />
            <Input
              type="email"
              name="email"
              id="email"
              inputStyleType="w-full"
            />
          </label>

          <br className="mb-10" /> */}

          <label htmlFor="username">
            <span className="font-thin">Username</span>
            <br />
            <Input
              type="text"
              name="username"
              id="username"
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

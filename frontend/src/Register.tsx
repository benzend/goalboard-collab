import { FormEventHandler } from 'react';
import { useNavigate, Form, Link } from 'react-router-dom';
import { Nav } from './components/Nav';
import { Input } from './components/Input';
import { Heading } from './components/Heading';
import { Button } from './components/Button';

export const Register = () => {
  const navigate = useNavigate();
  const handleSubmit: FormEventHandler = (e) => {
    e.preventDefault();

    const target = e.target as typeof e.target & {
      email: { value: string };
      password: { value: string };
      confirmPassword: { value: string };
    };

    const email = target.email.value;
    const password = target.password.value;
    const confirmPassword = target.confirmPassword.value;

    console.log({ email, password, confirmPassword });

    if (password !== confirmPassword) throw new Error('passwords do not match');
    if (password.length < 8) throw new Error('password is too short');
    if (password.length > 50)
      throw new Error('weird, why is your password this long?');

    // TODO: pass up to backend for account creation
    navigate('/ksjldfksjf/goals/new');
  };
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10">
        <Heading el="h1" type="h2" className="mb-10">
          Sign up
        </Heading>

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

          <label htmlFor="confirmPassword">
            <span className="font-thin">Confirm Password</span>
            <br />
            <Input
              type="password"
              name="comfirmPassword"
              id="confirmPassword"
              inputStyleType="w-full"
            />
          </label>

          <br className="mb-10" />

          <Link to="/register">
            <span className="text-cyan-600 underline">Sign in</span>
          </Link>

          <Button type="submit" className="fixed bottom-2 right-2">
            Sign Up
          </Button>

          {/**
           * TODO: Handle for password forgot
           */}
        </Form>
      </main>
    </div>
  );
};

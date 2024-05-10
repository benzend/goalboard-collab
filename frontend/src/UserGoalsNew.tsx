import { FormEvent, FormEventHandler } from 'react';
import { Form, useNavigate } from 'react-router-dom';
import { Nav } from './components/Nav';
import { Button } from './components/Button';
import { Input } from './components/Input';
import { Select } from './components/Select';
import { Heading } from './components/Heading';
import { createGoal } from './utils/goal';

export const UserGoalsNew = () => {
  const navigate = useNavigate();
  const handleSubmit: FormEventHandler = (e: FormEvent) => {
    e.preventDefault();

    const target = e.target as typeof e.target & {
      name: { value: string };
      dur: { value: string };
      durType: { value: 'seconds' | 'minutes' | 'hours' | 'days' | 'weeks' };
      durPer: { value: string };
    };

    const name = target.name.value;
    const dur = target.dur.value;
    const durType = target.durType.value;
    const durPer = target.durPer.value;

    console.log({ name, dur, durType, durPer });

    createGoal({ name, dur, durPer, durType })
      .then((goal) => {
        console.debug({ goal });
        navigate('/sdkjfkajdf/goals');
      })
      .catch((err) => {
        console.error({ err });
        alert('failed to create new goal');
      });
  };
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10 max-w-lg mx-auto">
        <Heading el="h1" type="h2" className="mb-10">
          Create new Goal
        </Heading>

        <Form onSubmit={handleSubmit}>
          <label htmlFor="name">
            <span className="font-thin">Goal</span>
            <br />
            <Input
              type="text"
              name="name"
              id="goalNameInput"
              inputStyleType="w-full"
            />
          </label>

          <br className="mb-10" />

          <label htmlFor="dur">
            <span className="font-thin">For</span>
            <br />
            <Input
              type="number"
              name="dur"
              id="goalDurInput"
              inputStyleType="w-lg"
            />

            <Select
              name="durType"
              id="durTypeSelect"
              className="border border-black px-2 h-10 rounded shadow-lg bg-white ml-2"
              selectStyleType="default"
            >
              <option value="seconds">Seconds</option>
              <option value="minutes">Minutes</option>
              <option value="hours">Hours</option>
              <option value="days">Days</option>
              <option value="weeks">Weeks</option>
            </Select>
          </label>

          <br className="mb-10" />

          <label htmlFor="durPer">
            <span className="font-thin">Every</span>
            <br />
            <Select
              name="durPer"
              id="durPerSelect"
              className="border border-black px-2 h-10 rounded shadow-lg bg-white"
              selectStyleType="default"
            >
              <option value="day">Day</option>
              <option value="week">Week</option>
              <option value="month">Month</option>
            </Select>
          </label>

          <Button type="submit" className="fixed bottom-2 right-2">
            Create
          </Button>
        </Form>
      </main>
    </div>
  );
};

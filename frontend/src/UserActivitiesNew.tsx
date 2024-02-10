import { FormEvent, FormEventHandler } from 'react';
import {
  Form,
  useNavigate,
  LoaderFunction,
  useLoaderData,
} from 'react-router-dom';
import { Nav } from './components/Nav';
import { Button } from './components/Button';
import { Input } from './components/Input';
import { Select } from './components/Select';
import { Heading } from './components/Heading';
import { LoaderData } from './utils/loader-data';

export const loader = (async () => {
  // TODO: redirect to login if user unauthed
  return {
    goals: [
      { id: 'ksjdfkjsdf', name: 'Read', dur: 'daily' },
      { id: 'ksjdfksf', name: 'Run', dur: 'daily' },
    ],
  };
}) satisfies LoaderFunction;

export const UserActivitiesNew = () => {
  const navigate = useNavigate();
  const data = useLoaderData() as LoaderData<typeof loader>;
  const handleSubmit: FormEventHandler = (e: FormEvent) => {
    e.preventDefault();

    const target = e.target as typeof e.target & {
      goal: { value: string };
      dur: { value: string };
      durType: { value: 'seconds' | 'minutes' | 'hours' | 'days' | 'weeks' };
    };

    const goal = target.goal.value;
    const dur = target.dur.value;
    const durType = target.durType.value;

    console.log({ goal, dur, durType });

    /**
     * TODO: submit activity to backend
     */

    navigate('/sdkjfkajdf/goals');
  };
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10">
        <Heading el="h1" type="h2" className="mb-10">
          Track an Activity
        </Heading>

        <Form onSubmit={handleSubmit}>
          <label htmlFor="goal">
            <span className="font-thin">Goal</span>
            <br />
            <Select selectStyleType="w-full" name="goal" id="goalSelect">
              {data.goals.map((goal) => (
                <option value={goal.id}>{goal.name}</option>
              ))}
            </Select>
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

          <Button type="submit" className="fixed bottom-2 right-2">
            Track
          </Button>
        </Form>
      </main>
    </div>
  );
};

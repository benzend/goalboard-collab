import { FormEvent, FormEventHandler } from 'react';
import { Form, useNavigate } from 'react-router-dom';
import { Nav } from './components/Nav';

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

    navigate('/sdkjfkajdf/goals');
  };
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10">
        <h1 className="text-xl font-semibold mb-10">Create new Goal</h1>

        <Form onSubmit={handleSubmit}>
          <label htmlFor="name">
            <span className="font-thin">Goal</span>
            <br />
            <input
              type="text"
              name="name"
              id="goalNameInput"
              className="border border-black w-full px-2 h-10 rounded shadow-lg"
            />
          </label>

          <br className="mb-10" />

          <label htmlFor="dur">
            <span className="font-thin">For</span>
            <br />
            <input
              type="number"
              name="dur"
              id="goalDurInput"
              className="border border-black px-2 h-10 rounded shadow-lg w-20"
            />

            <select
              name="durType"
              id="durTypeSelect"
              className="border border-black px-2 h-10 rounded shadow-lg bg-white ml-2"
            >
              <option value="seconds">Seconds</option>
              <option value="minutes">Minutes</option>
              <option value="hours">Hours</option>
              <option value="days">Days</option>
              <option value="weeks">Weeks</option>
            </select>
          </label>

          <br className="mb-10" />

          <label htmlFor="durPer">
            <span className="font-thin">Every</span>
            <br />
            <select
              name="durPer"
              id="durPerSelect"
              className="border border-black px-2 h-10 rounded shadow-lg bg-white"
            >
              <option value="day">Day</option>
              <option value="week">Week</option>
              <option value="month">Month</option>
            </select>
          </label>

          <button
            type="submit"
            className="fixed bottom-2 right-2 bg-cyan-600 text-white px-5 py-3 rounded-2xl"
          >
            Create
          </button>
        </Form>
      </main>
    </div>
  );
};

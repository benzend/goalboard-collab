import { LoaderFunction, useLoaderData } from 'react-router-dom';
import { Nav } from './components/Nav';
import { LoaderData } from './utils/loader-data';

export const loader = (async () => {
  // TODO: redirect to login if user unauthed
  return {
    goals: [
      { name: 'Read', dur: 'daily' },
      { name: 'Read', dur: 'daily' },
    ],
  };
}) satisfies LoaderFunction;

export const UserGoals = () => {
  const data = useLoaderData() as LoaderData<typeof loader>;
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10  s">
        {data.goals.length ? (
          data.goals.map((g) => {
            // TODO: handle real goals
            return (
              <section className="flex justify-between align-middle mb-6">
                <h2 className="text-lg font-semibold self-center">{g.name}</h2>

                <section className="text-right">
                  {/**
                   * TODO: Inject correct goal durations
                   */}
                  <p>Today's goal: 1hr</p>
                  {/**
                   * TODO: Inject correct goal durations
                   */}
                  <p>Time tracked: 20min</p>
                </section>
              </section>
            );
          })
        ) : (
          <section>No goals</section>
        )}
      </main>
    </div>
  );
};
